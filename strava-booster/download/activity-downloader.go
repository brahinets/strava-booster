package download

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Activities(from time.Time, auth string) []ActivityEntity {
	var activities []ActivityEntity

	var hasData = true
	var currentPage = 1

	for hasData == true {
		page := downloadPage(from, currentPage, auth)
		activities = append(activities, page.Activities...)

		totalPages := int(math.Ceil(float64(page.Total) / float64(page.PerPage)))
		hasData = currentPage < totalPages
		currentPage++
	}

	return activities
}

func downloadPage(from time.Time, page int, auth string) ActivitiesPage {
	err, req := buildRequest(from, page, auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var data ActivitiesPage
	json.NewDecoder(resp.Body).Decode(&data)

	return data
}

func buildRequest(from time.Time, page int, auth string) (error, *http.Request) {
	stravaUrl, err := buildUrl(from, page)

	req, err := http.NewRequest("GET", stravaUrl.String(), nil)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "_strava4_session="+auth)

	return err, req
}

func buildUrl(from time.Time, page int) (*url.URL, error) {
	stravaUrl, err := url.Parse("https://www.strava.com/athlete/training_activities")
	if err != nil {
		log.Fatal(err)
	}

	values := stravaUrl.Query()
	values.Add("start_date", from.Format("01/02/2006"))
	values.Set("activity_type", "Run")
	values.Set("per_page", strconv.Itoa(20))
	values.Set("page", strconv.Itoa(page))

	stravaUrl.RawQuery = values.Encode()

	return stravaUrl, err
}
