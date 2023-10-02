package redislocation

import (
	"context"
	"fmt"
	"injam/adapter/redis"
	"time"
)

type DB struct {
	adapter redis.Adapter
}

func New(adapter redis.Adapter) DB {
	return DB{adapter: adapter}
}

func (db DB) SaveUserLocation(userID string, od string) error {
	if _, err := db.adapter.Client().Set(context.TODO(), "user:"+userID, od, time.Hour*3).Result(); err != nil {
		return err
	}
	return nil
}

func (db DB) GetUserLocation(userID string) (string, error) {
	od, err := db.adapter.Client().Get(context.TODO(), "user"+userID).Result()
	fmt.Println("user" + userID)
	fmt.Println("od ", od, "   err in get user:", userID, err)
	if err != nil {
		return "", err
	}
	return od, nil
}
