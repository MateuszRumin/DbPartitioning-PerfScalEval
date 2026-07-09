package sqlgenerate

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
	"time"
)

var ReadGenerators = []WeightedGenerator{
	{Weight: 18, Gen: QuestionPage{}},
	{Weight: 17, Gen: QuestionAnswers{}},
	{Weight: 11, Gen: CommentsForPost{}},
	{Weight: 4, Gen: AcceptedAnswer{}},
	{Weight: 9, Gen: NewestQuestions{}},
	{Weight: 6, Gen: ActiveQuestions{}},
	{Weight: 7, Gen: TagNewestQuestions{}},
	{Weight: 5, Gen: UserProfile{}},
	{Weight: 4, Gen: UserRecentPosts{}},
	{Weight: 4, Gen: RelatedPosts{}},
	{Weight: 3, Gen: UnansweredQuestions{}},
	{Weight: 3, Gen: TextSearchPosts{}},
	{Weight: 2, Gen: HotQuestions{}},
	{Weight: 1, Gen: PostTimeline{}},
	{Weight: 2, Gen: VoteBreakdown{}},
	{Weight: 2, Gen: UserRecentComments{}},
	{Weight: 1, Gen: UserBadges{}},
	{Weight: 1, Gen: PopularTags{}},
}

var (
	datasetStart        = time.Date(2008, 7, 31, 0, 0, 0, 0, time.UTC)
	datasetEndExclusive = time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)
)

var commonTags = []string{
	"c#", "java", "javascript", "php", "c++", "python", "sql", "mysql",
	"asp.net", ".net", "jquery", "html", "css", "android", "iphone",
	"objective-c", "ruby-on-rails", "regex", "xml", "database", "linux",
	"windows", "visual-studio", "web-services", "performance", "multithreading",
}

var commonSearchTerms = []string{
	"error", "exception", "null", "array", "string", "database", "query",
	"function", "class", "server", "file", "connection", "timeout", "memory",
	"thread", "json", "xml", "date", "sort", "join", "index", "update",
	"authentication", "encoding", "request", "response",
}

const noOpQuery = `SELECT 1 WHERE FALSE;`

func NewWorkerGenerator(generators []WeightedGenerator) (*WorkerGenerator, error) {
	if len(generators) == 0 {
		return nil, errors.New("sqlgenerate: empty generator list")
	}

	wg := &WorkerGenerator{
		generators: append([]WeightedGenerator(nil), generators...),
		prefix:     make([]int, len(generators)),
	}

	for i, item := range wg.generators {
		if item.Gen == nil {
			return nil, fmt.Errorf("sqlgenerate: generator %d is nil", i)
		}
		if item.Weight <= 0 {
			return nil, fmt.Errorf("sqlgenerate: generator %d has non-positive weight", i)
		}

		wg.total += item.Weight
		wg.prefix[i] = wg.total
	}

	return wg, nil
}

func (wg *WorkerGenerator) GenerateRandomQuery(
	r *rand.Rand,
	idp []int,
	idu []int,
) string {
	pick := r.IntN(wg.total)
	idx := sort.Search(len(wg.prefix), func(i int) bool {
		return pick < wg.prefix[i]
	})

	return wg.generators[idx].Gen.Generate(r, idp, idu)
}

func ValidateIDPools(idp []int, idu []int) error {
	if len(idp) == 0 {
		return errors.New("sqlgenerate: post ID pool is empty")
	}
	if len(idu) == 0 {
		return errors.New("sqlgenerate: user ID pool is empty")
	}
	return nil
}

func randomID(r *rand.Rand, ids []int) (int, bool) {
	if len(ids) == 0 {
		return 0, false
	}
	return ids[r.IntN(len(ids))], true
}

func randomFrom[T any](r *rand.Rand, values []T) T {
	return values[r.IntN(len(values))]
}

func sqlString(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}

func clampDate(value time.Time) time.Time {
	if value.Before(datasetStart) {
		return datasetStart
	}
	if !value.Before(datasetEndExclusive) {
		return datasetEndExclusive.AddDate(0, 0, -1)
	}
	return value
}

func randomDateBetween(r *rand.Rand, start, endExclusive time.Time) time.Time {
	start = clampDate(start)

	if endExclusive.After(datasetEndExclusive) {
		endExclusive = datasetEndExclusive
	}
	if !endExclusive.After(start) {
		return start
	}

	days := int(endExclusive.Sub(start) / (24 * time.Hour))
	if days <= 0 {
		return start
	}

	return start.AddDate(0, 0, r.IntN(days))
}

// randomAnchorDate odzwierciedla długi ogon forum:
// 70% ruchu dotyczy ostatnich 90 dni zbioru,
// 20% ostatniego roku, 10% całej historii.
func randomAnchorDate(r *rand.Rand) time.Time {
	lastDay := datasetEndExclusive.AddDate(0, 0, -1)
	roll := r.IntN(100)

	var start time.Time
	switch {
	case roll < 70:
		start = lastDay.AddDate(0, 0, -89)
	case roll < 90:
		start = lastDay.AddDate(0, 0, -364)
	default:
		start = datasetStart
	}

	return randomDateBetween(r, start, datasetEndExclusive)
}

func randomWindow(r *rand.Rand) (time.Time, time.Time) {
	days := randomFrom(r, []int{
		1, 1, 1,
		7, 7, 7,
		30, 30,
		90,
	})

	end := randomAnchorDate(r).AddDate(0, 0, 1)
	if end.After(datasetEndExclusive) {
		end = datasetEndExclusive
	}

	start := end.AddDate(0, 0, -days)
	if start.Before(datasetStart) {
		start = datasetStart
	}

	return start, end
}

func randomPageSize(r *rand.Rand) int {
	return randomFrom(r, []int{15, 15, 15, 20, 30, 30, 50})
}

func randomOffset(r *rand.Rand, pageSize int) int {
	roll := r.IntN(100)

	switch {
	case roll < 82:
		return 0
	case roll < 95:
		return pageSize
	case roll < 99:
		return pageSize * (2 + r.IntN(3))
	default:
		return pageSize * (5 + r.IntN(16))
	}
}

func randomTag(r *rand.Rand) string {
	return randomFrom(r, commonTags)
}

func randomSearchTerm(r *rand.Rand) string {
	return randomFrom(r, commonSearchTerms)
}

type QueryGenerator interface {
	Generate(r *rand.Rand, idp []int, idu []int) string
}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

type WorkerGenerator struct {
	generators []WeightedGenerator
	prefix     []int
	total      int
}
