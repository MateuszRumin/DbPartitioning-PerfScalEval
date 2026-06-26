package sqlgenerate

import (
	"math/rand/v2"
	"sort"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, idp int, idu int) string
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
	// bardzo częste

	{20, PostByIDQuery{}},
	{15, QuestionDetailAndAuthor{}},
	{15, QuestionAnswers{}}, //lq
	{8, CommentsForPost{}},  //480

	// feedy

	{8, NewQuestions{}},   //lq18-0
	{4, PopularPosts7d{}}, //sl19-0
	{3, LastGlobalActivity{}},

	// profile

	{6, UserProfile{}},
	{4, UserActivity{}},

	// wyszukiwanie

	{4, TextSearchPosts{}},
	{3, QuestionsForTag{}},
	{2, TagSearchPosts{}},

	// powiązania

	{2, ConnectedPosts{}},
	{2, ConnectedPostsMultiple{}},

	// pojedyncze dodatkowe odczyty

	{2, QuestionBestAnswers{}},

	// rzadkie

	{1, PostHistory{}},
	{1, ReputationTrend{}},
	{1, TopCommentedPosts{}},
	{1, TopControversialPosts{}},
	{1, UserRankingByPostPopularity30d{}},
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
func (wg *WorkerGenerator) GenerateRandomQuery(r *rand.Rand, idp, idu int) string {
	pick := r.IntN(wg.Total)

	i := sort.Search(len(wg.Prefix), func(i int) bool {
		return pick < wg.Prefix[i]
	})

	return wg.Generators[i].Gen.Generate(r, idp, idu)
}
