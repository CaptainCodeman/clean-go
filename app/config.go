// +build !appengine

package main

type (
	Config struct {
		MongoURL	string
	}
)

var (
	config *Config
)

func init() {
	config = &Config{
		MongoURL: "mongodb://localhost/clean",
	}
}