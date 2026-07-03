idp, idu, err := sqlgen.GenerateInitValues()
var wg sync.WaitGroup
start := time.Now()
deadline := time.Now().Add(10 * time.Minute)
for i := 0; i < threadsCount; i++ {
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		r := newWorkerRand()
		wg := sqlgen.NewWorkerGenerator(r)
		runTest(deadline, id, r, wg, idp, idu)
	}(i)
}
wg.Wait()
