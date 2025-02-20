package test

import (
	"fmt"
	"math/rand"
	"perfscaleval/config"
	"perfscaleval/queries"
	"sync"
	"time"
)

func createControl() {
	// Parametry kontrolujące
	executionsPerSecond := 1             // Liczba wykonań pętli na sekundę
	duration := 30 * time.Second         // Czas działania wątku
	constantSpeed := false               // Czy pętla ma działać z tą samą szybkością
	speedChangeAfter := 20 * time.Second // Po jakim czasie zmienić szybkość
	newExecutionsPerSecond := 20         // Nowa liczba wykonań pętli na sekundę po zmianie

	var wg sync.WaitGroup

	for i := 0; i < config.NumConnections; i++ {
		wg.Add(1)
		go func(connIndex int) {
			defer wg.Done()

			bufor := makeBuforDataForThread()

			runThreadLoop(bufor, executionsPerSecond, duration, constantSpeed, speedChangeAfter, newExecutionsPerSecond, i)
		}(i)
	}

	wg.Wait()
}

func runThreadLoop(bufor BuforData, executionsPerSecond int, duration time.Duration, constantSpeed bool, speedChangeAfter time.Duration, newExecutionsPerSecond int, conNum int) {
	startTime := time.Now()
	// Inicjalizacja tickera z początkową częstotliwością
	interval := time.Second / time.Duration(executionsPerSecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Zmienna do przechowywania nowego interwału po zmianie szybkości
	var newInterval time.Duration
	if !constantSpeed {
		newInterval = time.Second / time.Duration(newExecutionsPerSecond)
	}

	// Pętla główna
	for {
		select {
		case <-ticker.C:
			// Wybierz akcję do wykonania na buforze
			chooseAction(bufor, conNum)

			// Jeśli nie ma stałej szybkości i minął czas zmiany szybkości
			if !constantSpeed && time.Since(startTime) > speedChangeAfter {
				// Zatrzymaj stary ticker i utwórz nowy z nowym interwałem
				ticker.Stop()
				ticker = time.NewTicker(newInterval)
				constantSpeed = true // Zapobiegaj wielokrotnej zmianie
			}
		default:
			// Sprawdź, czy czas działania wątku minął
			if time.Since(startTime) > duration {
				return
			}
		}
	}
}

func chooseAction(bufor BuforData, conNum int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rng.Intn(7)
	randomElem := rand.Intn(config.BuforCreateSize)
	switch randomIndex {
	case 0:
		fmt.Printf("Wylosowano Badges (liczba elementów: %d)\n", len(bufor.Badges))
		queries.CreateBadge(bufor.Badges[randomElem], conNum)
	case 1:
		fmt.Printf("Wylosowano Comments (liczba elementów: %d)\n", len(bufor.Comments))
		queries.CreateComment(bufor.Comments[randomElem], conNum)
	case 2:
		fmt.Printf("Wylosowano PostHistory (liczba elementów: %d)\n", len(bufor.PostHistory))
		queries.CreatePostHistory(bufor.PostHistory[randomElem], conNum)
	case 3:
		fmt.Printf("Wylosowano PostLinks (liczba elementów: %d)\n", len(bufor.PostLinks))
		queries.CreatePostLink(bufor.PostLinks[randomElem], conNum)
	case 4:
		fmt.Printf("Wylosowano Posts (liczba elementów: %d)\n", len(bufor.Posts))
		queries.CreatePost(bufor.Posts[randomElem], conNum)
	case 5:
		fmt.Printf("Wylosowano Users (liczba elementów: %d)\n", len(bufor.Users))
		queries.CreateUser(bufor.Users[randomElem], conNum)
	case 6:
		fmt.Printf("Wylosowano Votes (liczba elementów: %d)\n", len(bufor.Votes))
		queries.CreateVote(bufor.Votes[randomElem], conNum)
	default:
		fmt.Println("Nieznany indeks")
	}

}

// fmt.Printf("Print randomIndex %d \n", randomIndex)
// fmt.Printf("Print randomElem %d \n", randomElem)
// fmt.Printf("Print bufor len %d \n", len(bufor.Badges))
// fmt.Print(bufor.Badges[0])
