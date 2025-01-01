package io.ysb.strava.booster.statistic

import java.time.Duration
import java.time.temporal.ChronoUnit

object Statistic {

    fun countUniqueDays(activities: List<Activity>): Int {
        val uniqueDays = activities.map { it.date.truncatedTo(ChronoUnit.DAYS) }.toSet()
        return uniqueDays.size
    }

    fun calculateTotalDuration(activities: List<Activity>): Duration {
        return activities.fold(Duration.ZERO) { total, activity -> total.plus(activity.time) }
    }

    fun calculateTotalDistance(activities: List<Activity>): Double {
        return activities.sumOf { it.distance }
    }

    fun findShortestRun(activities: List<Activity>): Activity? {
        return activities.minByOrNull { it.distance }
    }

    fun findLongestRun(activities: List<Activity>): Activity? {
        return activities.maxByOrNull { it.distance }
    }
}
