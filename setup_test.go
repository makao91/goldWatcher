package main

import (
	"bytes"
	"goldwatcher/repository"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1654782060772,
	"tsj": 1654782056216,
	"date": "Jun 9th 2022, 09:40:56 am NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 1849,
		"xagPrice": 21.9115,
		"chgXau": -3.735,
		"chgXag": -0.1425,
		"pcXau": -0.2016,
		"pcXag": -0.6461,
		"xauClose": 1852.735,
		"xagClose": 22.054
	  }
	]
  }
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
