package io.ysb.strava.booster.raw

import com.fasterxml.jackson.annotation.JsonIgnoreProperties
import com.fasterxml.jackson.annotation.JsonProperty

@JsonIgnoreProperties(ignoreUnknown = true)
data class ActivityEntity(
    @JsonProperty("id")
    val id: Long,

    @JsonProperty("name")
    val name: String = "",

    @JsonProperty("type")
    val type: String? = null,

    @JsonProperty("display_type")
    val displayType: String = "",

    @JsonProperty("activity_type_display_name")
    val activityTypeDisplayName: String = "",

    @JsonProperty("private")
    val myPrivate: Boolean = false,

    @JsonProperty("bike_id")
    val bikeId: Any? = null,

    @JsonProperty("athlete_gear_id")
    val athleteGearId: Int = 0,

    @JsonProperty("start_date")
    val startDate: String = "",

    @JsonProperty("start_date_local_raw")
    val startDateLocalRaw: Long = 0,

    @JsonProperty("start_time")
    val startTime: String = "",

    @JsonProperty("start_day")
    val startDay: String = "",

    @JsonProperty("distance")
    val distance: String = "",

    @JsonProperty("distance_raw")
    val distanceRaw: Double = 0.0,

    @JsonProperty("long_unit")
    val longUnit: String = "",

    @JsonProperty("short_unit")
    val shortUnit: String = "",

    @JsonProperty("moving_time")
    val movingTime: String = "",

    @JsonProperty("moving_time_raw")
    val movingTimeRaw: Int = 0,

    @JsonProperty("elapsed_time")
    val elapsedTime: String = "",

    @JsonProperty("elapsed_time_raw")
    val elapsedTimeRaw: Int = 0,

    @JsonProperty("trainer")
    val trainer: Boolean = false,

    @JsonProperty("static_map")
    val staticMap: String? = null,

    @JsonProperty("has_latlng")
    val hasLatlng: Boolean = false,

    @JsonProperty("commute")
    val commute: Any? = null,

    @JsonProperty("elevation_gain")
    val elevationGain: String = "",

    @JsonProperty("elevation_unit")
    val elevationUnit: String = "",

    @JsonProperty("elevation_gain_raw")
    val elevationGainRaw: Double = 0.0,

    @JsonProperty("description")
    val description: Any? = null,

    @JsonProperty("activity_url")
    val activityUrl: String = "",

    @JsonProperty("activity_url_for_twitter")
    val activityUrlForTwitter: String = "",

    @JsonProperty("twitter_msg")
    val twitterMsg: String = "",

    @JsonProperty("is_new")
    val isNew: Boolean = false,

    @JsonProperty("is_changing_type")
    val isChangingType: Boolean = false,

    @JsonProperty("suffer_score")
    val sufferScore: Double = 0.0,

    @JsonProperty("workout_type")
    val workoutType: Any? = null,

    @JsonProperty("flagged")
    val flagged: Boolean = false,

    @JsonProperty("hide_power")
    val hidePower: Boolean = false,

    @JsonProperty("hide_heartrate")
    val hideHeartrate: Boolean = false,

    @JsonProperty("leaderboard_opt_out")
    val leaderboardOptOut: Boolean = false,

    @JsonProperty("visibility")
    val visibility: String = ""
)
