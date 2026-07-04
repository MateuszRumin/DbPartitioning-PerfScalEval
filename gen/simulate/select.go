package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand/v2"
	sqlgen "simulate/sqlgenerate"
	"time"
)

func newWorkerRand() *rand.Rand {
	var b [16]byte
	_, _ = crand.Read(b[:])

	s1 := binary.LittleEndian.Uint64(b[:8])
	s2 := binary.LittleEndian.Uint64(b[8:])

	return rand.New(rand.NewPCG(s1, s2))
}

func selectTest(id int, r *rand.Rand, wg *sqlgen.WorkerGenerator, idp int, idu int, deadline time.Time) {
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
		err := executeQuerySelect(db, query)
		if err != nil {
			log.Printf("[worker %d] query error: %v", id, err)
			continue
		}

		qr = append(qr, QueryResults{
			qtype:    "SELECT", // np. SELECT, INSERT, UPDATE
			end:      time.Now(),
			duration: time.Since(start),
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
