package rcache

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type CacheConfig struct {
	Host string
	Password string
	Port string
	Database string
}

var CTX = context.Background()
var RS *redis.Client

func Connect(cfg *CacheConfig) {
	dbnum, e := strconv.Atoi(cfg.Database)

	if e != nil {
		log.Fatal("Failed to connect to redis host")
	}

	if rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB: dbnum,
	}); rdb == nil {
		log.Fatal("Failed to connect to redis host")
	} else {
		RS = rdb
	}
}