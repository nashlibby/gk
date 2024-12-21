package gk

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type GRedis struct {
	Client *redis.Client
	Ctx    context.Context
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

func NewRedisClient(config RedisConfig) *GRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	return &GRedis{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func (r *GRedis) Set(key string, val interface{}, expiration time.Duration) (err error) {
	err = r.Client.Set(r.Ctx, key, val, expiration).Err()
	return
}

func (r *GRedis) Get(key string) (val interface{}, err error) {
	val, err = r.Client.Get(r.Ctx, key).Result()
	return
}

func (r *GRedis) Has(key string) error {
	_, err := r.Client.Exists(r.Ctx, key).Result()
	return err
}

func (r *GRedis) Del(key string) error {
	_, err := r.Client.Del(r.Ctx, key).Result()
	return err
}

func (r *GRedis) Hset(key string, values ...interface{}) error {
	_, err := r.Client.HSet(r.Ctx, key, values).Result()
	return err
}

func (r *GRedis) Hget(key, field string) error {
	_, err := r.Client.HGet(r.Ctx, key, field).Result()
	return err
}

func (r *GRedis) HgetAll(key, field string) error {
	_, err := r.Client.HGetAll(r.Ctx, key).Result()
	return err
}
