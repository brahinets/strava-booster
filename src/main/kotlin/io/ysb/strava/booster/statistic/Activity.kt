package io.ysb.strava.booster.statistic

import java.time.Duration
import java.time.LocalDateTime

data class Activity(
    val sport: String,
    val date: LocalDateTime,
    val title: String,
    val time: Duration,
    val distance: Double,
    val elevation: Int,
    val effort: Int
)
