package model

// IncomingRequest encapsulates the body from the APIGateway Proxy Request
type IncomingRequest struct {
	ID          string `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
