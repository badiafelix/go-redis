package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	return client
}

func setData(rdb *redis.Client, key string, data string, ttl time.Duration) *redis.StatusCmd {
	exec := rdb.Set(context.Background(), key, data, ttl)
	if err := exec.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
	}
	log.Println("set operation success")
	return exec
}

func getData(rdb *redis.Client, key string) string {
	exec := rdb.Get(context.Background(), key)
	if err := exec.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
	}
	res, err := exec.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
	}
	log.Println("get operation success. result:", res)
	return res
}

func main() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)
	key := "name"
	data := "badai engineerr"
	ttl := time.Duration(100) * time.Second

	//insert data
	insert := setData(rdb, key, data, ttl)
	fmt.Println(insert)
	//ambil data
	fmt.Println(getData(rdb, key))

}
