package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	dandeliongo "github.com/henomis/dandelion-go"
	"github.com/henomis/dandelion-go/pkg/request"
)

const Token = "YOUR_API_TOKEN"

func main() {

	dandelionClient := dandeliongo.New(
		dandeliongo.DandelionEndpoint,
		Token,
		10*time.Second,
	)

	response, err := dandelionClient.SentimentAnalysis(
		&request.SentimentAnalysis{
			Lang: newString("en"),
			Text: "I really love your APIs",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if !response.IsSuccess() {
		log.Fatal(response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}

// Support methods

func newString(s string) *string {
	return &s
}
