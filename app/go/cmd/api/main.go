package main

import (
	_ "mickamy.com/playground/internal/lib/logger"

	"mickamy.com/playground/internal/api"
)

func main() {
	api.Run()
}
