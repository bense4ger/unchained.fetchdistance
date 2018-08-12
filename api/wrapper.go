package api

import (
	"fmt"

	"googlemaps.github.io/maps"
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
