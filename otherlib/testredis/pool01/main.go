package main

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

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			*/
			return c, err
		},
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
