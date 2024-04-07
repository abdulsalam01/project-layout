package app

import (
	"context"
	"fmt"
	"time"

	"github.com/api-sekejap/config"
	"github.com/api-sekejap/internal/constant"
	db "github.com/api-sekejap/pkg/database"
	redis "github.com/api-sekejap/pkg/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type BaseAppInitializer struct {
	db.DatabaseHelper
	redis.RedisHelper
	redis.RedisLockerHelper
}

// Base initializer function to return base of application requirements.
// This function will contains:
// 1. Database.
// 2. Logging.
// 3. Redis.
// 4. ErrorWrapper.
func Initializer(ctx context.Context, config *config.Config) (BaseAppInitializer, error) {
	var (
		initializer BaseAppInitializer
		err         error
	)

	/*
	 * Configuration layer.
	 * Database section.
	 */
	dbUrlConnection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.Extras[constant.DatabaseSSLMode].(string),
	)

	dbDurationMax, err := time.ParseDuration(config.Database.Extras[constant.DatabaseTimeout].(string))
	if err != nil {
		logrus.Fatalf("Unable to parse duration of connection config pool: %v\n", err)
		return initializer, err
	}

	dbConfig, err := pgxpool.ParseConfig(dbUrlConnection)
	if err != nil {
		logrus.Fatalf("Unable to parse connection config pool: %v\n", err)
		return initializer, err
	}

	// Setup base db connection.
	dbConfig.MaxConns = int32(config.Database.Extras[constant.DatabaseMaxConnection].(int))
	dbConfig.MinConns = int32(config.Database.Extras[constant.DatabaseMinConnection].(int))
	dbConfig.MaxConnLifetime = dbDurationMax
	dbConfig.ConnConfig.ConnectTimeout = dbDurationMax
	dbConfig.ConnConfig.Tracer = &dbQueryTracer{logrus.StandardLogger()} // Tracer settings.

	// Tie with database pool configuration.
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		logrus.Fatalf("Unable to create connection pool: %v\n", err)
		return initializer, err
	}
	defer dbPool.Reset()

	/*
	 * Configuration layer.
	 * Logger section.
	 */

	/*
	 * Configuration layer.
	 * Redis section.
	 */
	redisMemory := redis.New(config.MemoryCache)
	redisLocker := redis.NewLocker(redisMemory)

	/*
	 * Configuration layer.
	 * Errorwraper section.
	 */

	// Initializer all here.
	initializer = BaseAppInitializer{
		db.DatabaseHelper{Database: dbPool},
		redis.RedisHelper{Memory: redisMemory},
		redis.RedisLockerHelper{Locker: redisLocker},
	}
	return initializer, nil
}
