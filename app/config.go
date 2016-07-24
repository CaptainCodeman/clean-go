// +build !appengine

package main

type (
	// Config is an example of provider-specific configuration
	// in this case it's for the standalone version only to set
	// the mongodb database connection string
	Config struct {
		MongoURL string
	}
)

var (
	config *Config
)

func init() {
	// this would likely be loaded from flags or a conf file
	config = &Config{
		MongoURL: "mongodb://localhost/clean",
	}
}
