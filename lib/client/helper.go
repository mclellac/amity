package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func buildRequest(method string, url string, entity interface{}) (*http.Request, error) {
	body, err := encodeEntity(entity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	req.Header.Set("content-type", "application/json")

	return req, err
}

func encodeEntity(entity interface{}) (io.Reader, error) {
	if entity == nil {
		return nil, nil
	}

	buff, err := json.Marshal(entity)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(buff), nil
}

func processResponse(r *http.Response, expectedStatus int) error {
	if r.StatusCode != expectedStatus {
		return errors.New("HTTP: " + r.Status)
	}

	return nil
}

func processResponseEntity(r *http.Response, entity interface{}, expectedStatus int) error {
	if err := processResponse(r, expectedStatus); err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, entity); err != nil {
		return err
	}

	return nil
}

func makeRequest(method string, url string, entity interface{}) (*http.Response, error) {
	req, err := buildRequest(method, url, entity)

	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
