package redisUtil

/**
* 用go语言对redis连接池进行了封装，GetConn之后，需要先调用下TestConn哈，防止断开连接的情况
* 外部获取conn之后再进行第二次封装的时候，切记要调用Close函数，防止连接未被放回连接池的情况，导致连接过多
*
*
 */

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	. "moqikaka.com/Test/src/Model/redisUtil"
	"time"
)

// redis连接池对象
type RedisPool struct {
	// 名字
	name string

	// 连接地址
	address string

	// 连接池对象
	pool redis.Pool
}

// 获取连接池名字
// 返回值：
// 连接池名字
func (this *RedisPool) GetName() string {
	return this.name
}

// 获取连接池地址
// 返回值：
// 连接池地址
func (this *RedisPool) GetAddress() string {
	return this.address
}

// 获取连接对象，在使用后需要调用Close方法，将连接放入连接池，方便继续使用
// 返回值：
// 连接对象
func (this *RedisPool) GetConn() redis.Conn {
	return this.pool.Get()
}

// 测试连接是否生效
// 返回值：
// error：如果error为nil，表示连接生效的，否则未生效
func (this *RedisPool) TestConn() error {
	conn := this.GetConn()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		return err
	}

	return nil
}

// 判断指定的key是否存在
// 参数：
// key：指定的key值
// 返回值：
// exists：true：存在，反之不存在
// err：错误信息
func (this *RedisPool) Exists(key string) (exists bool, err error) {
	conn := this.GetConn()
	defer conn.Close()

	var result int = 0
	if result, err = redis.Int(conn.Do("EXISTS", key)); err != nil {
		return
	}

	exists = result == 1

	return
}

// 判断指定的key是否存在
// 参数：
// key：指定的key值
// 返回值：
// value：返回值
// exists：true：存在，反之不存在
// err：错误信息
func (this *RedisPool) GetKey(key string) (value string, exists bool, err error) {
	conn := this.GetConn()
	defer conn.Close()

	value, err = redis.String(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}

		return
	}

	exists = true

	return
}

// 设置值
// 参数：
// key：指定的key值
// value：待设置的值
// 返回值：
// ok：是否设置成功
// err：错误对象
func (this *RedisPool) Set(key, value interface{}) (bool, error) {
	conn := this.GetConn()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("Set", key, value))

	return ok, err
}

//----------------------------------------其余的方法可以自己封装----------------------------------

//----------------------------------------创建连接池对象-----------------------------------------

// 创建redis池信息
// 参数：
// _name：连接名字
// redisConfig：redis配置文件
// 返回值：
// RedisPool：redis连接池对象
func NewRedisPool(_name string, redisConfig *RedisConfig) *RedisPool {
	_pool := redis.Pool{
		MaxActive:   redisConfig.MaxActive,
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: redisConfig.IdleTimeout,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0, 4)
			options = append(options, redis.DialConnectTimeout(redisConfig.DialConnectTimeout))
			options = append(options, redis.DialPassword(redisConfig.Password))
			options = append(options, redis.DialDatabase(redisConfig.Database))
			c, err := redis.Dial("tcp", redisConfig.ConnectionString, options...)
			if err != nil {
				return nil, fmt.Errorf("Dial failed, err:%s", err)
			}

			return c, err
		},
	}

	return &RedisPool{
		name:    _name,
		address: redisConfig.ConnectionString,
		pool:    _pool,
	}
}
