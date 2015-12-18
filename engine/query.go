package engine

type Direction byte

const (
	Ascending Direction = 1 << iota
	Descending
)

type Condition byte

const (
	Equal Condition = 1 << iota
	NotEqual
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
)

type (
	Query struct {
		Name    string
		Offset  int
		Limit   int
		Filters []*Filter
		Orders  []*Order
	}

	QueryBuilder interface {
		Filter(property string, value interface{}) QueryBuilder
		Order(property string, direction Direction)
	}

	Filter struct {
		Property  string
		Condition Condition
		Value     interface{}
	}

	Order struct {
		Property  string
		Direction Direction
	}
)

func NewQuery(name string) *Query {
	return &Query{
		Name: name,
	}
}

func (q *Query) Filter(property string, condition Condition, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

func (q *Query) Order(property string, direction Direction) *Query {
	order := NewOrder(property, direction)
	q.Orders = append(q.Orders, order)
	return q
}

func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit
	return q
}

func NewFilter(property string, condition Condition, value interface{}) *Filter {
	return &Filter{
		Property:  property,
		Condition: condition,
		Value:     value,
	}
}

func NewOrder(property string, direction Direction) *Order {
	return &Order{
		Property:  property,
		Direction: direction,
	}
}
