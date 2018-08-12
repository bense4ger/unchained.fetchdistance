package api

import "testing"

func TestNewMapsClientCreatesClient(t *testing.T) {
	c, err := newMapsClient("api-key")
	if c == nil || err != nil {
		t.Errorf("Expected client to be created, it wasn't: %s", err.Error())
	}
}

func TestNewMapsClientReturnsError(t *testing.T) {
	c, err := newMapsClient("")
	if c != nil || err == nil {
		t.Error("Expected an error, there wasn't one")
	}
}

func TestNewGMapsClient(t *testing.T) {
	_, err := NewGMapsClient("api-key")
	if err != nil {
		t.Errorf("Expected error to be nil, but it wasn't: %s", err.Error())
	}
}

func TestMustReturnsClient(t *testing.T) {
	c := Must(NewGMapsClient("api-key"))
	if c == nil {
		t.Error("Expected a client to be created, but it wasn't")
	}
}

func TestMustReturnsNil(t *testing.T) {
	c := Must(NewGMapsClient(""))
	if c != nil {
		t.Errorf("Expected no client to be created, but it was")
	}
}
