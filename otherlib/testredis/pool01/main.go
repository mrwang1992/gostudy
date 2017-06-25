package main

// demo for using redisgo pool
// from blog: http://blog.csdn.net/u010471121/article/details/52779039

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool          *redis.Pool
	redisServer   = flag.String("redisServer", ":6700", "")
	redisPassword = flag.String("redisPassword", "", "")
)

// init redis pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		// max redis conn idle
		MaxIdle: 3,
		// max redis conn num
		MaxActive: 5,
		// max idle wait time for redis conn
		IdleTimeout: 240 * time.Second,
		// make redis conn function
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			// use for auth
			/*
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			*/
			return c, err
		},
		// test conn is alive function
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword)

	conn := pool.Get()
	defer conn.Close()

	v, err := conn.Do("SET", "pool", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(conn.Do("GET", "pool"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
