package statistic

import (
	"math"
	a "strava-booster/activity"
	"time"
)

func CountUniqueDays(activities []a.Activity) int {
	uniqueDays := make(map[time.Time]bool)

	for _, activity := range activities {
		uniqueDays[activity.Date.Truncate(24*time.Hour)] = true
	}

	return len(uniqueDays)
}

func CalculateTotalDuration(activities []a.Activity) time.Duration {
	var totalDuration time.Duration

	for _, activity := range activities {
		totalDuration += activity.Time
	}

	return totalDuration
}

func CalculateTotalDistance(activities []a.Activity) float64 {
	var totalDistance float64

	for _, activity := range activities {
		totalDistance += activity.Distance
	}

	return totalDistance
}

func FindShortestRun(activities []a.Activity) a.Activity {
	var shortestRun a.Activity
	shortestDistance := math.MaxFloat64

	for _, activity := range activities {
		if activity.Distance < shortestDistance {
			shortestRun = activity
			shortestDistance = activity.Distance
		}
	}

	return shortestRun
}

func FindFarthestRun(activities []a.Activity) a.Activity {
	var farthestRun a.Activity
	farthestDistance := 0.0

	for _, activity := range activities {
		if activity.Distance > farthestDistance {
			farthestRun = activity
			farthestDistance = activity.Distance
		}
	}

	return farthestRun
}
