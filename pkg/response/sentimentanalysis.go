package response

import (
	"io"
	"time"
)

type SentimentAnalysis struct {
	Headers
	Status
	Timestamp      string    `json:"timestamp"`
	Time           int       `json:"time"`
	Lang           string    `json:"lang"`
	LangConfidence *float64  `json:"langConfidence,omitempty"`
	Text           *string   `json:"text,omitempty"`
	Sentiment      Sentiment `json:"sentiment"`
}

type Sentiment struct {
	Type  string  `json:"type"`
	Score float64 `json:"score"`
}

func (s *SentimentAnalysis) Decode(body io.ReadCloser) error {
	return decode(body, s)
}

func (s *SentimentAnalysis) SetHeaders(units, unitsLeft float64, unitsReset time.Time) {
	s.Headers.Set(units, unitsLeft, unitsReset)
}
