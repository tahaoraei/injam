package main

import (
	"context"
	"fmt"
	"injam/adapter/redis"
	"injam/pkg/config"
	"time"
)

func main() {
	cfg := config.Load("config.yml")
	rdb := redis.New(cfg.Redis)
	if err := SaveUserLocation(&rdb, "2", 55.656, 33.4354, time.Now().UnixMicro()); err != nil {
		fmt.Println(err)
	}

	location, timestamp, _ := GetUserLocation(&rdb, "2")
	fmt.Println(location, timestamp)

	if err := SaveUserLocation(&rdb, "2", 55.657, 33.4351, time.Now().UnixMicro()); err != nil {
		fmt.Println(err)
	}

	location, timestamp, _ = GetUserLocation(&rdb, "2")
	fmt.Println(location, timestamp)

}

func SaveUserLocation(rdb *redis.Adapter, userID string, longitude float64, latitude float64, timestamp int64) error {
	ctx := context.Background()

	// Set the user's location and timestamp in the Redis hash
	err := rdb.Client().HSet(ctx, "user:"+userID, "location", fmt.Sprintf("%f,%f", longitude, latitude)).Err()
	if err != nil {
		return err
	}

	err = rdb.Client().HSet(ctx, "user:"+userID, "timestamp", timestamp).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserLocation(rdb *redis.Adapter, userID string) (string, int64, error) {
	ctx := context.Background()

	// Retrieve the user's location and timestamp from the Redis hash
	location, err := rdb.Client().HGet(ctx, "user:"+userID, "location").Result()
	if err != nil {
		return "", 0, err
	}

	timestamp, err := rdb.Client().HGet(ctx, "user:"+userID, "timestamp").Int64()
	if err != nil {
		return "", 0, err
	}

	return location, timestamp, nil
}
