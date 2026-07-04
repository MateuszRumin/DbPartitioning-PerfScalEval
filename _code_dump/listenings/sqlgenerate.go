


type WeightedGenerator struct {
	Weight int
	Gen    QueryGenerator
}

var Generators = []WeightedGenerator{
	{20, PostByIDQuery{}},
	{15, QuestionDetailAndAuthor{}},
	{15, QuestionAnswers{}}, //lq
	{8, CommentsForPost{}},  //480
}

type WorkerGenerator struct {
	Generators []WeightedGenerator
	Prefix     []int
	Total      int
}

func (wg *WorkerGenerator) GenerateRandomQuery(r *rand.Rand, idp, idu int) string {
	pick := r.IntN(wg.Total)

	i := sort.Search(len(wg.Prefix), func(i int) bool {
		return pick < wg.Prefix[i]
	})

	return wg.Generators[i].Gen.Generate(r, idp, idu)
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

