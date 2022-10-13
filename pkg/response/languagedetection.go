package response

import (
	"io"
	"time"
)

type LanguageDetection struct {
	Headers
	Status
	Timestamp     string         `json:"timestamp"`
	Time          int            `json:"time"`
	DetectedLangs []DetectedLang `json:"detectedLangs,omitempty"`
}

type DetectedLang struct {
	Lang        string  `json:"lang"`
	Confindence float64 `json:"confidence"`
}

func (l *LanguageDetection) Decode(body io.ReadCloser) error {
	return decode(body, l)
}

func (l *LanguageDetection) SetHeaders(units, unitsLeft float64, unitsReset time.Time) {
	l.Headers.Set(units, unitsLeft, unitsReset)
}
