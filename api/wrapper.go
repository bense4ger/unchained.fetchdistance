package api

import (
	"fmt"

	"googlemaps.github.io/maps"
)

// NewMapsClient creates a new Google Maps API Client
func NewMapsClient(apiKey string) (*maps.Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("NewMapsClient: API Key is empty")
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("NewMapsClient: %s", err.Error())
	}

	return c, nil
}
