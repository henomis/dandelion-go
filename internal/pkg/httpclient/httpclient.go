package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/henomis/dandelion-go/internal/pkg/postform"
)

type HttpClient struct {
	httpClient *http.Client
	baseURL    string
	token      string
	keepAlive  bool //bonus point
}

const dandelionTimeFormat = "2006-01-02 15:04:05 -0700"

type RequestData interface {
	ToPostForm() *postform.PostForm
}

type ResponseData interface {
	Decode(body io.ReadCloser) error
	SetHeaders(units, unitsLeft float64, unitsReset time.Time)
}

func New(baseURL string, token string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		token:   token,
	}
}

func (h *HttpClient) Request(
	path string,
	requestData RequestData,
	responseData ResponseData,
) error {

	postFormData := requestData.ToPostForm()
	postFormData.Add("token", &h.token)

	request, err := http.NewRequest("POST", h.baseURL+path, strings.NewReader(postFormData.Encode()))
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return err
	}

	err = responseData.Decode(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(response.Header)

	responseData.SetHeaders(
		getHeaderAsFloat(response.Header, "X-DL-units"),
		getHeaderAsFloat(response.Header, "X-DL-units-left"),
		getHeaderAsTime(response.Header, "X-DL-units-reset"),
	)

	return nil
}

// Support methods

func getHeaderAsFloat(header http.Header, key string) float64 {

	value, err := strconv.ParseFloat(header.Get(key), 64)
	if err != nil {
		return 0
	}

	return value
}

func getHeaderAsTime(header http.Header, key string) time.Time {

	value, err := time.Parse(dandelionTimeFormat, header.Get(key))
	if err != nil {
		return time.Time{}
	}

	return value
}
