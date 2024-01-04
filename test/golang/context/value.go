package main

import (
	"context"
	"fmt"
)

func withValue(ctx context.Context)  {
	fmt.Printf("name is %s\n", ctx.Value("name").(string))
}