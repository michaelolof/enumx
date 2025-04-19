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

## Usage

Normal enums* in Golang
```go
package main

type Color int

const (
    Red Color = 1 + iota
    Green
    Blue
)
```
<br>

Enumx enums
```go
package main

import "github.com/michaelolof/enumx"

type Color int

var Colors, set, _ = enumx.New[Color, any](nil)
var (
    Red Color = set(1, nil)
    Green = set(2, nil)
    Blue = set(3, nil)
)
```
`Red`, `Green` and `Blue` are still variables of type `Color` each holding a value of `1`, `2` and `3` respectively except we now have a new enum collection variable `Colors` which holds the list of colors defined by the `set` function
<br/>
The library works by using a dictionary under the hood. The `set` function adds the variant value to the collection and returns the variant type. `set` will panic if you try adding a new variant with a key that already exists.

<br />
<br />
Enumx enums can also be used for detailed enums

```go
package main

import "github.com/michaelolof/enumx"

type KnownCountry string

var KnownCountries, set, get = enumx.New[KnownCountry, countryInfo](nil)
var (
    Nigeria         = set("NG", countryInfo{iso3: "NGA", dial: "+234", name: "Nigeria"})
    Ghana           = set("GH", countryInfo{iso3: "GHA", dial: "+233", name: "Ghana"})
    UnitedStates    = set("US", countryInfo{iso3: "USA", dial: "+1", name: "United States of America"})
)

func (k KnownCountry) Iso3() string {
    return get(k).iso3
}

func (k KnownCountry) Dial() string {
    return get(k).dial
}

func (k KnownCountry) Name() string {
    return get(k).name
}

type countryInfo struct {
    iso3 string
    dial string
    name string
}

func main() {
    fmt.Println(Nigeria.Name()) // Nigeria
    fmt.Println(Ghana.Dial()) // +233
    fmt.Println(UnitedStates.Iso3()) // USA
    fmt.Println(KnownCountries.Len()) // 3
    fmt.Println(KnownCountries.Values()) // [GH NG US]
}
```