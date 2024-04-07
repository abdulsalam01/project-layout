package redis

import "github.com/redis/go-redis/v9"

func New(config MemoryCache) redisResources {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		Username: config.Username,
		DB:       0, // Use default DB.
	})

	return rdb
}
