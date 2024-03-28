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
	data := downloadRawData()
	activities := mapActivities(data)
	runAnalytics(activities)
}

func downloadRawData() []download.ActivityEntity {
	analyticsStart := time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC)
	sessionCookie := os.Getenv("STRAVA_SESSION_TOKEN")
	return download.Activities(analyticsStart, sessionCookie)
}

func runAnalytics(activities []activity.Activity) {
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

func mapActivities(entities []download.ActivityEntity) []activity.Activity {
	var activities []activity.Activity

	for _, entity := range entities {
		activities = append(activities, activity.Activity{
			Sport:     entity.Type,
			Date:      time.Unix(entity.StartDateLocalRaw, 0),
			Title:     entity.Name,
			Time:      time.Duration(entity.ElapsedTimeRaw) * time.Second,
			Distance:  float64(entity.DistanceRaw) / 1000,
			Elevation: int(entity.ElevationGainRaw),
			Effort:    int(entity.SufferScore),
		})
	}

	return activities
}
