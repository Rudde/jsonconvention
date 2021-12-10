# JsonConvention

A package that allow automatic marshaling and unmarshaling of json data based on a convention defined by an input function.

The package is a copy of the contents of the encoding/json package in the offical go compiler and adjusted to give this flexibility. For this reason this package is not designed to be used without being given a convention function, if you have a mixed usecase, just use the offical encoding/json for places where this package ins't needed.

The package can be used exctatly like the offical one, with one more parameter `convention func(string) string` which is the function who will take the name of the struct field and output it in the desierd convention.

## Installation

```go
import "github.com/rudde/jsonconvention"
```

## Usage

```go
package main

import (
    "strings"
    "bytes"

    "github.com/rudde/jsonconvention"
)

type FooBar struct {
    Foo string
    Bar string
}

jsonStr := []byte(`{"Foo": "foo", "Bar": "bar"}`)

fooBar := FooBar{}

_ = jsonconvention.Unmarshal(jsonStr, &fooBar, strings.ToLower)

jsonDta, _ := jsonconvention.Marshal(fooBar, strings.ToLower)

d := jsonconvention.NewDecoder(bytes.NewBuffer(jsonDta))
_ = d.Decode(&fooBar, strings.ToLower)
```

Note that if you have actually defined a name with the json tag, the convention function will bem overwritten with what you have defined in the json tag.