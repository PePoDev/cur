// Package redis contain fuction to connect redis database and utility function
package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}

// Get key and return value string with error
func Get(key string) (string, error) {
	value, err := rdb.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("key %v does not exist\n", key)
	} else if err != nil {
		panic(err)
	}

	return value, err
}

// Set key and value to redis and retrun string
func Set(key, value string) error {
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}
