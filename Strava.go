package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	activities := ReadActivities("activities.csv")

	numDays := countUniqueDays(activities)
	numRuns := len(activities)
	totalDuration := calculateTotalDuration(activities)
	totalDistance := calculateTotalDistance(activities)
	shortestRun := findShortestRun(activities)
	farthestRun := findFarthestRun(activities)

	fmt.Println("# of Days:", numDays)
	fmt.Println("# of Runs:", numRuns)
	fmt.Println("Total Duration:", totalDuration)
	fmt.Println("Total Distance:", formatDistance(totalDistance))
	fmt.Println("Shortest Run:", formatDistance(shortestRun.Distance))
	fmt.Println("Farthest Run:", formatDistance(farthestRun.Distance))
}

func countUniqueDays(activities []Activity) int {
	uniqueDays := make(map[time.Time]bool)

	for _, activity := range activities {
		uniqueDays[activity.Date.Truncate(24*time.Hour)] = true
	}

	return len(uniqueDays)
}

func formatDistance(distanceMeters float64) string {
	return fmt.Sprintf("%.2fkm", distanceMeters)
}

func calculateTotalDuration(activities []Activity) time.Duration {
	var totalDuration time.Duration

	for _, activity := range activities {
		totalDuration += activity.Time
	}

	return totalDuration
}

func calculateTotalDistance(activities []Activity) float64 {
	var totalDistance float64

	for _, activity := range activities {
		totalDistance += activity.Distance
	}

	return totalDistance
}

func findShortestRun(activities []Activity) Activity {
	var shortestRun Activity
	shortestDistance := math.MaxFloat64

	for _, activity := range activities {
		if activity.Distance < shortestDistance {
			shortestRun = activity
			shortestDistance = activity.Distance
		}
	}

	return shortestRun
}

func findFarthestRun(activities []Activity) Activity {
	var farthestRun Activity
	farthestDistance := 0.0

	for _, activity := range activities {
		if activity.Distance > farthestDistance {
			farthestRun = activity
			farthestDistance = activity.Distance
		}
	}

	return farthestRun
}

type Activity struct {
	Sport     string
	Date      time.Time
	Title     string
	Time      time.Duration
	Distance  float64
	Elevation int
	Effort    int
}

func ReadActivities(activitiesFile string) []Activity {
	file, err := os.Open(activitiesFile)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var activities []Activity
	for _, record := range records {
		duration, err := parseDuration(record[3])
		if err != nil {
			fmt.Println("Error parsing duration:", err)
			return nil
		}

		distanceStr := strings.TrimSuffix(strings.TrimSpace(record[4]), " km")
		distance, err := strconv.ParseFloat(distanceStr, 64)
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			return nil
		}

		elevationStr := strings.TrimSuffix(strings.TrimSpace(record[5]), " m")
		elevation, err := strconv.Atoi(elevationStr)
		if err != nil {
			fmt.Println("Error parsing elevation:", err)
			return nil
		}

		effort, err := strconv.Atoi(record[6])
		if err != nil {
			fmt.Println("Error parsing effort:", err)
			return nil
		}

		date, err := time.Parse("1/2/2006", record[1])
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil
		}

		a := Activity{
			Sport:     record[0],
			Date:      date,
			Title:     record[2],
			Time:      duration,
			Distance:  distance,
			Elevation: elevation,
			Effort:    effort,
		}
		activities = append(activities, a)
	}

	return activities
}

func parseDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, ":")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid duration format")
	}

	seconds, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		return 0, err
	}

	hours := 0
	if len(parts) == 3 {
		hours, err = strconv.Atoi(parts[len(parts)-3])
		if err != nil {
			return 0, err
		}
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}
