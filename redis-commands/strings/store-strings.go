package storeStrings

import (
	"fmt"
	"redis-educative/util"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

type InsRetCmds interface {
	Set(key, value string)
	Get(key string) string
	SetEx(key, value string, duration int)
	PSetEx(key, value string, duration int)
}

// This is to make sure Redis implements all of the InsRetCmds interface functions.
var val InsRetCmds = (*Redis)(nil)

func (r *Redis) Set(key, value string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("SET", key, value))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("Set ", val)
}

func (r *Redis) Get(key string) string {
	conn := r.Pool.Get()
	defer conn.Close()
	gValue, err := redis.String(conn.Do("GET", key))
	util.FailOnError(err, "Failed to set value")
	return gValue
}

func (r *Redis) SetEx(key, value string, duration int) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("SETEX", key, duration, value))
	util.FailOnError(err, "Failed to set value with duration")
	fmt.Println("SetEx ", val)
}

func (r *Redis) PSetEx(key, value string, duration int) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("PSETEX", key, duration, value))
	util.FailOnError(err, "Failed to set value with duration")
	fmt.Println("PSetEx ", val)
}

func (r *Redis) SetNx(key, value string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("SETNX", key, value))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("SetNx ", val)
}

func (r *Redis) Strlen(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("STRLEN", key))
	util.FailOnError(err, "Failed to get length of value")
	fmt.Println("Strlen ", val)
}

func (r *Redis) MSet() {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("MSET", "ZNMD", "10", "KKKG", "8"))
	util.FailOnError(err, "Failed to set value")
	fmt.Println("MSet ", val)
}

func (r *Redis) MGet() {
	conn := r.Pool.Get()
	defer conn.Close()
	var keys []interface{}
	keys = append(keys, "ZNMD")
	keys = append(keys, "KKKG")
	val, err := redis.Values(conn.Do("MGET", keys...))
	util.FailOnError(err, "Failed to get value")
	ratingZNMD, err := redis.Int64(val[0], nil)
	util.FailOnError(err, "Failed to scan value")
	ratingKKKG, err := redis.Int64(val[1], nil)
	util.FailOnError(err, "Failed to scan value")
	fmt.Println("MGet ", ratingZNMD, ratingKKKG)
}
