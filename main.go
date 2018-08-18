package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"bense4ger/unchained.fetchdistance/api"
	"bense4ger/unchained.fetchdistance/debug"
	"bense4ger/unchained.fetchdistance/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	env    string
	apiKey string
	origin = flag.String("origin", "", "The journey's origin.  This should be a lat long")
	dest   = flag.String("dest", "", "The journey's destination.  This should be an address")
)

func notOk(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	h := make(map[string]string)
	h["content-type"] = "application/json"

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       fmt.Sprintf("{\"message\":\"%s\"}", err.Error()),
		Headers:    h,
	}, err
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := api.Must(api.NewGMapsClient(apiKey))
	if c == nil {
		log.Warn("Failed to create API Client")
		return notOk(http.StatusInternalServerError, fmt.Errorf("Failed to create API Client"))
	}

	if req.Body == "" {
		log.Warn("Empty body passed")
		return notOk(http.StatusBadRequest, fmt.Errorf("Empty body passed"))
	}

	bytes := []byte(req.Body)
	ir := &model.IncomingRequest{}

	err := json.Unmarshal(bytes, ir)
	if err != nil {
		log.Warnf("Error unmarshalling incoming request: %s", err.Error())
		return notOk(http.StatusInternalServerError, fmt.Errorf("Unable to parse request body"))
	}

	o := &model.Location{
		LatLong: ir.Origin,
	}

	d := &model.Location{
		ID:      ir.ID,
		LatLong: ir.Destination,
	}

	dist, err := c.GetDistance(context.Background(), o, d)
	if err != nil {
		log.Warnf("Error getting distance: %s", err.Error)
		return notOk(http.StatusInternalServerError, fmt.Errorf("Unable to get distance"))
	}

	rb, err := json.Marshal(dist)
	if err != nil {
		log.Warnf("Error marshalling distance struct: %s", err.Error())
		return notOk(http.StatusInternalServerError, fmt.Errorf("Unable to get distance"))
	}

	h := make(map[string]string)
	h["content-type"] = "application/json"

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    h,
		Body:       string(rb),
	}, nil
}

func main() {
	switch env {
	case "prod":
		lambda.Start(handleRequest)
	case "dev":
		if *origin == "" || *dest == "" {
			log.Fatal("Origin and destination must not be empty")
		}
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
