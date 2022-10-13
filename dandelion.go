package dandeliongo

import (
	"fmt"
	"time"

	"github.com/henomis/dandelion-go/internal/pkg/httpclient"
	"github.com/henomis/dandelion-go/pkg/request"
	"github.com/henomis/dandelion-go/pkg/response"
)

const (
	DandelionEndpoint     = "https://api.dandelion.eu"
	entityExtractionPath  = "/datatxt/nex/v1/"
	textSimilarityPath    = "/datatxt/sim/v1/"
	languageDetectionPath = "/datatxt/li/v1/"
	sentimentAnalysisPath = "/datatxt/sent/v1/"
	wikiSearchPath        = "/datagraph/wikisearch/v1/"
)

type DandelionClient struct {
	httpClient *httpclient.HttpClient
}

func New(endpoint string, token string, timeout time.Duration) *DandelionClient {
	return &DandelionClient{
		httpClient: httpclient.New(endpoint, token, timeout),
	}
}

func (d *DandelionClient) EntityExtraction(
	entityExtractionRequest *request.EntityExtraction,
) (*response.EntityExtraction, error) {

	entityExtractionResponse := &response.EntityExtraction{}

	err := d.httpClient.Request(
		entityExtractionPath,
		entityExtractionRequest,
		entityExtractionResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return entityExtractionResponse, nil
}

func (d *DandelionClient) TextSimilarity(
	textSimilarityRequest *request.TextSimilarity,
) (*response.TextSimilarity, error) {

	textSimilarityResponse := &response.TextSimilarity{}

	err := d.httpClient.Request(
		textSimilarityPath,
		textSimilarityRequest,
		textSimilarityResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return textSimilarityResponse, nil
}

func (d *DandelionClient) LanguageDetection(
	languageDetectionRequest *request.LanguageDetection,
) (*response.LanguageDetection, error) {

	languageDetectionResponse := &response.LanguageDetection{}

	err := d.httpClient.Request(
		languageDetectionPath,
		languageDetectionRequest,
		languageDetectionResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return languageDetectionResponse, nil
}

func (d *DandelionClient) SentimentAnalysis(
	sentimentAnalysisRequest *request.SentimentAnalysis,
) (*response.SentimentAnalysis, error) {

	sentimentAnalysisResponse := &response.SentimentAnalysis{}

	err := d.httpClient.Request(
		sentimentAnalysisPath,
		sentimentAnalysisRequest,
		sentimentAnalysisResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return sentimentAnalysisResponse, nil
}

func (d *DandelionClient) WikiSearch(
	wikiSearchRequest *request.WikiSearch,
) (*response.WikiSearch, error) {

	wikiSearchResponse := &response.WikiSearch{}

	err := d.httpClient.Request(
		wikiSearchPath,
		wikiSearchRequest,
		wikiSearchResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	return wikiSearchResponse, nil
}
