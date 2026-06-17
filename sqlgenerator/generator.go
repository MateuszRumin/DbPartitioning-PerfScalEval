package main

import (
	"fmt"
	"math/rand/v2"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, ids map[string][]int) string
}

type UserByIDQuery struct{}

func (q UserByIDQuery) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT *
        FROM users
        WHERE id = %d
    `, id)
}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var generators = []WeightedGenerator{
	{25, UserByIDQuery{}},
}

func GenerateRandomQuery(r *rand.Rand, ids map[string][]int) string {
	total := 0

	for _, g := range generators {
		total += g.Weight
	}

	pick := r.IntN(total)

	current := 0

	for _, g := range generators {
		current += g.Weight

		if pick < current {
			return g.Gen.Generate(r, ids)
		}
	}

	panic("unreachable")
}
