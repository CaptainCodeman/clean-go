package engine

// Direction represents a query sort direction
type Direction byte

const (
	// Ascending means going up, A-Z
	Ascending Direction = 1 << iota

	// Descending means reverse order, Z-A
	Descending
)

// Condition represents a filter comparison operation
// between a field and a value
type Condition byte

const (
	// Equal if it should be the same
	Equal Condition = 1 << iota

	// LessThan if it should be smaller
	LessThan

	// LessThanOrEqual if it should be smaller or equal
	LessThanOrEqual

	// GreaterThan if it should be larger
	GreaterThan

	// GreaterThanOrEqual if it should be equal or greater than
	GreaterThanOrEqual
)

type (
	// Query represents a query specification for filtering
	// sorting, paging and limiting the data requested
	Query struct {
		Name    string
		Offset  int
		Limit   int
		Filters []*Filter
		Orders  []*Order
	}

	// QueryBuilder helps with query creation
	QueryBuilder interface {
		Filter(property string, value interface{}) QueryBuilder
		Order(property string, direction Direction)
	}

	// Filter represents a filter operation on a single field
	Filter struct {
		Property  string
		Condition Condition
		Value     interface{}
	}

	// Order represents a sort operation on a single field
	Order struct {
		Property  string
		Direction Direction
	}
)

// NewQuery creates a new database query spec. The name is what
// the storage system should use to identify the types, usually
// a table or collection name.
func NewQuery(name string) *Query {
	return &Query{
		Name: name,
	}
}

// Filter adds a filter to the query
func (q *Query) Filter(property string, condition Condition, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

// Order adds a sort order to the query
func (q *Query) Order(property string, direction Direction) *Query {
	order := NewOrder(property, direction)
	q.Orders = append(q.Orders, order)
	return q
}

// Slice adds a slice operation to the query
func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit
	return q
}

// NewFilter creates a new property filter
func NewFilter(property string, condition Condition, value interface{}) *Filter {
	return &Filter{
		Property:  property,
		Condition: condition,
		Value:     value,
	}
}

// NewOrder creates a new query order
func NewOrder(property string, direction Direction) *Order {
	return &Order{
		Property:  property,
		Direction: direction,
	}
}
