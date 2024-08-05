package main

import (
	"context"
	"fmt"

	"mickamy.com/playground/internal/cli/db"
)

func main() {
	fmt.Println("Dropping database...")
	if err := db.Drop(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
