# pickup-news
Output news that matches the keyword.

## Usage
### Preparations
- You get NewsAPI API key.  
https://newsapi.org/

- You get Slack incoming webhook url.
https://slack.com/intl/ja-jp/help/articles/115005265063

- Install Docker.  
https://www.docker.com/

### local
```
$ docker build -t pickup-news .
```

```
$ docker run \
-e PICKUPNEWS_KEYWORD=Google \
-e PICKUPNEWS_FROM=2019-12-15 \
-e PICKUPNEWS_TO=2019-12-15 \
-e PICKUPNEWS_APIKEY=Your News API key \
-e PICKUPNEWS_WEBHOOKURL=Your Slack incoming webhook url \
-it --rm --name running-pickup-news pickup-news
```

