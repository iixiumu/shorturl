package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     128,
		MaxActive:   1024,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "127.0.0.1:6379") },
	}
}

type RedisIncr struct {
}

func (RedisIncr) GenID() uint64 {
	conn := RedisPool.Get()
	defer conn.Close()
	id, _ := redis.Uint64(conn.Do("INCR", "uid"))
	return id
}

func (RedisIncr) Long2Short(longUrl string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	url, err := redis.String(conn.Do("GET", longUrl))
	return url, err
}

func (RedisIncr) Short2Long(shortUrl string) (string, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	url, err := redis.String(conn.Do("GET", shortUrl))
	return url, err
}

func (RedisIncr) SetLong2Short(longUrl, shortUrl string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("SET", longUrl, shortUrl, "EX", 10))
	return err
}

func (RedisIncr) SetShort2Long(shortUrl, longUrl string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("SET", shortUrl, longUrl, "EX", 10))
	return err
}

func (RedisIncr) Expire(key string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("EXPIRE", key, 10))
	return err
}
