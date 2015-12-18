// +build appengine

package appengine

import (
	"google.golang.org/appengine/datastore"

	"github.com/captaincodeman/clean/engine"
)

func translateQuery(kind string, query *engine.Query) *datastore.Query {
	q := datastore.NewQuery(kind)

	for _, filter := range query.Filters {
		switch filter.Condition {
		case engine.Equal:
			q = q.Filter(filter.Property + " =", filter.Value)
		case engine.NotEqual:
			q = q.Filter(filter.Property + " >", filter.Value)	// TODO
		case engine.LessThan:
			q = q.Filter(filter.Property + " <", filter.Value)
		case engine.LessThanOrEqual:
			q = q.Filter(filter.Property + " <=", filter.Value)
		case engine.GreaterThan:
			q = q.Filter(filter.Property + " >", filter.Value)
		case engine.GreaterThanOrEqual:
			q = q.Filter(filter.Property + " >=", filter.Value)
		}
	}

	for _, order := range query.Orders {
		switch order.Direction {
		case engine.Ascending:
			q = q.Order(order.Property)
		case engine.Descending:
			q = q.Order("-" + order.Property)
		}
	}

	if query.Offset > 0 {
		q = q.Offset(query.Offset)
	}

	if query.Limit > 0 {
		q = q.Limit(query.Limit)
	}

	return q
}
