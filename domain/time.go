package domain

import "time"

var (
	// Now returns the current UTC time
	//
	// It is a replaceable function to allow for easy unit testing
	Now = func() time.Time {
		return time.Now().UTC()
	}
)
