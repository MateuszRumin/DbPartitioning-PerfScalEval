package sqlgenerate

import (
	"math/rand/v2"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, ids map[string][]int) string
}

type UserByIDQuery struct{}
type PostByIDQuery struct{}
type PostsByScoreEq struct{}
type PostByScoreHigh struct{}
type PostByScoreLow struct{}
type PostByViewCoEq struct{}
type PostByViewCoHigh struct{}
type PostByViewCoLow struct{}
type PostByOvnId struct{}
type PostByLastEditorId struct{}
type PostHiById struct{}
type PostHiByTypeId struct{}
type PostHiByPosId struct{}
type PostHiByUsrId struct{}
type PostLiByPosId struct{}
type PostLiByRelPosId struct{}
type PostLiByLiTyId struct{}

type ComByPosId struct{}
type ComByScorEq struct{}
type ComByScorHigh struct{}

type ComByScorLow struct{}
type ComByUsrID struct{}
type VotCountByPosId struct{}

//type PostByAcpAnsIDNULL struct{}
//type PostByAcpAnsIDExist struct{}
//type PostByParIDNULL struct{}
//type PostByParIDExist struct{}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var generators = []WeightedGenerator{
	{10, UserByIDQuery{}},
	{10, PostByIDQuery{}},
	{10, PostsByScoreEq{}},
	{10, PostByScoreHigh{}},
	{10, PostByScoreLow{}},
	{10, PostByViewCoEq{}},
	{10, PostByViewCoHigh{}},
	{10, PostByViewCoLow{}},
	{10, PostByOvnId{}},
	{10, PostByLastEditorId{}},
	{10, PostHiById{}},
	{10, PostHiByTypeId{}},
	{10, PostHiByPosId{}},
	{10, PostHiByUsrId{}},
	{10, PostLiByPosId{}},
	{10, PostLiByRelPosId{}},
	{10, PostLiByLiTyId{}},
	{10, ComByPosId{}},
	{10, ComByScorEq{}},
	{10, ComByScorHigh{}},
	{10, ComByScorLow{}},
	{10, ComByUsrID{}},
	{10, VotCountByPosId{}},
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
