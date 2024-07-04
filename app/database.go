package app

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Redis   *redis.Client
	Mongodb *mongo.Client
	MySql   *gorm.DB
}

func (app *App) initDB() *DB {
	db := &DB{}

	db.Redis = newRedisClient(app.ctx, app.cfg)
	// db.Mongodb = newMongodbClient(app.ctx, app.cfg)
	db.MySql = newMySqlClient(app.cfg)

	return db
}

func newRedisClient(ctx context.Context, cfg *AppConfig) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisURL)

	if err != nil {
		log.Panicln("Redis parse url error", err)
	}

	client := redis.NewClient(opts)

	if err = client.Ping(ctx).Err(); err != nil {
		log.Panicln("Redis ping error", err)
	}

	fmt.Println("Redis connected")

	return client
}

func newMongodbClient(ctx context.Context, cfg *AppConfig) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongodbURI))

	if err != nil {
		log.Panicln("Mongodb connect error", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Panicln("Mongodb ping error", err)
	}

	fmt.Println("Mongodb connected")

	return client
}

func newMySqlClient(cfg *AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MySqlDNS), &gorm.Config{})

	if err != nil {
		log.Panicln("MySql connect error", err)
	}

	fmt.Println("MySql connected")

	return db
}
