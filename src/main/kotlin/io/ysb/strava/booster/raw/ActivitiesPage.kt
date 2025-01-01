package io.ysb.strava.booster.raw

import com.fasterxml.jackson.annotation.JsonProperty

data class ActivitiesPage(
    @JsonProperty("models")
    val activities: List<ActivityEntity>,

    @JsonProperty("page")
    val page: Int,

    @JsonProperty("perPage")
    val perPage: Int,

    @JsonProperty("total")
    val total: Int
)
