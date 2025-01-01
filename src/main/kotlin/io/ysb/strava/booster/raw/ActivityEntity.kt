package io.ysb.strava.booster.raw

import com.fasterxml.jackson.annotation.JsonProperty

data class ActivityEntity(
    @JsonProperty("id")
    val id: Long,

    @JsonProperty("name")
    val name: String,

    @JsonProperty("type")
    val type: String?,

    @JsonProperty("display_type")
    val displayType: String,

    @JsonProperty("activity_type_display_name")
    val activityTypeDisplayName: String,

    @JsonProperty("private")
    val myPrivate: Boolean,

    @JsonProperty("bike_id")
    val bikeId: Any?,

    @JsonProperty("athlete_gear_id")
    val athleteGearId: Int,

    @JsonProperty("start_date")
    val startDate: String,

    @JsonProperty("start_date_local_raw")
    val startDateLocalRaw: Long,

    @JsonProperty("start_time")
    val startTime: String,

    @JsonProperty("start_day")
    val startDay: String,

    @JsonProperty("distance")
    val distance: String,

    @JsonProperty("distance_raw")
    val distanceRaw: Double,

    @JsonProperty("long_unit")
    val longUnit: String,

    @JsonProperty("short_unit")
    val shortUnit: String,

    @JsonProperty("moving_time")
    val movingTime: String,

    @JsonProperty("moving_time_raw")
    val movingTimeRaw: Int,

    @JsonProperty("elapsed_time")
    val elapsedTime: String,

    @JsonProperty("elapsed_time_raw")
    val elapsedTimeRaw: Int,

    @JsonProperty("trainer")
    val trainer: Boolean,

    @JsonProperty("static_map")
    val staticMap: String?,

    @JsonProperty("has_latlng")
    val hasLatlng: Boolean,

    @JsonProperty("commute")
    val commute: Any?,

    @JsonProperty("elevation_gain")
    val elevationGain: String,

    @JsonProperty("elevation_unit")
    val elevationUnit: String,

    @JsonProperty("elevation_gain_raw")
    val elevationGainRaw: Double,

    @JsonProperty("description")
    val description: Any?,

    @JsonProperty("activity_url")
    val activityUrl: String,

    @JsonProperty("activity_url_for_twitter")
    val activityUrlForTwitter: String,

    @JsonProperty("twitter_msg")
    val twitterMsg: String,

    @JsonProperty("is_new")
    val isNew: Boolean,

    @JsonProperty("is_changing_type")
    val isChangingType: Boolean,

    @JsonProperty("suffer_score")
    val sufferScore: Double,

    @JsonProperty("workout_type")
    val workoutType: Any?,

    @JsonProperty("flagged")
    val flagged: Boolean,

    @JsonProperty("hide_power")
    val hidePower: Boolean,

    @JsonProperty("hide_heartrate")
    val hideHeartrate: Boolean,

    @JsonProperty("leaderboard_opt_out")
    val leaderboardOptOut: Boolean,

    @JsonProperty("visibility")
    val visibility: String
)
