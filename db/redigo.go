package db

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/somewhere/config"
)

var RedisDb redis.Conn

func InitRedisDatabase() error {
	RedisDb, err := redis.Dial("tcp", config.Config.DbConfig.RedisAddress)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}

	_, err = RedisDb.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(RedisDb.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	return nil
}
