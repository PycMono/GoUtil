package redisUtil

/*
* 测试redis连接池
*
*
*
 */
import (
	"fmt"
	"moqikaka.com/Test/src/Model/redisUtil"
	"testing"
)

func TestSend(context *testing.T) {
	redisConfig := redisUtil.NewRedisConfig("10.1.0.40:6380", "pwd_123", 0, 100, 100, 10000, 0)
	redisPool := NewRedisPool("1", redisConfig)
	redisPool.TestConn()
	redisPool.Set("UserID", 200)
	for i := 0; i <= 1000; i++ {
		userID, _, err := redisPool.GetKey("UserID")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(userID)
	}
}
