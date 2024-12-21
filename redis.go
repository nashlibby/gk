package gk

import (
	"context"
	"encoding/json"
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
	jData, err := json.Marshal(val)
	if err != nil {
		return err
	}
	err = r.Client.Set(r.Ctx, key, string(jData), expiration).Err()
	return
}

func (r *GRedis) Get(key string) (val string, err error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *GRedis) Has(key string) (int64, error) {
	count, err := r.Client.Exists(r.Ctx, key).Result()
	return count, err
}

func (r *GRedis) Del(key string) (int64, error) {
	count, err := r.Client.Del(r.Ctx, key).Result()
	return count, err
}

func (r *GRedis) Hset(key string, values ...interface{}) error {
	err := r.Client.HSet(r.Ctx, key, values).Err()
	return err
}

func (r *GRedis) Hget(key, field string) (string, error) {
	result, err := r.Client.HGet(r.Ctx, key, field).Result()
	return result, err
}

func (r *GRedis) HgetAll(key, field string) (map[string]string, error) {
	result, err := r.Client.HGetAll(r.Ctx, key).Result()
	return result, err
}
