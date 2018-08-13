package main

import (
	"context"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	"unchained.fetchdistance/debug"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	env    string
	apiKey string
	origin = flag.String("origin", "", "The journey's origin.  This should be a lat long")
	dest   = flag.String("dest", "", "The journey's destination.  This should be an address")
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	switch env {
	case "prod":
		lambda.Start(handleRequest)
	case "dev":
		debug.RunLocal(*origin, *dest, apiKey)
	}
}

func init() {
	env = os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("No API Key specified")
	}

	if env == "dev" {
		flag.Parse()
	}
}
