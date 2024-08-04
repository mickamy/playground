package main

import (
	"context"
	"fmt"

	"mickamy.com/playground/internal/cli/db"
)

func main() {
	fmt.Println("Creating database...")
	if err := db.Create(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
