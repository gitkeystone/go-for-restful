# 安装 RabbitMQ
```bash
# https://hub.docker.com/_/rabbitmq?xk=ShowRecommendedBadge&xt=Disabled
docker pull rabbitmq:4.3.0
docker run -d --hostname rabbitmq-host --name rabbitmq-server -p 5672:5672 -p 15672:15672 rabbitmq:4.3.0

docker run --rm --detach --hostname rabbitmq-host --name rabbitmq-server \
    --env RABBITMQ_DEFAULT_USER=<guest> \
    --env RABBITMQ_DEFAULT_PASS=<guest> \
    --publish 15672:15672 \
    --publish 5672:5672 \
     rabbitmq:4.3.0
```


# Go RabbitMQ Client Library
```bash
go get -u github.com/rabbitmq/amqp091-go

export AMQP_URL="amqp://guest:guest@localhost:5672"
```


# 安装 Redis
```bash
docker run --name some-redis -p 6379:6379 -d redis
docker exec -i -t some-redis redis-cli

KEYS *
SET topic async
```

# Redis client for Go
```bash
go get -u github.com/redis/go-redis/v9

"github.com/redis/go-redis/v9"

```