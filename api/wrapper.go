package api

import (
	"context"
	"fmt"

	"googlemaps.github.io/maps"
	"unchained.fetchdistance/model"
)

// GMapsClient wraps the Google Maps API Client
type GMapsClient struct {
	c *maps.Client
}

// NewGMapsClient creates a new GMapsClient
func NewGMapsClient(apiKey string) (*GMapsClient, error) {
	c, err := newMapsClient(apiKey)
	if err != nil {
		return nil, fmt.Errorf("NewGMapsClient: %s", err.Error())
	}

	return &GMapsClient{
		c: c,
	}, nil
}

// GetDistance takes an origin and destination and gets the distance from Google API
func (c *GMapsClient) GetDistance(ctx context.Context, origin, destination *model.Location) (*model.Distance, error) {
	o, err := origin.Get()
	if err != nil {
		return nil, fmt.Errorf("GetDistance: %s", err.Error())
	}

	d, err := origin.Get()
	if err != nil {
		return nil, fmt.Errorf("GetDistance: %s", err.Error())
	}

	req := &maps.DistanceMatrixRequest{
		Origins:      []string{o},
		Destinations: []string{d},
		Mode:         maps.TravelModeWalking,
		Units:        maps.UnitsImperial,
	}

	resp, err := c.c.DistanceMatrix(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetDistance: %s", err.Error())
	}

	if len(resp.Rows) == 0 {
		return nil, fmt.Errorf("GetDistance: No DistanceMatrixRows")
	}

	dist := &model.Distance{
		ID:       destination.ID,
		Duration: resp.Rows[0].Elements[0].Duration,
		Length:   resp.Rows[0].Elements[0].Distance.HumanReadable,
	}

	return dist, nil
}

// Must ensures that a GMapsClient has been created correctly
func Must(c *GMapsClient, err error) *GMapsClient {
	if err != nil || c.c == nil {
		return nil
	}

	return c
}

func newMapsClient(apiKey string) (*maps.Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("newMapsClient: API Key is empty")
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("newMapsClient: %s", err.Error())
	}

	return c, nil
}
