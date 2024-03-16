package activity

import "time"

type Activity struct {
	Sport     string
	Date      time.Time
	Title     string
	Time      time.Duration
	Distance  float64
	Elevation int
	Effort    int
}
