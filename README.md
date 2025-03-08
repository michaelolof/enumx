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

type Colors_ = enumx.Enum[string, colors]
Colors, set := enumx.Group[string, colors]

var (
    Red     = set(colors{name: "red", hex: "#FF0000"})
    Green   = set(colors{name: "green", hex: "#00FF00"})
    Blue    = set(colors{name: "blue", hex: "#0000FF"})
)

func main() {
    c := Green
    size := Colors.Len()
    isGood := isGoodGuy(c)
    hex := c.Item().hex
}

func isGoodGuy(c Colors_) {
    return Colors.Equal(c, Green)
}

type colors {
    name string
    hex string
}

func (c colors) Id() string {
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

type Colors_ = enumx.Enum[string, enumsx.EnumString]
Colors, set := enumx.Group[string, enumx.EnumString]

var (
    Red     = set(enumx.EnumString("red"))
    Green   = set(enumx.EnumString("green"))
    Blue    = set(enumx.EnumString("blue"))
)
```