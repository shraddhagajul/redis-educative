package lists

import (
	"fmt"
	"redis-educative/util"

	"github.com/gomodule/redigo/redis"
)
type Redis struct {
	Pool *redis.Pool
}

func (r *Redis) Lpush(key string, values []string) {
	conn := r.Pool.Get()
	defer conn.Close()
	for _, element := range values {
	val, err := redis.Int64(conn.Do("LPUSH", key, element))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("Lpush : ", val)
	}
}

func (r *Redis) LRange(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LRANGE : ", val)	
}

func (r *Redis) Lpop(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("LPOP", key))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LPOP : ", val)	
}

func (r *Redis) Rpush(key string, values []string) {
	conn := r.Pool.Get()
	defer conn.Close()
	for _, element := range values {
	val, err := redis.Int64(conn.Do("RPUSH", key, element))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("Lpush : ", val)
	}
}

func (r *Redis) Rpop(key string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("RPOP", key))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("RPOP : ", val)	
}