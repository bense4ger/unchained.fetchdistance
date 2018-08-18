package debug

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"unchained.fetchdistance/api"
	"unchained.fetchdistance/model"
)

// RunLocal allows for local debugging
func RunLocal(origin, dest, apiKey string) {
	fmt.Println("Running Locally")

	o := &model.Location{
		ID:      "origin",
		LatLong: origin,
	}

	d := &model.Location{
		ID:      "destination",
		LatLong: dest,
	}

	c := api.Must(api.NewGMapsClient(apiKey))
	if c == nil {
		log.Fatal("Failed to create GMaps client")
	}

	dist, err := c.GetDistance(context.Background(), o, d)
	if err != nil {
		log.Warnf("runLocal: %s", err.Error())
	}

	log.Info(dist.String())
}
