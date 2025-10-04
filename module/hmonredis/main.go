package hmonredis

import (
    "context"
    "errors"
    "fmt"
    "log"
    "syscall"
    "time"
    "crypto/tls"
    "github.com/go-redis/redis/v8"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/module/hmonstd"
)

var redisHost string
var redisPort int
var redisPwd string
var redisDb int
var redisMaxRetries int
var redisMinRetryBackoff time.Duration
var redisMaxRetryBackoff time.Duration
var redisDialTimeout time.Duration
var redisReadTimeout time.Duration
var redisWriteTimeout time.Duration
var redisPoolSize int
var redisMinIdleConns int
var redisMaxConnAge time.Duration
var redisPoolTimeout time.Duration
var redisIdleTimeout time.Duration
var redisIdleCheckFrequency time.Duration
var redisClient *redis.Client
var redisTlsConfig *tls.Config
var ctx context.Context

func onRedisConnect(ctx context.Context, conn *redis.Conn) error {
    fmt.Println("Connection to Redis established!")
    _, err := conn.Ping(ctx).Result()
    if err != nil {
        log.Printf("Error pinging Redis on connect: %v", err)
        return err
    }
    fmt.Println("Redis server successfully pinged on connect.")
    return nil
}

func getRedisClient() *redis.Client {
    return redis.NewClient(&redis.Options{
      Addr: fmt.Sprintf(
        "%s:%d",
        redisHost,
        redisPort),
      Password: redisPwd,
      DB: redisDb,
      OnConnect: onRedisConnect,
      MaxRetries: redisMaxRetries,
      MinRetryBackoff: redisMinRetryBackoff,
      MaxRetryBackoff: redisMaxRetryBackoff,
      DialTimeout: redisDialTimeout,
      ReadTimeout: redisReadTimeout,
      WriteTimeout: redisWriteTimeout,
      PoolSize: redisPoolSize,
      MinIdleConns: redisMinIdleConns,
      MaxConnAge: redisMaxConnAge,
      PoolTimeout: redisPoolTimeout,
      IdleTimeout: redisIdleTimeout,
      IdleCheckFrequency: redisIdleCheckFrequency,
      TLSConfig: redisTlsConfig,
  })
}

func SetRaw(key string, data []uint8) error {
    result := redisClient.Set(ctx, key, data, 0)
    err := result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result = redisClient.Set(ctx, key, data, 0)
        err = result.Err()
    }
    return err
}

func SetRawWithExpired(key string, data []uint8) error {
    ttl := time.Duration(60) * time.Second
    result := redisClient.Set(ctx, key, data, ttl)
    err := result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result = redisClient.Set(ctx, key, data, ttl)
        err = result.Err()
    }
    return err
}

func SetStr(key string, val string) error {
    result := redisClient.Set(ctx, key, val, 0)
    err := result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result = redisClient.Set(ctx, key, val, 0)
        err = result.Err()
    }
    return err
}

func SetInt(key string, val int) error {
    result := redisClient.Set(ctx, key, val, 0)
    err := result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result = redisClient.Set(ctx, key, val, 0)
        err = result.Err()
    }
    return err
}

func Get(key string) (string, error) {
    var err error
    result := redisClient.Get(ctx, key)
    err = result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result := redisClient.Get(ctx, key)
        err = result.Err()
    }
    if err != nil {
        if err == redis.Nil {
            return "", nil
        }
        return "", err
    }
    var out string
    out, _ = result.Result()
    if err != nil {
        if err == redis.Nil {
            return "", nil
        }
        return "", err
    }
    return out, nil
}

func Del(key string) error {
    result := redisClient.Del(ctx, key)
    err := result.Err()
    for errors.Is(err, syscall.ECONNRESET) {
        result = redisClient.Del(ctx, key)
        err = result.Err()
    }
    return err
}

func Subscribe(c context.Context, key string) *redis.PubSub {
    return redisClient.Subscribe(c, key)
}

func SubscriberReceiveMessage(subscriber *redis.PubSub) (*redis.Message, error) {
    return subscriber.ReceiveMessage(ctx)
}

func init() {

    redisHost = hmonenv.GetRedisHost()
    redisPort = hmonenv.GetRedisPort()
    redisPwd = hmonenv.GetRedisPwd()
    redisDb = hmonenv.GetRedisDb()
    redisMaxRetries = hmonenv.GetRedisMaxRetries()
    redisMinRetryBackoff = hmonenv.GetRedisMinRetryBackoff()
    redisMaxRetryBackoff = hmonenv.GetRedisMaxRetryBackoff()
    redisDialTimeout = hmonenv.GetRedisDialTimeout()
    redisReadTimeout = hmonenv.GetRedisReadTimeout()
    redisWriteTimeout = hmonenv.GetRedisWriteTimeout()
    redisPoolSize = hmonenv.GetRedisPoolSize()
    redisMinIdleConns = hmonenv.GetRedisMinIdleConns()
    redisMaxConnAge = hmonenv.GetRedisMaxConnAge()
    redisPoolTimeout = hmonenv.GetRedisPoolTimeout()
    redisIdleTimeout = hmonenv.GetRedisIdleTimeout()
    redisIdleCheckFrequency = hmonenv.GetRedisIdleCheckFrequency()
    redisTlsConfig = hmonenv.GetRedisTlsConfig()

    redisClient = getRedisClient()
    err := hmonstd.ChkTcp(redisHost, redisPort)
    if err != nil {
        panic(fmt.Sprintf("Connection to Redis ERROR; %s", err.Error()))
    }
    ctx = context.Background()
}
