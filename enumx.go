package enumx

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func New[T constraints.Ordered, V any](fallback *T) (group collection[T, V], setter func(k T, v V) T, getter func(val T) V) {
	g := collection[T, V]{
		enums:    map[T]V{},
		fallback: fallback,
	}

	return g, g.set, g.info
}

type collection[T constraints.Ordered, V any] struct {
	fallback *T
	enums    map[T]V
}

func (p *collection[T, V]) set(k T, v V) T {
	if _, found := p.enums[k]; found {
		panic(fmt.Sprintf("enum item '%v' already exist", k))
	}

	p.enums[k] = v
	return k
}

func (e *collection[T, V]) info(id T) V {
	if v, found := (e.enums)[id]; !found {
		panic("enum not found")
	} else {
		return v
	}
}

// Returns the enum item and a found state for a given enum id.
func (e *collection[T, V]) ById(id T) (T, bool) {
	if _, found := (e.enums)[id]; !found && e.fallback != nil {
		return *e.fallback, true
	} else if !found && e.fallback == nil {
		var t T
		return t, found
	} else {
		return id, true
	}
}

// Returns the enum item and a found state for a given enum id.
func (e *collection[T, V]) MustById(id T) T {
	if v, ok := e.ById(id); !ok {
		panic(fmt.Sprintf("Enum item '%v' not found", id))
	} else {
		return v
	}
}

// Iterates through the matches an item in the enum based on the predicate function and returns the enum and a found state
func (e *collection[T, V]) Find(p func(T) bool) (T, bool) {
	for v := range e.enums {
		if p(v) {
			return v, true
		}
	}
	if e.fallback != nil {
		return *e.fallback, true
	}

	var t T
	return t, false
}

// Iterates through the matches an item in the enum based on the predicate function. Panics if no match is found.
func (e *collection[T, V]) MustFind(p func(T) bool) T {
	if v, ok := e.Find(p); !ok {
		panic("Enum item not found")
	} else {
		return v
	}
}

// Len returns the length of the enum list on the base enum
func (e *collection[T, V]) Len() int {
	return len(e.enums)
}

// Returns a list of all the items in a given enum
func (e *collection[T, V]) Items() map[T]V {
	return e.enums
}

// Returns a list of all the items in a given enum
func (e *collection[T, V]) Values() []T {
	rtn := make([]T, 0, e.Len())
	for v := range e.enums {
		rtn = append(rtn, v)
	}
	return rtn
}
