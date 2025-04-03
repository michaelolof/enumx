# enumx

An enum is a closed set of possible variants. This library tries to provide an enum implementation in Golang that meets that requirement
<br>
<br>

## Installation
```$
go get -u github.com/michaelolof/enumx
```
<br>
<br>

## Examples
```go
package main

import (
    "github.com/michaelolof/enumx"
)

type Color = enumx.Enum[string, color]
Colors, set := enumx.Group[string, color]

var (
    Red     = set(color{name: "red", hex: "#FF0000"})
    Green   = set(color{name: "green", hex: "#00FF00"})
    Blue    = set(color{name: "blue", hex: "#0000FF"})
)

func main() {
    c := Green
    size := Colors.Len()
    isGood := isGoodGuy(c)
    hex := c.Item().hex
}

func isGoodGuy(c Color) {
    return Colors.Equal(c, Green)
}

type color {
    name string
    hex string
}

func (c color) Id() string {
    return c.name
}
```
<br>

For simple primitives that need to be treated like a closed set
```go
package main

import (
    "github.com/michaelolof/enumx"
)

type Color = enumx.Enum[string, enumsx.EnumString]
Colors, set := enumx.Group[string, enumx.EnumString]

var (
    Red     = set(enumx.EnumString("red"))
    Green   = set(enumx.EnumString("green"))
    Blue    = set(enumx.EnumString("blue"))
)
```