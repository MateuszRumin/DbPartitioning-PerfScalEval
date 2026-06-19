package dbpartitioningperfscaleval

import (
	"fmt"
	"math/rand/v2"
	"time"
)

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
