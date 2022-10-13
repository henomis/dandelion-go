package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type SentimentAnalysis struct {
	// The following three parameters are mutually exclusive
	Text         string  `json:"text,omitempty" validate:"-"`
	HTML         string  `json:"html,omitempty" validate:"-"`
	HTMLFragment string  `json:"html_fragment,omitempty" validate:"-"`
	Lang         *string `json:"lang,omitempty" validate:"omitempty,oneof=it en auto"`
}

func (s *SentimentAnalysis) Validate() error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"text":          s.Text,
			"html":          s.HTML,
			"html_fragment": s.HTMLFragment,
		},
	)

}

func (s *SentimentAnalysis) ToPostForm() *postform.PostForm {

	postForm := postform.New()

	postForm.Add("text", &s.Text)
	postForm.Add("html", &s.HTML)
	postForm.Add("html_fragment", &s.HTMLFragment)
	postForm.Add("lang", s.Lang)

	return postForm
}
