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

func wantConnection(deadline time.Time, id int, r *rand.Rand, wg *sqlgen.WorkerGenerator, idp []int, idu []int) []QueryResults {
	fmt.Println("Worker:", id)

	db, err := setConnection()
	if err != nil {
		log.Printf("[worker %d] DB error: %v", id, err)
		return nil
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

	return qr
}
