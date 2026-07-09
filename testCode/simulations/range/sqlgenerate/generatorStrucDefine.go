package sqlgenerate

import (
	"math/rand/v2"
	"sort"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, idp []int, idu []int) string
}

type NewQuestions struct{}
type UserActivity struct{}
type TopControversialPosts struct{}
type TopCommentedPosts struct{}
type PopularPosts7d struct{}
type UserRankingByPostPopularity30d struct{}
type TextSearchPosts struct{}
type TagSearchPosts struct{}
type QuestionsForTag struct{}
type PostByScoreHigh struct{}
type PostByScoreLow struct{}
type PostByViewCoHigh struct{}
type PostByViewCoLow struct{}
type PostsByScoreRange struct{}
type PostsByViewRange struct{}
type PostByAnswerCount struct{}
type PostByOwnerAndScore struct{}
type UserPosts struct{}
type AveragePostViews struct{}
type PostPerUser struct{}
type SingleDayLookUp struct{}

type DateRandomAreaScore struct{}
type DateRandomAreaViewCount struct{}
type DateRandomAreaOrderScore struct{}

//type struct{}
//type struct{}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var Generators = []WeightedGenerator{
	{3, NewQuestions{}},   //lq18-0
	{3, PopularPosts7d{}}, //sl19-0
	{3, UserActivity{}},
	{3, TextSearchPosts{}},
	{3, QuestionsForTag{}},
	{3, TagSearchPosts{}},
	{3, TopCommentedPosts{}},
	{3, TopControversialPosts{}},
	{3, UserRankingByPostPopularity30d{}},
	{3, PostByScoreHigh{}},
	{3, PostByScoreLow{}},
	{3, PostByViewCoHigh{}},
	{3, PostByViewCoLow{}},
	{3, PostsByScoreRange{}},
	{3, PostsByViewRange{}},
	{3, PostByAnswerCount{}},
	{3, PostByOwnerAndScore{}},
	{3, UserPosts{}},
	{3, AveragePostViews{}},
	{3, PostPerUser{}},
	{3, SingleDayLookUp{}},
	{3, DateRandomAreaScore{}},
	{3, DateRandomAreaViewCount{}},
	{3, DateRandomAreaOrderScore{}},
}

type WorkerGenerator struct {
	Generators []WeightedGenerator
	Prefix     []int
	Total      int
}

func NewWorkerGenerator(r *rand.Rand) *WorkerGenerator {
	gs := Generators
	w := &WorkerGenerator{
		Generators: gs,
		Prefix:     make([]int, len(gs)),
	}

	sum := 0
	for i, g := range gs {
		sum += g.Weight
		w.Prefix[i] = sum
	}

	w.Total = sum
	return w
}
func (wg *WorkerGenerator) GenerateRandomQuery(r *rand.Rand, idp, idu []int) string {
	pick := r.IntN(wg.Total)

	i := sort.Search(len(wg.Prefix), func(i int) bool {
		return pick < wg.Prefix[i]
	})

	return wg.Generators[i].Gen.Generate(r, idp, idu)
}
