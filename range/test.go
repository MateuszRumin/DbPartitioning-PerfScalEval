package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand/v2"
	sqlgen "range/sqlgenerate"
	"time"
)

type QueryResults struct {
	qtype    string
	end      time.Time
	duration time.Duration
}

func newWorkerRand() *rand.Rand {
	var b [16]byte
	_, _ = crand.Read(b[:])

	s1 := binary.LittleEndian.Uint64(b[:8])
	s2 := binary.LittleEndian.Uint64(b[8:])

	return rand.New(rand.NewPCG(s1, s2))
}

func wantConnection(deadline time.Time, id int, r *rand.Rand, wg *sqlgen.WorkerGenerator, idp []int, idu []int) {
	fmt.Println("Worker:", id)

	db, err := setConnection()
	if err != nil {
		log.Printf("[worker %d] DB error: %v", id, err)
		return
	}
	defer db.Close()
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
			qtype:    "SELECT", // np. SELECT, INSERT, UPDATE
			end:      stop,
			duration: duration,
		})

	}

	db2, err := slc()
	if err != nil {

		return
	}
	defer db2.Close()

	for _, d := range qr {

		_, err = db2.Exec("Insert INTO QueryResults (query_type,timeEnded,duration_ms) values (?,?,?)", d.qtype, d.end.Format("2006-01-02 15:04:05"), d.duration.Milliseconds())
		if err != nil {
			log.Printf("[worker %d] result insert error: %v", id, err)
		}

	}
}
