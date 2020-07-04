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
$ make build
```
TODO: Does not work properly
TODO: fix stdin
```
$ sam local invoke PickupNewsFunction -n env.json -e event.json
```

### release
When you push the tag(v*), the `GitHub Actions` execute builds and releases.  
https://github.com/KirariGlock/pickup-news/releases  

```
$ git tag v0.0.1
$ git push origin v0.0.1
```
