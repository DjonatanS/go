package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "password")
	getToken(ctx)
}

func getToken(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)

}
