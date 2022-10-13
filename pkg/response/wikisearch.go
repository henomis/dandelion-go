package response

import (
	"io"
	"time"
)

type WikiSearch struct {
	Headers
	Status
	Timestamp string             `json:"timestamp"`
	Time      int                `json:"time"`
	Count     int                `json:"count"`
	Query     string             `json:"query"`
	Lang      string             `json:"lang"`
	Entities  []WikiSearchEntity `json:"entities,omitempty"`
}

type WikiSearchEntity struct {
	Weight     float64  `json:"weight"`
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	URI        string   `json:"uri"`
	Label      string   `json:"label"`
	Types      []string `json:"types,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Abstract   *string  `json:"abstract,omitempty"`
	Lod        *Lod     `json:"lod,omitempty"`
	Image      *Image   `json:"image,omitempty"`
}

func (w *WikiSearch) Decode(body io.ReadCloser) error {
	return decode(body, w)
}

func (w *WikiSearch) SetHeaders(units, unitsLeft float64, unitsReset time.Time) {
	w.Headers.Set(units, unitsLeft, unitsReset)
}
