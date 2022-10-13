# Dandelion SDK for Go


[![Build Status](https://github.com/henomis/dandelion-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/dandelion-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/dandelion-go?status.svg)](https://godoc.org/github.com/henomis/dandelion-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/dandelion-go)](https://goreportcard.com/report/github.com/henomis/dandelion-go) [![GitHub release](https://img.shields.io/github/release/henomis/dandelion-go.svg)](https://github.com/henomis/dandelion-go/releases)

This is Danelion's **unofficial** Go client, designed to enable you to use Dandelion's services easily from your own applications.

## Danelion

Danelion is a cloud-based text analytics service that through APIs allows you extract informations from a text content.


## Supported API versions

|                        |    |
|------------------------|----|
| Entity extraction      | v1 |
| Text similarity        | v1 |
| Language detection	 | v1 |
| Sentiment analysis     | v1 |
| Wikisearch 		     | v1 |


## Getting started

### Installation

You can load dandelion-go into your project by using:
```
go get github.com/henomis/dandelion-go
```


### Configuration

The only thing you need to start using Dandelio's APIs is the `Token`. 


### Usage

Please refer to the [examples folder](examples/) to see how to use the SDK.

Here below a simple usage example:

```go
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

	response, err := dandelionClient.EntityExtraction(
		&request.EntityExtraction{
			Lang:    newString("en"),
			Include: newString("types,abstract,categories,lod"),
			Text:    "The doctor says an apple is better than an orange",
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
```