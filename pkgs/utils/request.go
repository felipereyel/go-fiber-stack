package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Headers map[string]string

func GetHTTPClient(baseUrl string, headers Headers) func(string, string, []byte) ([]byte, error) {
	return func(method string, path string, body []byte) ([]byte, error) {
		url := fmt.Sprintf("%s/%s", baseUrl, path)
		bodyReader := bytes.NewReader(body)

		req, err := http.NewRequest(method, url, bodyReader)
		if err != nil {
			return nil, errors.New("failed to create request: " + err.Error())
		}

		req.Header.Set("Content-Type", "application/json")
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		client := http.Client{Timeout: 30 * time.Second}

		res, err := client.Do(req)
		if err != nil {
			return nil, errors.New("failed to perform request: " + err.Error())
		}

		rawBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("could not read response body: " + err.Error())
		}

		if res.StatusCode != 200 {
			return nil, errors.New("not ok response from api: " + string(rawBody))
		}

		return rawBody, nil
	}
}
