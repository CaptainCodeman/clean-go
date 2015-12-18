package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Id struct {
		Next int64 `bson:"n"`
	}
)

var (
	idCollection = "id"
)

func getNextSequence(s *mgo.Session, name string) int64 {
	c := s.DB("").C(idCollection)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	id := new(Id)
	c.Find(bson.M{"_id": name}).Apply(change, id)
	return id.Next
}
