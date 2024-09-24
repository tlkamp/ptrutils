# ptrutils

A library for converting to and from pointers using generics.

## Examples

```go
package main

import (
    "fmt"

    "github.com/tlkamp/ptrutils"
)

func main() {
    var i *int // nil
    example := "example"

    // From pointers
    fmt.Printf("%+v\n", ptrutils.FromPointer(i))
    fmt.Printf("%+v\n", ptrutils.FromPointer(&example))

    // To Pointer
    fmt.Printf("%+v\n", ptrutils.ToPointer(example))
}
```
