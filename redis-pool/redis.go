package redispool

import (
	"fmt"
	"os"
	"redis-educative/util"

	"github.com/gomodule/redigo/redis"
)

func BuildPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			util.FailOnError(err, "Failed to connect to redis")
			if err != nil {
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func Ping(conn redis.Conn) {
	pong, err := redis.String(conn.Do("PING"))
	util.FailOnError(err, "Failed to get response for PING")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(pong)
}
