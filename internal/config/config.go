package config

import (
	"errors"
	"shortify/internal/repositories/mongo"
	"shortify/internal/repositories/redis"
)

// Config - a struct that holds a redis client
type Config struct {
	Redis *redis.RedisInternal
	Mongo *mongo.MongoInternal
}

// NewConfig - a function that returns a new Config struct
func NewConfig() (*Config, error) {

	cfg := new(Config)

	err := cfg.newClientRedis()
	if err != nil {
		return cfg, err
	}

	err = cfg.newClientMongo()
	if err != nil {
		return cfg, err
	}

	return cfg, err
}

// CloseAll - a function that closes all connections
func (cfg *Config) CloseAll() {
	if cfg.Redis != nil {
		cfg.Redis.Redis.Close()
	}
}

// newClientMongo initializes the MongoDB client
func (cfg *Config) newClientMongo() error {

	client, err := mongo.NewMongoInternal()
	if err != nil {
		return errors.New("Error creating mongo client: " + err.Error())
	}

	cfg.Mongo = client
	return nil
}

// newClientRedis is a function that returns a new Redis client
func (cfg *Config) newClientRedis() error {

	r, err := redis.NewRedisInternal()
	if err != nil {
		return errors.New("Error creating redis client: " + err.Error())
	}

	cfg.Redis = r

	return nil
}
