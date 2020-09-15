package blazegraph

import (
	"net/http"
	"github.com/cirss/geist/sparql"
)

var DefaultEndpoint = "http://127.0.0.1:9999/blazegraph/sparql"

type Client struct {
	sparql.Client
}

func NewClient() *Client {
	bc := new(Client)
	bc.HttpClient = &http.Client{}
	bc.Endpoint = DefaultEndpoint
	return bc
}


