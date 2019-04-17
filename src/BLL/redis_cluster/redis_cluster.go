package redis_cluster

import (
	"time"

	"github.com/gitstliu/go-redis-cluster"
)

func tt() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"10.1.0.40:6380", "10.1.0.40:6380", "10.1.0.40:6380"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		return
	}

	cluster.Do("SET", "foo", "bar")
}
