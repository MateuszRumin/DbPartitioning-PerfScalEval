package sqlgenerate

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type SingleDayLookUp struct{}

func (q SingleDayLookUp) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date = '%s';
	`, randomDate.Format("2006-01-02"))
}

type SingleMonthLookUp struct{}

func (q SingleMonthLookUp) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))

	month1 := time.Date(
		randomDate.Year(),
		randomDate.Month(),
		1,
		0, 0, 0, 0,
		time.UTC,
	)

	month2 := month1.AddDate(0, 1, 0)

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date >= '%s'
		  AND creation_date < '%s';
	`, month1.Format("2006-01-02"), month2.Format("2006-01-02"))
}

type RecentPosts struct{}

func (q RecentPosts) Generate(r *rand.Rand, ids map[string][]int) string {
	cutoff := time.Now().AddDate(0, 0, -1)

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date > '%s';
	`, cutoff.Format("2006-01-02 15:04:05"))
}

type DateRandomArea struct{}

func (q DateRandomArea) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	days2 := int(end.Sub(randomDate).Hours() / 24)
	randomDate2 := randomDate.AddDate(0, 0, r.IntN(days2+1))

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date >= '%s'
		  AND creation_date < '%s';
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

type DateRandomAreaScore struct{}

func (q DateRandomAreaScore) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	days2 := int(end.Sub(randomDate).Hours() / 24)
	randomDate2 := randomDate.AddDate(0, 0, r.IntN(days2+1))
	score := r.IntN(1000)

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date >= '%s'
		  AND creation_date < '%s'
		  AND score > %d ;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), score)
}

type DateRandomAreaViewCount struct{}

func (q DateRandomAreaViewCount) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	days2 := int(end.Sub(randomDate).Hours() / 24)
	randomDate2 := randomDate.AddDate(0, 0, r.IntN(days2+1))
	score := r.IntN(1000)

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date >= '%s'
		  AND creation_date < '%s'
		  AND view_count > %d ;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), score)
}

type DateRandomAreaOrderScore struct{}

func (q DateRandomAreaOrderScore) Generate(r *rand.Rand, ids map[string][]int) string {
	start := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	days2 := int(end.Sub(randomDate).Hours() / 24)
	randomDate2 := randomDate.AddDate(0, 0, r.IntN(days2+1))

	return fmt.Sprintf(`
		SELECT *
		FROM posts
		WHERE creation_date >= '%s'
		  AND creation_date < '%s'
		ORDER BY score DESC;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}
