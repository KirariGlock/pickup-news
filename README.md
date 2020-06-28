# pickup-news
Output news that matches the keyword.  
Notify if the value is less than the value set in PICKUPNEWS_NOTICELOWERLIMIT.

## Usage
### Preparations
- You get NewsAPI API key.  
https://newsapi.org/

- You get Slack incoming webhook url.
https://slack.com/intl/ja-jp/help/articles/115005265063

- Install Docker.  
https://www.docker.com/

- Install AWS SAM
https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html

### local
```
$ docker build -t pickup-news .
```

Currently it cannot work in local environment.  
TODO: Allow it to run on AWS SAM.  
```
$ docker run \
-e PICKUPNEWS_KEYWORD=Google \
-e PICKUPNEWS_FROM=2019-12-15 \
-e PICKUPNEWS_TO=2019-12-15 \
-e PICKUPNEWS_APIKEY=Your News API key \
-e PICKUPNEWS_WEBHOOKURL=Your Slack incoming webhook url \
-e PICKUPNEWS_NOTICELOWERLIMIT=0 \
-it --rm --name running-pickup-news pickup-news
```

### release
When you push the tag(v*), the `GitHub Actions` execute builds and releases.  
https://github.com/KirariGlock/pickup-news/releases  

```
$ git tag v0.0.1
$ git push origin v0.0.1
```
