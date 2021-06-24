package storage

import (
	"encoding/json"
	"errors"
	"github.com/1r0npipe/shortener-web-links/internal/model"
	"github.com/go-redis/redis"
	"time"
)

type ClientRedis struct {
	red *redis.Client
}

type Options struct {
	Addr     string
	Password string
	DB       int
}

var (
	DefaultOptions = Options{
		Addr:     "localhost:6379",
		Password: "RedisPassword",
		DB:       0,
	}
	ErrNotFound = errors.New("item not found in Redis")
)

func (c *ClientRedis) New(option Options) (ClientRedis, error) {
	result := ClientRedis{}
	if option.Addr == "" {
		option.Addr = DefaultOptions.Addr
	}
	if option.Password == "" {
		option.Password = DefaultOptions.Password
	}
	client := redis.NewClient(&redis.Options{
		Addr:     option.Addr,
		Password: option.Password,
		DB:       option.DB,
	})

	err := client.Ping().Err()
	if err != nil {
		return result, ErrNotFound
	}

	result.red = client
	return result, nil
}

func (c *ClientRedis) Get(key string) (*model.Item, error) {
	result := model.Item{}
	data, err := c.red.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *ClientRedis) Put(key string, m model.Item) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = c.red.Set(key, string(data), time.Duration(m.TTL)*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientRedis) Delete(key string, m model.Item) error {

	return nil
}
