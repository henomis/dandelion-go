package response

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Headers struct {
	XDLUnits      float64   `json:"X-DL-Units"`
	XDLUnitsLeft  float64   `json:"X-DL-Units-Left"`
	XDLUnitsReset time.Time `json:"X-DL-Units-Reset"`
}

func (h *Headers) Set(units, unitsLeft float64, unitsReset time.Time) {
	h.XDLUnits = units
	h.XDLUnitsLeft = unitsLeft
	h.XDLUnitsReset = unitsReset
}

type Status struct {
	Message string            `json:"message,omitempty"`
	Code    string            `json:"code,omitempty"`
	Data    map[string]string `json:"data,omitempty"`
	IsError bool              `json:"error,omitempty"`
}

func (s *Status) IsSuccess() bool {
	return !s.IsError
}

func (s *Status) Error() error {
	return fmt.Errorf("%s: %s [%v]", s.Code, s.Message, s.Data)
}

// Support methods

func decode(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(data)
}
