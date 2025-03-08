package enumx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeAliasedEnum(t *testing.T) {
	type Colors_ = Enum[string, EnumString]
	var Colors, set = Group[string, EnumString](nil)

	var (
		Red   = set(EnumString("red"))
		Green = set(EnumString("green"))
		Blue  = set(EnumString("blue"))
	)

	tester := func(c ...Colors_) {
		assert.Equal(t, Colors.Len(), len(c))
	}

	tester(Red, Green, Blue)
	assert.Equal(t, Red.Id(), "red")
	assert.True(t, Colors.Equals(Green, Green))
	assert.True(t, Colors.NotEquals(Blue, Red))
}

type color struct {
	Name string
	hex  string
	rgb  [3]int
}

func (c color) Id() string {
	return c.Name
}

func TestStructEnums(t *testing.T) {

	Colors, set := Group[string, color](nil)

	var (
		Red   = set(color{Name: "red", hex: "#FF0000", rgb: [3]int{255, 0, 0}})
		Green = set(color{Name: "green", hex: "#00FF00", rgb: [3]int{0, 255, 0}})
		Blue  = set(color{Name: "blue", hex: "#0000FF", rgb: [3]int{0, 0, 255}})
	)

	assert.Equal(t, Colors.Len(), 3)
	assert.Equal(t, Red.Item().hex, "#FF0000")
	assert.Equal(t, Green.Item().hex, "#00FF00")
	assert.Equal(t, Blue.Item().hex, "#0000FF")
}

func TestEnumIterate(t *testing.T) {
	type Colors_ = Enum[string, color]
	Colors, set := Group[string, color](nil)

	var (
		Red   = set(color{Name: "red", hex: "#FF0000", rgb: [3]int{255, 0, 0}})
		Green = set(color{Name: "green", hex: "#00FF00", rgb: [3]int{0, 255, 0}})
		Blue  = set(color{Name: "blue", hex: "#0000FF", rgb: [3]int{0, 0, 255}})
	)

	checkColor := func(c Colors_) {
		fmt.Println("Color:", c.Item().Name, "-", c.Item().hex)
	}

	for _, k := range Colors.Items() {
		checkColor(k)
	}

	assert.Equal(t, Colors.Len(), [3]Colors_{Red, Green, Blue})
}
