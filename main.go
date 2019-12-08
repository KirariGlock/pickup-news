package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	resuest, err := http.NewRequest("GET", "https://newsapi.org/v2/everything", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := url.Values{}
	values.Add("qInTitle", "Google")
	values.Add("from", "2019-12-07")
	values.Add("to", "2019-12-07")
	values.Add("apiKey", "") // TODO 環境変数で渡す
	resuest.URL.RawQuery = values.Encode()

	client := new(http.Client)

	resp, err := client.Do(resuest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO

}

type NewsAPIRespons struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}
