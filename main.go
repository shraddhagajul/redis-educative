package main

import (
	"fmt"
	"redis-educative/redis-commands/lists"
	rStrings "redis-educative/redis-commands/strings"
	utilityCmds "redis-educative/redis-commands/utility-commands"
	"redis-educative/redis-pool"
)

func main() {

	pool := redispool.BuildPool()
	redispool.Ping(pool.Get())

	//Insert Retrieve commands
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

	//Utility Commands
	utilityRpool := utilityCmds.Redis{Pool: pool}
	utilityRpool.Keys()
	utilityRpool.Incr("KKKG")
	utilityRpool.IncrBy("KKKG", 2)
	utilityRpool.IncrByFloat("maths", 2.2)
	utilityRpool.Decr("KKKG")
	utilityRpool.DecrBy("KKKG", 1)
	utilityRpool.Del("planet")
	utilityRpool.FlushAll()
	rPool.Set("flowers", "rose")
	utilityRpool.Append("flowers",",daffodiles")

	//List commands
	key := "wonders"
	listPool := lists.Redis{Pool: pool}
	values := []string{"taj mahal", "leaning tower of pisa"}
	listPool.Lpush(key, values)
	listPool.LRange(key)
	listPool.Lpop(key)
	listPool.Rpush(key, values)
	listPool.Rpop(key)
	listPool.Llen(key)
	listPool.LIndex(key, 0)
	listPool.LSET(key, 0, "statue of liberty")
	listPool.LpushX(key, "leaning tower of pisa")
	listPool.LInsert(key, "leaning tower of pisa", "wall of china")
	listPool.LInsertAfter(key, "wall of china", "Chichen Itza")

}
