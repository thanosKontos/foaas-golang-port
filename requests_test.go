package main

import (
	"net/http"
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestHtmlResponses(t *testing.T) {
	var cases = map[string]string{
		"http://127.0.0.1:8000/version": "<h1>Version 0.0.1</h1>",
		"http://127.0.0.1:8000/zero/John%20Doe": "<h1>Zero, thats the number of fucks I give.</h1>",
		"http://127.0.0.1:8000/pulp/Greek/John%20Doe": "<h1>Greek, motherfucker, do you speak it?</h1>",
	}

	for url, expectedString := range cases {
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		statusCode, body := getResponseCodeAndBody(req)

		assert.Equal(t, http.StatusOK, statusCode)
		assert.Contains(t, body, expectedString)
    }
}

func TestJsonResponses(t *testing.T) {
	var cases = map[string]string{
		"http://127.0.0.1:8000/version": "{\"message\":\"Version 0.0.1\",\"subtitle\":\"FOAAS\"}",
		"http://127.0.0.1:8000/zero/John%20Doe": "{\"message\":\"Zero, thats the number of fucks I give.\",\"subtitle\":\"- John Doe\"}",
		"http://127.0.0.1:8000/pulp/Greek/John%20Doe": "{\"message\":\"Greek, motherfucker, do you speak it?\",\"subtitle\":\"- John Doe\"}",
	}

	for url, expectedString := range cases {
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Content-Type", `application/json`)
		assert.NoError(t, err)
	
		statusCode, body := getResponseCodeAndBody(req)

		assert.Equal(t, http.StatusOK, statusCode)
		assert.Contains(t, body, expectedString)
	}
}

func TestNonExistingResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/nonexisting", nil)
	assert.NoError(t, err)

	statusCode, body := getResponseCodeAndBody(req)

	assert.Equal(t, http.StatusOK, statusCode)
	assert.Contains(t, body, "<h1>622 - All The Fucks</h1>")
}

func TestMethodNotAllowedResponse(t *testing.T) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/version", nil)
	assert.NoError(t, err)

	statusCode, body := getResponseCodeAndBody(req)

	assert.Equal(t, http.StatusOK, statusCode)
	assert.Contains(t, body, "<h1>622 - All The Fucks</h1>")
}

func getResponseCodeAndBody(req *http.Request) (int, string) {
	client := &http.Client{}
	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return resp.StatusCode, bodyString
}
