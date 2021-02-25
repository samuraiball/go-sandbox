package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	MOCK_URL  = "localhost:8081"
	PING_PATH = "/ping"
)

func TestName(t *testing.T) {
	defer gock.Off()

	gock.New(MOCK_URL).
		Get(PING_PATH).
		MatchHeader("x-api-version", "1.[0-9]+").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})
	//res, err := http.Get("http://" + MOCK_URL + PING_PATH)
	//handleError(err)
	req, err := http.NewRequest(http.MethodGet, "http://"+MOCK_URL+PING_PATH, nil)
	handleError(err)

	req.Header = map[string][]string{"x-api-version": {"1.0"}}

	client := &http.Client{}
	res, err := client.Do(req)

	handleError(err)

	bodyByte, err := io.ReadAll(res.Body)
	handleError(err)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, `{"foo":"bar"}`, string(bodyByte)[:13])
	assert.Equal(t, gock.IsDone(), true)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
