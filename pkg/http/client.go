package http

import (
	"backend-task/pkg/logger"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type FetchExternalAPIParam struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
	Params  map[string]interface{}
}

func FetchExternalAPI(param FetchExternalAPIParam) ([]byte, error) {
	parsedURL, err := url.Parse(param.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URL")
	}

	q := parsedURL.Query()
	for key, value := range param.Params {
		switch v := value.(type) {
		case string:
			q.Add(key, v)
		case int:
			q.Add(key, strconv.Itoa(v))
		default:
			logger.Log.Fatal(fmt.Sprintf("Unsupported param type for key %s", key))
		}
	}
	parsedURL.RawQuery = q.Encode()

	req, err := http.NewRequest(param.Method, parsedURL.String(), bytes.NewBuffer(param.Body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create http request")
	}

	for key, value := range param.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Errorf("received non-2xx response code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return responseBody, nil
}
