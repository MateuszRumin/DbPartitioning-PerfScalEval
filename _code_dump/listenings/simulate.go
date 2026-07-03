
var wg sync.WaitGroup

start := time.Now()

go func() {
	defer wg.Done()
	multiThreadSelect(workersCountSelect, deadline, idpi, idui)
}()

go func() {
	defer wg.Done()
	multiThreadInsert(workersCountInsert, deadline, idb, idc, idph, idp, idu)

}()

go func() {
	defer wg.Done()
	multiThreadUpdate(workersCountUpdate, deadline, idb, idc, idph, idp, idu)
}()

wg.Wait()

stop := time.Now()
	