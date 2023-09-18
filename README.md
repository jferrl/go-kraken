# go-kraken

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/jferrl/go-kraken)
[![Test Status](https://github.com/jferrl/go-kraken/workflows/tests/badge.svg)](https://github.com/jferrl/go-kraken/actions?query=workflow%3Atests)
[![CodeQL](https://github.com/jferrl/go-kraken/workflows/CodeQL/badge.svg)](https://github.com/jferrl/go-kraken/actions?query=workflow%3ACodeQL)
[![codecov](https://codecov.io/gh/jferrl/go-kraken/branch/main/graph/badge.svg?token=68I4BZF235)](https://codecov.io/gh/jferrl/go-kraken)
[![Go Report Card](https://goreportcard.com/badge/github.com/jferrl/go-kraken)](https://goreportcard.com/report/github.com/jferrl/go-kraken)

Go library for accessing the Kraken API. :octopus:

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

<https://pro.kraken.com/app/settings/api>

## Authentication

<https://docs.kraken.com/rest/#section/Authentication>

Authenticated requests must include both API-Key and API-Sign HTTP headers, and nonce in the request payload. otp is also required in the payload if two-factor authentication (2FA) is enabled.

### Nonce

Nonce must be an always increasing, unsigned 64-bit integer, for each request that is made with a particular API key. While a simple counter would provide a valid nonce, a more usual method of generating a valid nonce is to use e.g. a UNIX timestamp in milliseconds.

### 2FA

If two-factor authentication (2FA) is enabled for the API key and action in question, the one time password must be specified in the payload's otp value.

In order to set OTP in the request, you can use the `ContextWithOtp` function. Internally, OTP value is stored in the context and then used in the request.

 For example:

```go
package main

import (
 "context"

 "github.com/jferrl/go-kraken"
)

func main() {
 ctx := context.Background()

 ctxWithOpt := kraken.ContextWithOtp(ctx, "123456")
}
```

###  API-Key

The "API-Key" header should contain your API key.

Security Scheme Type: API Key
Header parameter name: API-Key

###  API-Sign

Authenticated requests should be signed with the "API-Sign" header, using a signature generated with your private key, nonce, encoded payload, and URI path according to:

`HMAC-SHA512 of (URI path + SHA256(nonce + POST data)) and base64 decoded secret API key`

## License

This library is distributed under the BSD-style license found in the LICENSE file.
