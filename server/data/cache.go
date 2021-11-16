package data

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RS *redis.Client


func ConnectCache() {
	dbn, e := strconv.Atoi(os.Getenv("REDIS_DB"))

	if e != nil {
		log.Fatal(e.Error())
	}

	if rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: dbn,
	}); rdb == nil {
		log.Fatal("Failed to connect to redis host")
	} else {
		RS = rdb
		fmt.Println("Successfully connected to Redis server")
	}
}