# leblad

[![Build Status](https://github.com/omdxp/leblad/workflows/Test%20CI/badge.svg)](https://github.com/omdxp/leblad/actions?query=branch%3Amain)

A go module providing a list of Algerian administrative areas with many useful APIs, based on [dzcode-io/leblad](https://github.com/dzcode-io/leblad)

## Installation

```bash
go get -u github.com/omdxp/leblad
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/omdxp/leblad"
)

func main() {
    l := leblad.New()

    // Get all wilayas
    wilayas, err := l.GetWilayaList()
    if err != nil {
        panic(err)
    }

    fmt.Println(wilayas)

    // Get only wilayas names
    wilayas, err = l.GetWilayaList("name")
    if err != nil {
        panic(err)
    }
}
```

## API

init a new leblad instance

```go
l := leblad.New()
```
