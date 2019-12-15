# pickup-news
Output news that matches the keyword.

## Usage
### Preparations
- You get NewsAPI API key.  
https://newsapi.org/

- Install Docker.  
https://www.docker.com/

### local
```
$ docker build -t pickup-news .
```

```
$ docker run \
-e PICKUPNEWS_KEYWORD="Google" \
-e PICKUPNEWS_FROM="2019-12-15" \
-e PICKUPNEWS_TO="2019-12-15" \
-e PICKUPNEWS_APIKEY="Your API Key" \
-it --rm --name running-pickup-news pickup-news
```

