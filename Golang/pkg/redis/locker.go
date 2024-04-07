package redis

import (
	"github.com/bsm/redislock"
)

func NewLocker(redis redisResources) redisLockerResources {
	// Create a new lock client.
	locker := redislock.New(redis)
	return locker
}
