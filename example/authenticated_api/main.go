package main

import (
	"context"
	"fmt"

	"github.com/jferrl/go-kraken"
)

func main() {
	ctx := context.Background()

	c := kraken.New(nil).
		WithAuth(
			kraken.Secrets{},
		)

	st, err := c.Account.Balance(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(st)
}
