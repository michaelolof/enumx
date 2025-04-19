package enumx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumGroupLength(t *testing.T) {
	assert.Equal(t, Weekdays.Len(), 2, "weekdays length equal to 2")
	assert.Equal(t, Countries.Len(), 2, "countries length equal to 2")
}

func TestEnumListIteration(t *testing.T) {
	assert.Contains(t, Weekdays.Values(), Monday)
	assert.Contains(t, Weekdays.Values(), Tuesday)
	assert.NotContains(t, Weekdays.Values(), Nigeria)
	fmt.Println(Weekdays.Values())
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
