package services

import (
	"chuckEspaNorris/internal/libs"
	"chuckEspaNorris/internal/repositories"
	"sync"
)

func GetRandomJokes() []string {
	var results []string = make([]string, 0)
	NUMBER_JOKES := 3

	var wg sync.WaitGroup
	for i := 0; i < NUMBER_JOKES; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := repositories.ReadRandomJoke()
			results = append(results, res)
		}()
	}
	wg.Wait()

	for i, englishJoke := range results {
		results[i] = libs.TranslateEnToEs(englishJoke)
	}

	return results
}