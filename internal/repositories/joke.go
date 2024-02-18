package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadRandomJoke() string {
	requestURL := "https://api.chucknorris.io/jokes/random"
	res, err := http.Get(requestURL)
	if err != nil || res.StatusCode != 200 {
		fmt.Printf("error making http request: %s\nWith response status %v", err, res.StatusCode)
		return ""
	} else {
		defer res.Body.Close()
		var result map[string]string
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &result)
		return result["value"]
	}
}