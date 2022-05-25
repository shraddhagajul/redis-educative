package main

import (
	"fmt"
	rStrings "redis-educative/redis-commands/strings"
	"redis-educative/redis-pool"
)

func main() {

	pool := redispool.BuildPool()
	redispool.Ping(pool.Get())

	rPool := rStrings.Redis{Pool: pool}
	rPool.Set("planet", "earth")
	value := rPool.Get("planet")
	fmt.Println("key : planet, value : ", value)
	rPool.SetEx("animal", "elephant", 15)

	rPool.PSetEx("stone", "diamond", 2000)
	rPool.SetNx("animal", "giraffe")
	rPool.SetNx("city", "Mumbai")
	rPool.Strlen("animal")
	rPool.MSet()
	rPool.MGet()

}
