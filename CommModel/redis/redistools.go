package redistools

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisDataStore struct {
	RedisHost 	string
	RedisDB    	string
	RedisPwd	string
	TimeOut		int64
	RedisPool	*redis.Pool
}

func (r *RedisDataStore) NewPool() *redis.Pool {
	return &redis.Pool{
		Dial: 			r.RedisConnect,
		MaxIdle:		10,
		MaxActive: 		0,
		IdleTimeout:  	1 * time.Second,
		Wait:			true,
	}
}

func (r *RedisDataStore) RedisConnect() (redis.Conn, error) {
		c, err := redis.Dial("tcp", r.RedisHost)

		if err != nil{
			return nil,err
		}

		_, err = c.Do("AUTH", r.RedisPwd)

		if err != nil {
			return nil, err
		}

		_, err = c.Do("SELECT", r.RedisDB)

		if err != nil {
			return nil, err
		}

		redis.DialConnectTimeout(time.Duration(r.TimeOut))
		redis.DialReadTimeout(time.Duration(r.TimeOut))
		redis.DialWriteTimeout(time.Duration(r.TimeOut))

		return c,nil
}

func (r *RedisDataStore) Get(k string) (interface{}, error) {
		c := r.RedisPool.Get()
		defer c.Close()
		v, err := c.Do("GET", k)
		if err != nil {
			return nil, err
		}

		return v, nil
}

func (r *RedisDataStore) Set(k, v string) error{
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)

	return err
}

func (r *RedisDataStore) SetEx(k string, v interface{}, ex int64) error{
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v, "EX", ex)

	return err
}


func (r *RedisDataStore) HSet(f, k, v string) error{
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("HSET", f, k, v)

	return err
}

func (r *RedisDataStore) HSetEx(f, k string, v interface{}, ex int64) error{
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("HSET", f, k, v)
	_, err = c.Do("Expire", f, ex)
	return err
}

func (r *RedisDataStore) HGet(f, k string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	v, err := c.Do("HGET", f, k)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *RedisDataStore) HGetAll(f string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	v, err := c.Do("HGETALL", f)
	if err != nil {
		return nil, err
	}

	return v, nil
}

