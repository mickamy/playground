package main

import (
	"context"
	"fmt"

	"mickamy.com/playground/internal/cli/db"
)

func main() {
	fmt.Println("Resetting database...")
	ctx := context.Background()
	if err := db.Drop(ctx); err != nil {
		panic(err)
	}
	if err := db.Create(ctx); err != nil {
		panic(err)
	}
	if err := db.Migrate(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
