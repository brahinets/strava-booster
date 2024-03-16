package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	activities := ReadActivities("activities.csv")

	for _, activity := range activities {
		fmt.Printf("Sport: %s, Date: %s, Title: %s, Time: %s, Distance: %.2f km, Elevation: %d m, Effort: %d\n",
			activity.Sport, activity.Date, activity.Title, activity.Time, activity.Distance, activity.Elevation, activity.Effort)
	}
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
