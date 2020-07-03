package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义redis
var pool = &redis.Pool{
	IdleTimeout: 300 * time.Second,
	Dial: func() (conn redis.Conn, err error) {
		return redis.Dial(networkType, address, redis.DialPassword(passwd))
	},
}
