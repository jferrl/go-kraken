# go-kraken

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/jferrl/go-kraken)
[![Test Status](https://github.com/jferrl/go-kraken/workflows/tests/badge.svg)](https://github.com/jferrl/go-kraken/actions?query=workflow%3Atests)
[![CodeQL](https://github.com/jferrl/go-kraken/workflows/CodeQL/badge.svg)](https://github.com/jferrl/go-kraken/actions?query=workflow%3ACodeQL)
[![codecov](https://codecov.io/gh/jferrl/go-kraken/branch/main/graph/badge.svg?token=68I4BZF235)](https://codecov.io/gh/jferrl/go-kraken)

Go library for accessing the Kraken API.

Docs url <https://docs.kraken.com/rest/>.

## Usage

go-kraken is compatible with modern Go releases.

Build a new client, then you can use the services to reach different parts of the Kraken API. For example:

```go
package main

import (
 "context"
 "fmt"

 "github.com/jferrl/go-kraken"
)

func main() {
 ctx := context.Background()

 c := kraken.New(nil)

 // Get server time
 st, err := c.MarketData.Time(ctx)
 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println(st)
}
```

Using the `context` package, you can easily pass cancelation signals and
deadlines to various services of the client for handling a request.

## Token Creation

## Authentication

## License

This library is distributed under the BSD-style license found in the LICENSE file.
