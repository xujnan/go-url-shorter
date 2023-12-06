package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 封装redis客户端
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// CacheDuration 实际使用中不应设置到期时间，应设置LRU策略，并在清除时存入数据库
const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init redis: %v", err))
	}

	fmt.Printf("\nRedis started srccrssfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

// SaveUrlMapping 数据库API设计，存储原URL到短url的映射
func SaveUrlMapping(shortUrl, originalUrl, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url Error: %v", err))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieveInitialUrl Error: %v", err))
	}

	return result
}
