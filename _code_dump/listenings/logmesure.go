
type QueryResults struct {
	qtype    string
	end      time.Time
	duration time.Duration
}
// //////////////////
var qr []QueryResults
// ///////////////////
for _, d := range qr {

	_, err = db2.Exec("Insert INTO QueryResults (query_type,timeEnded,duration_ms) values (?,?,?)", d.qtype, d.end.Format("2006-01-02 15:04:05"), d.duration.Milliseconds())
	if err != nil {
		log.Printf("[worker %d] result insert error: %v", id, err)
	}

}
_, err = db.Exec("Insert INTO Tests (name,timeStart,timeEnd) values (?,?,?)", "Select range workload 1h 20threads p", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05"))
if err != nil {
	log.Printf("result insert error: %v", err)
}
