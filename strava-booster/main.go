package main

import (
	"fmt"
	"os"
	"strava-booster/activity"
	"strava-booster/download"
	"strava-booster/statistic"
	"time"
)

func main() {
	downloadRawData()
	runAnalytics()
}

func downloadRawData() {
	weekAgo := time.Now().Add(-time.Duration(24*7) * time.Hour)
	sessionCookie := os.Getenv("strava_session_token")
	activities := download.Activities(weekAgo, sessionCookie)

	fmt.Println("Raw Data:", activities)
}

func runAnalytics() {
	activities := activity.ReadActivities("../activities.csv")
	numDays := statistic.CountUniqueDays(activities)
	numRuns := len(activities)
	totalDuration := statistic.CalculateTotalDuration(activities)
	totalDistance := statistic.CalculateTotalDistance(activities)
	shortestRun := statistic.FindShortestRun(activities)
	farthestRun := statistic.FindFarthestRun(activities)

	fmt.Println("# of Days:", numDays)
	fmt.Println("# of Runs:", numRuns)
	fmt.Println("Total Duration:", totalDuration)
	fmt.Println("Total Distance:", formatDistance(totalDistance))
	fmt.Println("Shortest Run:", formatDistance(shortestRun.Distance))
	fmt.Println("Farthest Run:", formatDistance(farthestRun.Distance))
}

func formatDistance(distanceKilometers float64) string {
	return fmt.Sprintf("%.2fkm", distanceKilometers)
}
