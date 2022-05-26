package lists

import (
	"fmt"
	"redis-educative/util"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) Llen(list string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("LLEN", list))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LLEN : ", val)
}

func (r *Redis) LIndex(list string, index int64 ) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("LINDEX", list, index))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LINDEX : ", val)
}

func (r *Redis) LSET(list string, index int64, value string ) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("LSET", list, index, value))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LSET : ", val)
}

func (r *Redis) LpushX(key string, element string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("LPUSHX", key, element))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("Lpushx : ", val)
}

func (r *Redis) LInsert(key string, beforeValue string, element string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("LINSERT", key, "before", beforeValue, element))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LInsert : ", val)
}

func (r *Redis) LInsertAfter(key string, afterValue string, element string) {
	conn := r.Pool.Get()
	defer conn.Close()
	val, err := redis.Int64(conn.Do("LINSERT", key, "after", afterValue, element))
	util.FailOnError(err, "Failed to store slice values")
	fmt.Println("LInsertAfter : ", val)
}