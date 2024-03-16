package statistic

import (
	"math"
	"strava-booster/activity"
	"time"
)

func CountUniqueDays(activities []activity.Activity) int {
	uniqueDays := make(map[time.Time]bool)

	for _, activity := range activities {
		uniqueDays[activity.Date.Truncate(24*time.Hour)] = true
	}

	return len(uniqueDays)
}

func CalculateTotalDuration(activities []activity.Activity) time.Duration {
	var totalDuration time.Duration

	for _, activity := range activities {
		totalDuration += activity.Time
	}

	return totalDuration
}

func CalculateTotalDistance(activities []activity.Activity) float64 {
	var totalDistance float64

	for _, activity := range activities {
		totalDistance += activity.Distance
	}

	return totalDistance
}

func FindShortestRun(activities []activity.Activity) activity.Activity {
	var shortestRun activity.Activity
	shortestDistance := math.MaxFloat64

	for _, activity := range activities {
		if activity.Distance < shortestDistance {
			shortestRun = activity
			shortestDistance = activity.Distance
		}
	}

	return shortestRun
}

func FindFarthestRun(activities []activity.Activity) activity.Activity {
	var farthestRun activity.Activity
	farthestDistance := 0.0

	for _, activity := range activities {
		if activity.Distance > farthestDistance {
			farthestRun = activity
			farthestDistance = activity.Distance
		}
	}

	return farthestRun
}
