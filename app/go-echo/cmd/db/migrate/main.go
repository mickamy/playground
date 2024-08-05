package main

import (
	"context"
	"fmt"

	"mickamy.com/playground/internal/cli/db"
)

func main() {
	fmt.Println("Migrating database...")
	if err := db.Migrate(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
