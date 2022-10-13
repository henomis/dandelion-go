package request

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type EntityExtraction struct {
	// The following three parameters are mutually exclusive
	Text          string   `json:"text,omitempty" validate:"-"`
	HTML          string   `json:"html,omitempty" validate:"-"`
	HTMLFragment  string   `json:"html_fragment,omitempty" validate:"-"`
	Lang          *string  `json:"lang,omitempty" validate:"omitempty,oneof=de en es fr it pt ru auto"`
	TopEntities   *int     `json:"top_entities,omitempty" validate:"omitempty, gte=0"`
	MinConfidence *float64 `json:"min_confidence,omitempty" validate:"omitempty, gte=0.0, lte=1.0"`
	MinLength     *int     `json:"min_length,omitempty" validate:"omitempty, gte=2"`
	SocialHashtag *bool    `json:"social.hashtag,omitempty" validate:"-"`
	SocialMention *bool    `json:"social.mention,omitempty" validate:"-"`
	Include       *string  `json:"include,omitempty" validate:"-"`
	ExtraTypes    *string  `json:"extra_types,omitempty" validate:"-"`
	Country       *string  `json:"country,omitempty" validate:"omitempty,oneof=AD AE AM AO AQ AR AU BB BR BS BY CA CH CL CN CX DE FR GB HU IT JP KR MX NZ PG PL RE SE SG US YT ZW"`
	Epsilon       *float64 `json:"epsilon,omitempty" validate:"omitempty, gte=0.0, lte=0.5"`
}

type Social struct {
	Hashtag *bool `json:"hashtag,omitempty" validate:"-"`
	Mention *bool `json:"mention,omitempty" validate:"-"`
}

func (e *EntityExtraction) Validate() error {
	validate := validator.New()

	if err := validate.Struct(e); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"text":          e.Text,
			"html":          e.HTML,
			"html_fragment": e.HTMLFragment,
		},
	)

}

func (e *EntityExtraction) ToPostForm() *postform.PostForm {

	postForm := postform.New()

	postForm.Add("text", &e.Text)
	postForm.Add("html", &e.HTML)
	postForm.Add("html_fragment", &e.HTMLFragment)
	postForm.Add("lang", e.Lang)
	postForm.AddInt("top_entities", e.TopEntities)
	postForm.AddFloat("min_confidence", e.MinConfidence)
	postForm.AddInt("min_length", e.MinLength)
	postForm.AddBool("social.hashtag", e.SocialHashtag)
	postForm.AddBool("social.mention", e.SocialMention)
	postForm.Add("include", e.Include)
	postForm.Add("extra_types", e.ExtraTypes)
	postForm.Add("country", e.Country)

	return postForm
}
