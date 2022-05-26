package rediscommands

import (
	"fmt"
	"redis-educative/util"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

type UtilityCommands interface {
	Keys()
}

func (r *Redis) Keys() {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Strings(conn.Do("KEYS", "*"))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("Keys ", val)
}

func (r *Redis) Incr(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("INCR", key))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("INCR : ", val)
}

func (r *Redis) IncrBy(key string, count int64) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("INCRBY", key, count))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("INCRBY : ", val)
}

func (r *Redis) IncrByFloat(key string, count float64) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Float64(conn.Do("INCRBYFLOAT", key, count))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("INCRBYFLOAT : ", val)
}

func (r *Redis) Decr(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("DECR", key))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("DECR : ", val)
}

func (r *Redis) DecrBy(key string, count int64) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("DECRBY", key, count))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("DECRBY : ", val)
}

func (r *Redis) Del(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("DEL", key))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("Del : ", val)
}

func (r *Redis) FlushAll() {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("FLUSHALL"))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("FLUSHALL ", val)
}

func (r *Redis) Append(key, value string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("APPEND", key, value))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("APPEND : ", val)
}