
var qr []QueryResults

for time.Now().Before(deadline) {

	query := wg.GenerateRandomQuery(r, idp, idu)
	if query == "" {
		continue
	}
	start := time.Now()
	err := executeQuery(db, query)
	if err != nil {
		log.Printf("[worker %d] query error: %v", id, err)
		continue
	}
	stop := time.Now()
	duration := time.Since(start)

	qr = append(qr, QueryResults{
		qtype:    "SELECT",
		end:      stop,
		duration: duration,
	})

}
