package download

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Activities(from time.Time, auth string) []ActivityEntity {
	err, req := buildRequest(from, auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var data ActivitiesPage
	json.NewDecoder(resp.Body).Decode(&data)

	return data.Activities
}

func buildRequest(from time.Time, auth string) (error, *http.Request) {
	stravaUrl, err := buildUrl(from)

	req, err := http.NewRequest("GET", stravaUrl.String(), nil)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "_strava4_session="+auth)

	return err, req
}

func buildUrl(from time.Time) (*url.URL, error) {
	stravaUrl, err := url.Parse("https://www.strava.com/athlete/training_activities")
	if err != nil {
		log.Fatal(err)
	}

	values := stravaUrl.Query()
	values.Add("start_date", from.Format("01/02/2006"))
	values.Set("activity_type", "Run")
	values.Set("per_page", strconv.Itoa(20))

	stravaUrl.RawQuery = values.Encode()

	return stravaUrl, err
}
