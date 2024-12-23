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

func (r *GRedis) Set(key string, value interface{}, expiration time.Duration) (err error) {
	err = r.Client.Set(r.Ctx, key, value, expiration).Err()
	return
}

func (r *GRedis) SetJson(key string, value interface{}, expiration time.Duration) (err error) {
	jData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.Client.Set(r.Ctx, key, string(jData), expiration).Err()
	return
}

func (r *GRedis) Get(key string) (string, error) {
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

func (r *GRedis) HgetAll(key string) (map[string]string, error) {
	return r.Client.HGetAll(r.Ctx, key).Result()
}

// 数据入队列
func (r *GRedis) QueuePush(queueName string, value interface{}) (int64, error) {
	return r.Client.LPush(r.Ctx, queueName, value).Result()
}

// json数据入队列
func (r *GRedis) QueuePushJson(queueName string, value interface{}) (int64, error) {
	jData, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}
	return r.Client.LPush(r.Ctx, queueName, jData).Result()
}

// 数据出队列
func (r *GRedis) QueuePop(queueName string, timeout time.Duration) ([]string, error) {
	return r.Client.BRPop(r.Ctx, timeout, queueName).Result()
}

// 获取队列所有数据
func (r *GRedis) QueueList(queueName string) ([]string, error) {
	return r.Client.LRange(r.Ctx, queueName, 0, -1).Result()
}

// 获取队列长度
func (r *GRedis) QueueLength(queueName string) (int64, error) {
	return r.Client.LLen(r.Ctx, queueName).Result()
}

// 清空队列
func (r *GRedis) QueueClear(queueName string) error {
	return r.Client.LTrim(r.Ctx, queueName, 1, 0).Err()
}
