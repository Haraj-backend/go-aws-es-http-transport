package main

import (
	"log"

	transport "github.com/Haraj-backend/go-aws-es-http-transport"
	"github.com/aws/aws-sdk-go/aws/session"
	es6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/riandyrn/go-env"
)

/*
	In this example we are trying to connect to elasticsearch
	endpoint by using default IAM user. We are using official
	elasticsearch client for Go.

	Set your endpoint on env variable named `ES_ENDPOINT` before
	running the example.
*/

const envKeyESEndpoint = "ES_ENDPOINT"

func main() {
	endpoint := env.GetString(envKeyESEndpoint)
	client, err := es6.NewClient(es6.Config{
		Addresses: []string{endpoint},
		Transport: transport.NewAWSESTransport(session.New()),
	})
	if err != nil {
		log.Fatalf("unable to initialize elasticsearch client due: %v", err)
	}
	resp, err := client.Ping()
	if err != nil {
		log.Fatalf("unable to ping elasticsearch endpoint due: %v", err)
	}
	defer resp.Body.Close()
	log.Printf("status: %v", resp.Status())
}
