package rediscommands

import "github.com/gomodule/redigo/redis"

type Redis struct {
	Pool *redis.Pool
}
