package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// ID represents the last used integer
	// id for any collection
	ID struct {
		Next int64 `bson:"n"`
	}
)

var (
	idCollection = "id"
)

// simple way of using integer IDs with MongoDB
func getNextSequence(s *mgo.Session, name string) int64 {
	c := s.DB("").C(idCollection)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	id := new(ID)
	c.Find(bson.M{"_id": name}).Apply(change, id)
	return id.Next
}
