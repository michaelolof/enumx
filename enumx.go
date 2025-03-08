package enumx

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Identity[T constraints.Ordered] interface {
	Id() T
}

// Define an enum group
func Group[T constraints.Ordered, V Identity[T]](fallback *V) (enumGroup[T, V], func(val V) Enum[T, V]) {
	enums := map[T]enum[T, V]{}
	base := enumGroup[T, V]{
		enums:    enums,
		fallback: fallback,
	}

	return base, base.item
}

type enumGroup[T constraints.Ordered, V Identity[T]] struct {
	enums    map[T]enum[T, V]
	fallback *V
}

func (b *enumGroup[T, V]) item(val V) Enum[T, V] {
	e := enum[T, V]{
		base: b,
		item: val,
	}

	if _, alredyExist := b.enums[val.Id()]; alredyExist {
		panic(fmt.Sprintf("multiple assignments to same enum key %v", val.Id()))
	}

	(b.enums)[val.Id()] = e
	return &e
}

// Returns the enum item and a found state for a given enum id.
func (e *enumGroup[T, V]) ById(id T) (Enum[T, V], bool) {
	en, found := (e.enums)[id]
	if !found && e.fallback != nil {
		return &enum[T, V]{base: e, item: *e.fallback}, true
	}
	return &en, found
}

// Iterates through the matches an item in the enum based on the predicate function and returns the enum and a found state
func (e *enumGroup[T, V]) Find(predicate func(V) bool) (Enum[T, V], bool) {
	for _, e := range e.enums {
		if predicate(e.item) {
			return &e, true
		}
	}
	if e.fallback != nil {
		return &enum[T, V]{base: e, item: *e.fallback}, true
	}
	return nil, false
}

// Returns the enum item for a given Id. Panics if item is not found.
func (e *enumGroup[T, V]) TryById(id T) Enum[T, V] {
	if e, ok := (e.enums)[id]; ok {
		return &e
	}
	if e.fallback != nil {
		return &enum[T, V]{base: e, item: *e.fallback}
	}
	panic(fmt.Sprintf("Enum item '%v' not found", id))
}

// Iterates through the matches an item in the enum based on the predicate function. Panics if no match is found.
func (e *enumGroup[T, V]) TryFind(predicate func(V) bool) Enum[T, V] {
	for _, e := range e.enums {
		if predicate(e.item) {
			return &e
		}
	}
	if e.fallback != nil {
		return &enum[T, V]{base: e, item: *e.fallback}
	}
	panic("Enum item not found")
}

// Returns a hash map for all the items in a given enum
func (e *enumGroup[T, V]) Items() map[T]enum[T, V] {
	return e.enums
}

// Returns a hash map fo all the items in a given enum
func (e *enumGroup[T, V]) Values() map[T]V {
	values := make(map[T]V)
	for f, e := range e.enums {
		values[f] = e.item
	}
	return values
}

// Len returns the length of the enum list on the base enum
func (e *enumGroup[T, V]) Len() int {
	return len(e.enums)
}

// Checks if one enum item matches another item
func (e *enumGroup[T, V]) Equals(a Enum[T, V], b Enum[T, V]) bool {
	return a.Id() == b.Id()
}

// Len returns the length of the enum list on the base enum
func (e *enumGroup[T, V]) NotEquals(a Enum[T, V], b Enum[T, V]) bool {
	return a.Id() != b.Id()
}

type Enum[T constraints.Ordered, V Identity[T]] interface {
	Id() T
	Item() V
}

type enum[T constraints.Ordered, V Identity[T]] struct {
	base *enumGroup[T, V]
	// References the actual enum field of type V
	item V
}

// Returns the enum identifier
func (e enum[T, V]) Id() T {
	return e.item.Id()
}

// Returns the instance of the enum field
func (e enum[T, V]) Item() V {
	return e.item
}

func (e enum[T, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.item)
}

func (e *enum[T, V]) UnmarshalJSON(data []byte) error {
	var val V
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	e.item = val
	return nil
}
