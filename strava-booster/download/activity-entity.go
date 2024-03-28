package download

type ActivityEntity struct {
	ID                      int64   `json:"id"`
	Name                    string  `json:"name"`
	Type                    string  `json:"type"`
	DisplayType             string  `json:"display_type"`
	ActivityTypeDisplayName string  `json:"activity_type_display_name"`
	Private                 bool    `json:"private"`
	BikeID                  any     `json:"bike_id"`
	AthleteGearID           int     `json:"athlete_gear_id"`
	StartDate               string  `json:"start_date"`
	StartDateLocalRaw       int     `json:"start_date_local_raw"`
	StartTime               string  `json:"start_time"`
	StartDay                string  `json:"start_day"`
	Distance                string  `json:"distance"`
	DistanceRaw             float64 `json:"distance_raw"`
	LongUnit                string  `json:"long_unit"`
	ShortUnit               string  `json:"short_unit"`
	MovingTime              string  `json:"moving_time"`
	MovingTimeRaw           int     `json:"moving_time_raw"`
	ElapsedTime             string  `json:"elapsed_time"`
	ElapsedTimeRaw          int     `json:"elapsed_time_raw"`
	Trainer                 bool    `json:"trainer"`
	StaticMap               string  `json:"static_map"`
	HasLatlng               bool    `json:"has_latlng"`
	Commute                 any     `json:"commute"`
	ElevationGain           string  `json:"elevation_gain"`
	ElevationUnit           string  `json:"elevation_unit"`
	ElevationGainRaw        float64 `json:"elevation_gain_raw"`
	Description             any     `json:"description"`
	ActivityURL             string  `json:"activity_url"`
	ActivityURLForTwitter   string  `json:"activity_url_for_twitter"`
	TwitterMsg              string  `json:"twitter_msg"`
	IsNew                   bool    `json:"is_new"`
	IsChangingType          bool    `json:"is_changing_type"`
	SufferScore             float64 `json:"suffer_score"`
	WorkoutType             any     `json:"workout_type"`
	Flagged                 bool    `json:"flagged"`
	HidePower               bool    `json:"hide_power"`
	HideHeartrate           bool    `json:"hide_heartrate"`
	LeaderboardOptOut       bool    `json:"leaderboard_opt_out"`
	Visibility              string  `json:"visibility"`
}
