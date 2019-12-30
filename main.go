package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Keyword string
	From    string
	To      string
	Apikey  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", returnPong)
	r.HandleFunc("/newsTitles", returnNewsTitles)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func returnPong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Printf("pong!")
}

func returnNewsTitles(w http.ResponseWriter, r *http.Request) {
	// import enviroment
	var env Env
	if err := envconfig.Process("pickupnews", &env); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create request
	resuest, err := http.NewRequest("GET", "https://newsapi.org/v2/everything", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := url.Values{}
	values.Add("qInTitle", env.Keyword)
	values.Add("from", env.From)
	values.Add("to", env.To)
	values.Add("apiKey", env.Apikey)
	resuest.URL.RawQuery = values.Encode()

	// execute NewsAPI
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

	fmt.Printf("Keyword: %s resultCount: %d \n", env.Keyword, naResp.TotalResults)
	for i, article := range naResp.Articles {
		fmt.Printf("No.%d, %s,%s\n", i+1, article.Title, article.URL)
	}
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
