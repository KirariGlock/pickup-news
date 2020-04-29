package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Apikey     string // NewsAPI api key
	WebhookURL string // Slack webhook url
}

type RequestParameter struct {
	Keyword          string
	From             string
	To               string
	NoticeLowerLimit int `default:"0"` // Don't notify if the number of news is below NoticeLowerLimit
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, rp RequestParameter) (string, error) {
	// import enviroment
	var env Env
	if err := envconfig.Process("pickupnews", &env); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create request
	resuest, err := http.NewRequest("GET", "http://newsapi.org/v2/everything", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := url.Values{}
	values.Add("qInTitle", rp.Keyword)
	values.Add("from", rp.From)
	values.Add("to", rp.To)
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

	if naResp.TotalResults <= rp.NoticeLowerLimit {
		return fmt.Sprintf("TotalResult is lower NoticeLowerLimit. TotalResult:%d, NoticeLowerLimit:%d\n", naResp.TotalResults, rp.NoticeLowerLimit), nil
	}

	messageHeader := "<!channel> Keyword: " + rp.Keyword + " resultCount: " + strconv.Itoa(naResp.TotalResults) + "\n"
	var messageDetail bytes.Buffer
	for i, article := range naResp.Articles {
		messageDetail.WriteString("No.")
		messageDetail.WriteString(strconv.Itoa(i + 1))
		messageDetail.WriteString(", ")
		messageDetail.WriteString(article.Title)
		messageDetail.WriteString(", ")
		messageDetail.WriteString(article.URL)
		messageDetail.WriteString("\n")
	}

	notificationSlack(env, messageHeader+messageDetail.String())
	return "Success notification.", nil
}

func notificationSlack(env Env, message string) {
	params := `{"text":"` + message + `"}`
	resuest, err := http.NewRequest("POST", env.WebhookURL, bytes.NewBuffer([]byte(params)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resuest.Header.Set("Content-Type", "application/json")

	// Execute slack webhook
	client := new(http.Client)
	resp, err := client.Do(resuest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Printf("Unable to post this url : http status is %d \n", resp.StatusCode)
	}
	defer resp.Body.Close()
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
