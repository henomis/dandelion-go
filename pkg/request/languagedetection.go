package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type LanguageDetection struct {
	// The following three parameters are mutually exclusive
	Text         string `json:"text,omitempty" validate:"-"`
	HTML         string `json:"html,omitempty" validate:"-"`
	HTMLFragment string `json:"html_fragment,omitempty" validate:"-"`
	Clean        *bool  `json:"clean,omitempty" validate:"-"`
}

func (l *LanguageDetection) Validate() error {
	validate := validator.New()

	if err := validate.Struct(l); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"text":          l.Text,
			"html":          l.HTML,
			"html_fragment": l.HTMLFragment,
		},
	)

}

func (l *LanguageDetection) ToPostForm() *postform.PostForm {

	postForm := postform.New()

	postForm.Add("text", &l.Text)
	postForm.Add("html", &l.HTML)
	postForm.Add("html_fragment", &l.HTMLFragment)
	postForm.AddBool("clean", l.Clean)

	return postForm
}
