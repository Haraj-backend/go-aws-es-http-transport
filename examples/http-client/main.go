package main

import (
	"io/ioutil"
	"log"
	"net/http"

	transport "github.com/Haraj-backend/go-aws-es-http-transport"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/riandyrn/go-env"
)

/*
	In this example we are trying to connect to elasticsearch
	endpoint by using default IAM user. We are using plain http
	client to connect with the endpoint.

	Set env variable named `ES_ENDPOINT` before running the
	example. The value of `ES_ENDPOINT` is endpoint of your
	elasticsearch domain.

	Example value of `ES_ENDPOINT`:
	https://this-is-sample-domain.eu-west-1.es.amazonaws.com
*/

const envKeyESEndpoint = "ES_ENDPOINT"

func main() {
	endpoint := env.GetString(envKeyESEndpoint)
	client := &http.Client{
		Transport: transport.NewAWSESTransport(session.New()),
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		log.Fatalf("unable to connect to ES cluster due: %v", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	log.Printf("status code: %v, body: %s", resp.Status, b)
}
