package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type TextSimilarity struct {
	// The following three parameters are mutually exclusive
	Text1         string `json:"text1,omitempty" validate:"-"`
	HTML1         string `json:"html1,omitempty" validate:"-"`
	HTMLFragment1 string `json:"html_fragment1,omitempty" validate:"-"`
	// The following three parameters are mutually exclusive
	Text2         string  `json:"text2,omitempty" validate:"-"`
	HTML2         string  `json:"html2,omitempty" validate:"-"`
	HTMLFragment2 string  `json:"html_fragment2,omitempty" validate:"-"`
	Lang          *string `json:"lang,omitempty" validate:"omitempty,oneof=de en es fr it pt ru auto"`
	Bow           *string `json:"bow,omitempty" validate:"omitempty,oneof=always one_empty both_empty never"`
}

func (t *TextSimilarity) Validate() error {
	validate := validator.New()

	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if err := validateMutualExclusiveFields(
		map[string]string{
			"text1":          t.Text1,
			"html1":          t.HTML1,
			"html_fragment1": t.HTMLFragment1,
		},
	); err != nil {
		return err
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"text2":          t.Text2,
			"html2":          t.HTML2,
			"html_fragment2": t.HTMLFragment2,
		},
	)

}

func (t *TextSimilarity) ToPostForm() *postform.PostForm {

	postForm := postform.New()

	postForm.Add("text1", &t.Text1)
	postForm.Add("html1", &t.HTML1)
	postForm.Add("html_fragment1", &t.HTMLFragment1)
	postForm.Add("text2", &t.Text2)
	postForm.Add("html2", &t.HTML2)
	postForm.Add("html_fragment2", &t.HTMLFragment2)
	postForm.Add("lang", t.Lang)
	postForm.Add("bow", t.Bow)

	return postForm
}
