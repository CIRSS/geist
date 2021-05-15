package geist

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RestClient struct {
	HttpClient *http.Client
	Logger     *log.Logger
	Endpoint   string
}

func NewRestClient(endpoint string) *RestClient {
	rc := new(RestClient)
	rc.HttpClient = &http.Client{}
	rc.Endpoint = endpoint
	return rc
}

func (sc *SparqlClient) GetRequest(url string, contentType string,
	acceptType string) (responseBody []byte, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()

	// create the http request using the provided body
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)

	// perform the request and obtain the response
	response, err := sc.HttpClient.Do(request.WithContext(ctx))
	if err != nil {
		return
	}

	// read the response
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	return
}

func (sc *SparqlClient) PostRequest(url string, contentType string, acceptType string,
	requestBody []byte) (responseBody []byte, err error) {

	// ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	// defer cancel()

	// create the http requeest using the provided body
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)

	if sc.Logger != nil {
		sc.Logger.Print(request)
	}

	// perform the request and obtain the response
	response, err := sc.HttpClient.Do(request)
	if err != nil {
		return
	}

	// read the response
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	return
}

func (sc *SparqlClient) DeleteRequest(url string) (
	responseBody []byte, err error) {

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	// perform the request and obtain the response
	response, err := sc.HttpClient.Do(request)
	if err != nil {
		return
	}

	// read the response
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	return
}
