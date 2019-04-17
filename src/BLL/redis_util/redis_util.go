package redis_util

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Temp() {

	temp1 := make([]int, 0)
	temp1 = append(temp1, 10)
	temp1 = append(temp1, 20)

	delete(temp1, 20)

	return
	clusterSlots := func() ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			{
				Nodes: []redis.ClusterNode{
					{
						Addr: "10.1.0.40:6380", // master
					},
					{
						Addr: "10.1.0.40:6380", // 1st slave
					},
					{
						Addr: "10.1.0.40:6380", // master
					},
					{
						Addr: "10.1.0.40:6380", // 1st slave
					}},
			},
		}

		return slots, nil
	}

	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true, //允许将只读命令路由到随机主节点或从节点,它自动启用ReadOnly。
		Password:      "pwd_123",
	})

	redisdb.Ping()
	pipe := redisdb.Pipeline()
	pipe.Select(1)
	_, err := pipe.Exec()
	if err != nil {
		fmt.Println(err)
	}

	err = redisdb.ReloadState()
	if err != nil {
		panic(err)
	}

	//statusCmd := redisdb.Set("Name1", 100000, 1)
	IntCmd := redisdb.Del("Name1")
	fmt.Println(IntCmd.Result())
	fmt.Println(redisdb.Get("Name").Val())
	fmt.Println("-------------------------------")
	fmt.Println(redisdb.Get("Name").Val())
	fmt.Println("-------------------------------")
}
