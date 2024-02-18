package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/bregydoc/gtranslate"
)

func main() {

	var results []string = make([]string, 0)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := fetchJoke()
			results = append(results, res)
		}()
	}
	wg.Wait()

	for _, eng := range(results) {
		trans, _ := gtranslate.TranslateWithParams(
			eng,
			gtranslate.TranslationParams{
				From: "en",
				To:   "es",
			},
		)
		fmt.Println(trans)
	}
}

type chuckResponse struct {
	value string
}

func fetchJoke() string {
	requestURL := "https://api.chucknorris.io/jokes/random"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return ""
	} else {
		// fmt.Printf("client: status code: %d\n", res.StatusCode)
		defer res.Body.Close()
		var result map[string]string
				// json.NewDecoder(res.Body).Decode(target) for making json to struct
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &result)
		//fmt.Println(result["value"])
		return result["value"]
	}
}