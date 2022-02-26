# formatparser

![test](https://github.com/ghosind/formatparser/workflows/Test/badge.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/79fba4a6e6084cc8b5755450e7bff9ef)](https://www.codacy.com/gh/ghosind/formatparser/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ghosind/formatparser&amp;utm_campaign=Badge_Grade)
![Version Badge](https://img.shields.io/github/v/release/ghosind/formatparser)
![License Badge](https://img.shields.io/github/license/ghosind/formatparser)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/formatparser.svg)](https://pkg.go.dev/github.com/ghosind/formatparser)

Parse format string to text token part or key token part arrays.

## Installation

```sh
go get -u github.com/ghosind/formatparser
```

## Getting Started

```go
package main

import (
  "log"

  "github.com/ghosind/formatparser"
)

func main() {
  format := "[%time] %message"

  parts := formatparser.Parse(format)

  for _, part := range parts {
    log.Println(part)
  }
}
```

```sh
$ go run app.go
&{1 [}
&{2 time}
&{1 ] }
&{2 %message}
```
