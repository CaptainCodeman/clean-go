package domain

import "time"

var (
	// now returns the current UTC time
	// It is a replaceable function to allow for easy unit testing
	now = func() time.Time {
		return time.Now().UTC()
	}
)
