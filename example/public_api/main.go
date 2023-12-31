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
	st, err := c.Market.Time(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(st)
}
