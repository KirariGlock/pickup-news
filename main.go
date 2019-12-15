package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Env struct {
	Keyword string
	From    string
	To      string
}

func main() {
	resuest, err := http.NewRequest("GET", "https://newsapi.org/v2/everything", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keyword := "Google"

	values := url.Values{}
	values.Add("qInTitle", keyword)
	values.Add("from", "2019-12-07")
	values.Add("to", "2019-12-07")
	values.Add("apiKey", "") // TODO 環境変数で渡す
	resuest.URL.RawQuery = values.Encode()

	client := new(http.Client)

	resp, err := client.Do(resuest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Printf("Unable to get this url : http status is %d \n", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	naResp := new(NewsAPIRespons)
	if err := json.Unmarshal(body, &naResp); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Keyword: %s resultCount: %d", keyword, naResp.TotalResults)
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
