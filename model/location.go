package model

import (
	"fmt"
)

// Location encapsulates venue location details
type Location struct {
	Address string
	LatLong string
}

// Get returns the appropriate member.  If no member has data, an error is returned
func (l *Location) Get() (string, error) {
	switch {
	case len(l.Address) > 0:
		return l.Address, nil
	case len(l.LatLong) > 0:
		return l.LatLong, nil
	default:
		return "", fmt.Errorf("No location value")
	}
}
