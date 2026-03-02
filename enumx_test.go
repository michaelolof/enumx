package enumx

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumGroupLength(t *testing.T) {
	assert.Equal(t, Weekdays.Len(), 2, "weekdays length equal to 2")
	assert.Equal(t, Countries.Len(), 2, "countries length equal to 2")
}

func TestEnumListIteration(t *testing.T) {
	foundMonday := false
	foundTuesday := false
	for val := range Weekdays.Values() {
		if val == Monday {
			foundMonday = true
		}
		if val == Tuesday {
			foundTuesday = true
		}
	}
	assert.True(t, foundMonday)
	assert.True(t, foundTuesday)
}

type Weekday string

var Weekdays, s1, _ = New[Weekday, any](nil)

var (
	Monday  = s1("monday", nil)
	Tuesday = s1("tuesday", nil)
)

type Country string

var Countries, s2, g2 = New[Country, ci1](nil)
var (
	Nigeria      = s2("NG", ci1{iso3: "NGA", currency: "NGN"})
	UnitedStates = s2("US", ci1{iso3: "USA", currency: "USD"})
)

type ci1 struct {
	iso3     string
	currency string
}

func (c Country) Currency() string {
	return g2(c).currency
}

// ---------- additional tests for comprehensive coverage ----------

func TestNewWithFallback(t *testing.T) {
	fallback := "missing"
	col, set, _ := New[string, any](&fallback)
	set("a", nil)
	// existing
	id, ok := col.ById("a")
	assert.True(t, ok)
	assert.Equal(t, "a", id)
	// missing -> fallback
	id, ok = col.ById("b")
	assert.True(t, ok)
	assert.Equal(t, "missing", id)
}

func TestSetDuplicatePanics(t *testing.T) {
	_, set, _ := New[int, string](nil)
	set(1, "one")
	assert.PanicsWithValue(t, "enum item '1' already exist", func() {
		set(1, "uno")
	})
}

func TestInfoLookupAndPanic(t *testing.T) {
	_, set, info := New[string, int](nil)
	set("x", 42)
	assert.Equal(t, 42, info("x"))
	assert.PanicsWithValue(t, "enum not found", func() {
		info("nope")
	})
}

func TestByIdVariants(t *testing.T) {
	// no fallback
	col, set, _ := New[string, any](nil)
	set("k", nil)
	id, ok := col.ById("k")
	assert.True(t, ok)
	assert.Equal(t, "k", id)

	id, ok = col.ById("missing")
	assert.False(t, ok)
	assert.Equal(t, "", id)

	// with fallback pointer
	fb := "fb"
	col2, _, _ := New[string, any](&fb)
	id, ok = col2.ById("notthere")
	assert.True(t, ok)
	assert.Equal(t, "fb", id)
}

func TestMustById(t *testing.T) {
	col, set, _ := New[int, any](nil)
	set(10, nil)
	assert.Equal(t, 10, col.MustById(10))
	assert.PanicsWithValue(t, "Enum item '20' not found", func() {
		col.MustById(20)
	})
}

func TestFindAndMustFind(t *testing.T) {
	col, set, _ := New[int, bool](nil)
	set(1, true)
	set(2, false)
	n, ok := col.Find(func(i int) bool { return i%2 == 0 })
	assert.True(t, ok)
	assert.Equal(t, 2, n)

	// nothing matches, no fallback
	n, ok = col.Find(func(i int) bool { return i > 100 })
	assert.False(t, ok)
	assert.Equal(t, 0, n) // zero value

	assert.PanicsWithValue(t, "Enum item not found", func() {
		col.MustFind(func(i int) bool { return i < 0 })
	})
}

func TestLenEmptyAndItemsValues(t *testing.T) {
	col, set, _ := New[string, int](nil)
	assert.Equal(t, 0, col.Len())
	set("one", 1)
	set("two", 2)
	assert.Equal(t, 2, col.Len())

	iterCount := 0
	for key, val := range col.Items() {
		assert.Contains(t, []string{"one", "two"}, key)
		assert.Contains(t, []int{1, 2}, val)
		iterCount++
	}
	assert.Equal(t, 2, iterCount, "iterator should yield exactly 2 items")

	valsCount := 0
	for val := range col.Values() {
		assert.Contains(t, []string{"one", "two"}, val)
		valsCount++
	}
	assert.Equal(t, 2, valsCount)
}

func TestFallbackPointerMutation(t *testing.T) {
	value := "orig"
	col, _, _ := New[string, any](&value)
	id, ok := col.ById("missing")
	assert.True(t, ok)
	assert.Equal(t, "orig", id)

	value = "changed"
	id, ok = col.ById("missing")
	assert.True(t, ok)
	assert.Equal(t, "changed", id)
}

func TestVariousKeyAndValueTypes(t *testing.T) {
	// integer key, struct value
	type data struct{ X int }
	_, set, info := New[int, data](nil)
	set(5, data{X: 100})
	assert.Equal(t, data{X: 100}, info(5))

	// string key, float value
	_, set2, info2 := New[string, float64](nil)
	set2("pi", 3.14)
	assert.Equal(t, 3.14, info2("pi"))

	// custom alias types
	col3, set3, _ := New[UInt32, String](nil)
	set3(7, "seven")
	v, ok := col3.ById(7)
	assert.True(t, ok)
	assert.Equal(t, UInt32(7), v)
}

// While enumx collections are not safe for concurrent writes, this example
// illustrates that attempting to write from multiple goroutines can panic
// (or at least trigger the race detector). We include it as documentation
// rather than a required guarantee.
func TestConcurrentWriteUnsafe(t *testing.T) {
	_, set, _ := New[int, any](nil)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		set(1, nil)
	}()
	go func() {
		defer wg.Done()
		set(2, nil)
	}()
	wg.Wait()
	// if the runtime panics due to concurrent map write the test will fail,
	// reinforcing that the API is not goroutine-safe.
}

// Example tests act as documentation and are executed by `go test`.
func ExampleNew() {
	_, set, get := New[string, int](nil)
	set("one", 1)
	fmt.Println(get("one"))
	// Output: 1
}
