package esi

import (
	"encoding/json"
	"net/http"

	"github.com/eveisesi/neo"
	"github.com/pkg/errors"
)

func (s *service) GetStatus() (*neo.ServerStatus, *Meta) {

	response, m := s.request(request{
		method: http.MethodGet,
		path:   "/v1/status",
	})
	if m.IsError() {
		return nil, m
	}

	status := new(neo.ServerStatus)
	err := json.Unmarshal(response, status)
	if err != nil {
		m.Msg = errors.Wrapf(err, "unable to unmarshal response body on request %s", "/v1/status")
		return nil, m
	}

	return status, m

}
