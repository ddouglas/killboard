package main

import (
	"context"
	"log"
	"os"
	"time"

	core "github.com/eveisesi/neo/app"
	"github.com/eveisesi/neo/server"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var (
	app *cli.App
)

func init() {
	_ = godotenv.Load(".env")

	app = cli.NewApp()
	app.Name = "neo"
	app.UsageText = "neo [parent] [child command] [--options]"
	app.Usage = "Service that manages all services related to Neo and its stable operation"
	app.Version = "v0.0.1"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "killmail",
			Usage:       "Parent command for all administrative task around killmails",
			Subcommands: killmailCommands(),
		},
		cronCommand(),
		// cli.Command{
		// 	Name:  "restore",
		// 	Usage: "Reaches out to the DO Space and restore killmails that may have of been lost",
		// 	Action: func(c *cli.Context) error {

		// 		app := core.New()

		// 		client := s3.New(app.Spaces)
		// 		// hash=ef45d50a8ab98eb33fd386214dd0c906b6444a52 id=80621595
		// 		input := &s3.HeadObjectInput{
		// 			Bucket: aws.String("neo"),
		// 			Key:    aws.String(fmt.Sprintf("killmails/%d:%s.json", 80625595, "ef45d50a8ab98eb33fd386214dd0c906b6444a52")),
		// 		}

		// 		object, err := client.HeadObject(input)
		// 		if err != nil {
		// 			spew.Dump(err)
		// 			app.Logger.WithError(err).Error("failed to list objects")
		// 			return err
		// 		}

		// 		spew.Dump(object)

		// 		return nil

		// 	},
		// },
		cli.Command{
			Name:  "history",
			Usage: "Reaches out to the Zkillboard API and downloads historical killmail hashes, then reaches out to CCP for Killmail Data",
			Action: func(c *cli.Context) error {
				app := core.New()
				maxdate := c.String("maxdate")
				mindate := c.String("mindate")
				threshold := c.Int64("threshold")
				datehold := c.Bool("datehold")

				err := app.Killmail.HistoryExporter(mindate, maxdate, datehold, threshold)
				if err != nil {
					return cli.NewExitError(err, 1)
				}

				return cli.NewExitError(nil, 0)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "maxdate",
					Usage:    "Date to start the loop at when calling the zkillboard history api. (Format: YYYYMMDD)",
					Required: true,
				},
				cli.StringFlag{
					Name:     "mindate",
					Usage:    "Date to stop the history loop at when calling zkillboard history api. (Format: YYYYMMDD)",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "datehold",
					Usage: "Hold after each date until the processing queue has reached a threshold. Threshold must be defined, else this command will be ignored",
				},
				cli.IntFlag{
					Name:  "threshold",
					Usage: "Threshold that the queue must be below process processing the next date",
				},
			},
		},
		cli.Command{
			Name:   "serve",
			Usage:  "Starts an HTTP Server to serve killmail data",
			Action: server.Action,
		},

		cli.Command{
			Name:  "listen",
			Usage: "Opens a WSS Connection to ZKillboard and lsitens to the stream",
			Action: func(c *cli.Context) error {
				_ = core.New().Killmail.Websocket()

				return nil
			},
		},
		cli.Command{
			Name: "top",
			Action: func(c *cli.Context) error {
				return core.New().Top.Run()
			},
		},
		cli.Command{
			Name: "tracking",
			Action: func(c *cli.Context) error {

				app := core.New()

				beginning := time.Now().In(time.UTC)
				start := time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 10, 58, 0, 0, time.UTC)
				end := time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 11, 25, 0, 0, time.UTC)

				app.Tracker.Run(start, end)

				return nil
			},
		},
		cli.Command{
			Name:        "autocompleter",
			Description: "Manually rebuild the autocompleter index",
			Action: func(c *cli.Context) error {

				app := core.New()

				err := app.Search.Build()
				if err != nil {
					return cli.NewExitError(err, 1)
				}

				return nil
			},
		},
		cli.Command{
			Name:        "notifications",
			Description: "Notifications subscribe to the a Redis PubSub. When the importer detects a killmail with a value greater than the configured notification value, it publishes the id and hash to this pubsub and this service will format the message for slack and post the killmail to slack",
			Action: func(c *cli.Context) error {
				app := core.New()

				app.Notification.Run()
				return nil
			},
		},
		cli.Command{
			Name:        "updater",
			Description: "Updater ensures that all updatable records in the database are update date according to their CacheUntil timestamp.",
			Action: func(c *cli.Context) error {
				app := core.New()
				var ctx = context.Background()

				ch := make(chan int, 3)

				go app.Character.UpdateExpired(ctx)
				go app.Corporation.UpdateExpired(ctx)
				go app.Alliance.UpdateExpired(ctx)

				<-ch

				return nil

			},
		},
		migrateCommand(),
		cli.Command{
			Name:  "recalculate",
			Usage: "Dispatches Go Routines to handle recalculable killmails in the recalculate queue",
			Action: func(c *cli.Context) error {
				app := core.New()

				workers := c.Int64("workers")

				app.Killmail.Recalculator(workers)

				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "workers",
					Usage: "Number of Go Routines to that should be used to process messages.",
					Value: 10,
				},
			},
		},
		cli.Command{
			Name:  "recalculable",
			Usage: "Finds Killmails where the DestroyedValue and the DroppedValue do not equal the TotalValue and dispatches them to a queue to have these properties recalculated",
			Action: func(c *cli.Context) error {

				app := core.New()

				limit := c.Int64("limit")
				trigger := c.Int64("trigger")
				after := c.Uint64("after")

				app.Killmail.RecalculatorDispatcher(limit, trigger, after)

				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "limit",
					Usage: "number of records to fetch from the db",
					Value: 10000,
				},
				cli.IntFlag{
					Name:  "trigger",
					Usage: "this number of less must remain on the queue before triggering another pull from the db",
					Value: 2500,
				},
				cli.Int64Flag{
					Name:  "after",
					Usage: "Start at a specific killmail id",
					Value: 0,
				},
			},
		},
		cli.Command{
			Name:        "market",
			Usage:       "Updates market prices in the Db",
			Subcommands: marketCommands(),
		},
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
