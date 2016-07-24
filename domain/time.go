package domain

import "time"

var (
	// now returns the current UTC time
	// It is a replaceable function to allow for easy unit testing
	now = defaultNow
)

// set it back to this function to restore normal functionality
func defaultNow() time.Time {
	return time.Now().UTC()
}
