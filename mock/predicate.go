package mock

import "fmt"

// Predicate defines an interface of functions for all mock predicates
type Predicate interface {
	Query() string
	Type() int
}

const (
	_ = iota
	predicateURL
	predicatePath
	predicateMethod
	predicateQuery
	predicateBody
	predicateBodyPart
	predicateHeader
	predicateBasicAuth
	predicateBearerAuth
)

type basePredicate struct {
	typ int
}

func (p basePredicate) Type() int {
	return p.typ
}

func newBasePredicate(typ int) basePredicate {
	return basePredicate{
		typ: typ,
	}
}

type keyValuePredicate struct {
	basePredicate
	key   string
	value interface{}
}

func newKeyValuePredicate(typ int, key string, value interface{}) *keyValuePredicate {
	return &keyValuePredicate{
		key:           key,
		value:         value,
		basePredicate: newBasePredicate(typ),
	}
}
func (p *keyValuePredicate) Query() string {
	return fmt.Sprintf("%d-%s-%v", p.typ, p.key, p.value)
}

type valuePredicate struct {
	basePredicate
	value interface{}
}

func newValuePredicate(typ int, value interface{}) *valuePredicate {
	return &valuePredicate{
		value:         value,
		basePredicate: newBasePredicate(typ),
	}
}

func (p *valuePredicate) Query() string {
	return fmt.Sprintf("%d,%s", p.typ, p.value)
}
