package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type WikiSearch struct {
	// The following three parameters are mutually exclusive
	Text    string  `json:"text" validate:"-"`
	Lang    string  `json:"lang" validate:"oneof=de en es fr it pt ru"`
	Limit   *int    `json:"limit" validate:"omitempty,gte=1 lte=50"`
	Offset  *int    `json:"offset" validate:"omitempty,gte=0"`
	Query   *string `json:"query" validate:"omitempty,oneof=full prefix"`
	Include *string `json:"include" validate:"-"`
}

func (w *WikiSearch) Validate() error {
	validate := validator.New()

	if err := validate.Struct(w); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return nil
}

func (w *WikiSearch) ToPostForm() *postform.PostForm {

	postForm := postform.New()

	postForm.Add("text", &w.Text)
	postForm.Add("lang", &w.Lang)
	postForm.AddInt("limit", w.Limit)
	postForm.AddInt("offset", w.Offset)
	postForm.Add("query", w.Query)
	postForm.Add("include", w.Include)

	return postForm
}
