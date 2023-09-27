package entity

import "time"

type Location struct {
	ID        int
	Lat       float64
	Long      float64
	UserID    int
	Timestamp time.Time
	Accuracy  int
}
