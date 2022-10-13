package response

import (
	"io"
	"time"
)

type EntityExtraction struct {
	Headers
	Status
	Timestamp      string       `json:"timestamp"`
	Time           int          `json:"time"`
	Lang           string       `json:"lang"`
	LangConfidence float64      `json:"langConfidence"`
	Text           *string      `json:"text,omitempty"`
	Annotations    []Annotation `json:"annotations,omitempty"`
	TopEntities    []Entity     `json:"topEntities,omitempty"`
}

type Lod struct {
	Wikipedia *string `json:"wikipedia,omitempty"`
	Dbpedia   *string `json:"dbpedia,omitempty"`
}

type Image struct {
	Full      *string `json:"full,omitempty"`
	Thumbnail *string `json:"thumbnail,omitempty"`
}

type Annotation struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	URI             string   `json:"uri"`
	Label           string   `json:"label"`
	Confidence      float64  `json:"confidence"`
	Spot            string   `json:"spot"`
	Start           int      `json:"start"`
	End             int      `json:"end"`
	Types           []string `json:"types,omitempty"`
	Categories      []string `json:"categories,omitempty"`
	Abstract        *string  `json:"abstract,omitempty"`
	Lod             *Lod     `json:"lod,omitempty"`
	AlternateLabels []string `json:"alternateLabels,omitempty"`
	Image           *Image   `json:"image,omitempty"`
}

type Entity struct {
	ID    string `json:"id"`
	URI   string `json:"uri"`
	Score string `json:"score"`
}

func (e *EntityExtraction) Decode(body io.ReadCloser) error {
	return decode(body, e)
}

func (e *EntityExtraction) SetHeaders(units, unitsLeft float64, unitsReset time.Time) {
	e.Headers.Set(units, unitsLeft, unitsReset)
}
