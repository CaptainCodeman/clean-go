package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/captaincodeman/clean-go/engine"
)

// translateQuery converts an application query spec into
// a mongodb specific query
func translateQuery(c *mgo.Collection, query *engine.Query) *mgo.Query {
	m := bson.M{}
	for _, filter := range query.Filters {
		switch filter.Condition {
		case engine.Equal:
			m[filter.Property] = filter.Value
		case engine.LessThan:
			m[filter.Property] = bson.M{"$lt": filter.Value}
		case engine.LessThanOrEqual:
			m[filter.Property] = bson.M{"$lte": filter.Value}
		case engine.GreaterThan:
			m[filter.Property] = bson.M{"$gt": filter.Value}
		case engine.GreaterThanOrEqual:
			m[filter.Property] = bson.M{"$gte": filter.Value}
		}
	}
	q := c.Find(m)

	for _, order := range query.Orders {
		switch order.Direction {
		case engine.Ascending:
			q = q.Sort(order.Property)
		case engine.Descending:
			q = q.Sort("-" + order.Property)
		}
	}

	if query.Offset > 0 {
		q = q.Skip(query.Offset)
	}

	if query.Limit > 0 {
		q = q.Limit(query.Limit)
	}

	return q
}
