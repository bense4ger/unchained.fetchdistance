package api

import "testing"

func TestNewMapsClientCreatesClient(t *testing.T) {
	c, err := NewMapsClient("api-key")
	if c == nil || err != nil {
		t.Errorf("Expected client to be created, it wasn't: %s", err.Error())
	}
}

func TestNewMapsClientReturnsError(t *testing.T) {
	c, err := NewMapsClient("")
	if c != nil || err == nil {
		t.Errorf("Expected an error, there wasn't one")
	}
}
