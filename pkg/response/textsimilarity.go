package response

import (
	"io"
	"time"
)

type TextSimilarity struct {
	Headers
	Status
	Timestamp      string  `json:"timestamp"`
	Time           int     `json:"time"`
	Lang           string  `json:"lang"`
	LangConfidence float64 `json:"langConfidence"`
	Text1          *string `json:"text1,omitempty"`
	Text2          *string `json:"text2,omitempty"`
	Similarity     float64 `json:"similarity"`
}

func (t *TextSimilarity) Decode(body io.ReadCloser) error {
	return decode(body, t)
}

func (t *TextSimilarity) SetHeaders(units, unitsLeft float64, unitsReset time.Time) {
	t.Headers.Set(units, unitsLeft, unitsReset)
}
