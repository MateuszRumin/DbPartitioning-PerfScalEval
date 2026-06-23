package sqlgenerate

import (
	"math/rand/v2"
)

type QueryGenerator interface {
	Generate(r *rand.Rand, ids map[string][]int) string
}

type UserByIDQuery struct{}
type PostByIDQuery struct{}
type PostByScoreHigh struct{}
type PostByScoreLow struct{}
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
type ComByScorHigh struct{}
type ComByScorLow struct{}
type ComByUsrID struct{}
type VotCountByPosId struct{}
type PostPerUser struct{}
type AveragePostViews struct{}
type AgregatePostScore struct{}
type AgregateVotes struct{}
type UserPostVotes struct{}
type PostVotes struct{}
type PostComments struct{}
type UserPosts struct{}
type PostByOwnerAndScore struct{}
type PostByAnswerCount struct{}
type PostTopViewed struct{}
type PostsByViewRange struct{}
type PostsByScoreRange struct{}
type DateRandomAreaOrderScore struct{}
type DateRandomAreaViewCount struct{}
type DateRandomAreaScore struct{}
type DateRandomArea struct{}
type RecentPosts struct{}
type SingleMonthLookUp struct{}
type SingleDayLookUp struct{}

type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var generators = []WeightedGenerator{

	// ===== PK LOOKUPS =====

	{80, PostByIDQuery{}},
	{40, UserByIDQuery{}},
	{25, PostHiById{}},

	// ===== USER-CENTRIC =====

	{60, PostByOvnId{}},
	{20, PostByLastEditorId{}},
	{40, UserPosts{}},
	{25, ComByUsrID{}},

	// ===== POST DETAILS =====

	{50, ComByPosId{}},
	{40, PostComments{}},
	{35, PostVotes{}},
	{25, VotCountByPosId{}},

	// ===== PARTITION PRUNING TESTS =====

	{80, SingleDayLookUp{}},
	{70, SingleMonthLookUp{}},
	{50, DateRandomArea{}},

	// ===== PARTITION + SECONDARY FILTER =====

	{45, DateRandomAreaScore{}},
	{45, DateRandomAreaViewCount{}},
	{30, DateRandomAreaOrderScore{}},

	// ===== RANGE SCANS =====

	{35, PostsByScoreRange{}},
	{35, PostsByViewRange{}},

	// ===== COMMON FILTERS =====

	{20, PostByScoreHigh{}},
	{15, PostByScoreLow{}},

	{20, PostByViewCoHigh{}},
	{15, PostByViewCoLow{}},

	// ===== JOINS =====

	{20, UserPostVotes{}},

	// ===== HISTORY =====

	{15, PostHiByPosId{}},
	{10, PostHiByUsrId{}},
	{5, PostHiByTypeId{}},

	// ===== LINKS =====

	{10, PostLiByPosId{}},
	{10, PostLiByRelPosId{}},
	{5, PostLiByLiTyId{}},

	// ===== ANALYTICAL =====

	{8, AveragePostViews{}},
	{8, AgregatePostScore{}},
	{8, PostPerUser{}},
	{8, AgregateVotes{}},

	// ===== EXPENSIVE REPORTS =====

	{8, PostByOwnerAndScore{}},
	{8, PostByAnswerCount{}},
	{8, PostTopViewed{}},

	// ===== RECENT DATA =====

	{5, RecentPosts{}},

	// ===== COMMENTS =====

	{5, ComByScorHigh{}},
	{5, ComByScorLow{}},
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
