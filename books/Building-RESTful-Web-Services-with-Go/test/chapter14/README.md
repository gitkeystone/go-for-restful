

```bash
go get -u "github.com/gorilla/mux"
go get -u "github.com/gorilla/sessions"


openssl rand -base64 32
export SESSION_SECRET="pyX5jjHtg0+yYiFgumO/1KfC5/ahKqDF0s/UZ1avUgQ="
```


# 安装 Redis
```bash
docker run --name some-redis -p 6379:6379 -d redis
docker exec -i -t some-redis redis-cli

KEYS *
SET topic async


go get "gopkg.in/boj/redistore.v1"

```

# JWT
```bash
go get -u "github.com/dgrijalva/jwt-go/v4"
```
