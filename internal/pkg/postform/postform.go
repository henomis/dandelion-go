package postform

import (
	"fmt"
	"net/url"
)

type PostForm url.Values

func New() *PostForm {
	return &PostForm{}
}

func (p *PostForm) Add(key string, value *string) {
	if value != nil {
		(*url.Values)(p).Add(key, *value)
	}
}

func (p *PostForm) AddInt(key string, value *int) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%d", *value))
	}
}

func (p *PostForm) AddBool(key string, value *bool) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%t", *value))
	}
}

func (p *PostForm) AddFloat(key string, value *float64) {
	if value != nil {
		(*url.Values)(p).Add(key, fmt.Sprintf("%f", *value))
	}
}

func (p *PostForm) Encode() string {
	return (*url.Values)(p).Encode()
}
