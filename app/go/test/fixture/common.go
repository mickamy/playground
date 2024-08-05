package fixture

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
)

var (
	random *rand.Rand
)

func init() {
	var source = rand.NewSource(1)
	random = rand.New(source)

	gofakeit.GlobalFaker = gofakeit.NewFaker(random, true)
}
