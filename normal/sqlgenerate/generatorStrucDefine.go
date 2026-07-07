package sqlgenerate

import (
	"math/rand/v2"
	"sort"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, idp []int, idu []int) string
}

type PostByIDQuery struct{}
type NewQuestions struct{}
type QuestionDetailAndAuthor struct{}
type QuestionAnswers struct{}
type QuestionBestAnswers struct{}
type CommentsForPost struct{}
type UserProfile struct{}
type UserActivity struct{}
type PostHistory struct{}
type ConnectedPosts struct{}
type ReputationTrend struct{}
type TopControversialPosts struct{}
type TopCommentedPosts struct{}
type LastGlobalActivity struct{}
type ConnectedPostsMultiple struct{}
type PopularPosts7d struct{}
type UserRankingByPostPopularity30d struct{}
type TextSearchPosts struct{}
type TagSearchPosts struct{}
type QuestionsForTag struct{}

//type struct{}
//type struct{}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var Generators = []WeightedGenerator{

	{15, PostByIDQuery{}},
	{12, NewQuestions{}},
	{12, QuestionDetailAndAuthor{}},
	{10, QuestionAnswers{}},
	{4, QuestionBestAnswers{}},
	{8, CommentsForPost{}},
	{5, UserProfile{}},
	{6, UserActivity{}},
	{2, PostHistory{}},
	{3, ConnectedPosts{}},
	{1, ReputationTrend{}},
	{2, TopControversialPosts{}},
	{3, TopCommentedPosts{}},
	{4, LastGlobalActivity{}},
	{2, ConnectedPostsMultiple{}},
	{3, PopularPosts7d{}},
	{1, UserRankingByPostPopularity30d{}},
	{2, TextSearchPosts{}},
	{2, TagSearchPosts{}},
	{3, QuestionsForTag{}},
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
