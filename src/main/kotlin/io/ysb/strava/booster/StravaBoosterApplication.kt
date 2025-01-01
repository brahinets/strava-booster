package io.ysb.strava.booster

import io.ysb.strava.booster.raw.ActivityDownloader
import io.ysb.strava.booster.raw.ActivityEntity
import io.ysb.strava.booster.statistic.Activity
import io.ysb.strava.booster.statistic.Statistic
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.ApplicationContext
import java.time.Duration
import java.time.LocalDate
import java.time.LocalDateTime
import java.time.ZoneOffset

@SpringBootApplication
class StravaBoosterApplication

fun main(args: Array<String>) {
    val context = runApplication<StravaBoosterApplication>(*args)
    val data = downloadRawData(context)
    val activities = mapActivities(data)
    runAnalytics(activities)
}

fun downloadRawData(context: ApplicationContext): List<ActivityEntity> {
    val analyticsStart = LocalDate.of(2024, 12, 1)
    val sessionCookie = System.getenv("STRAVA_SESSION_TOKEN")
    val downloader = context.getBean(ActivityDownloader::class.java)

    return downloader.downloadActivities(analyticsStart, sessionCookie)
}

fun runAnalytics(activities: List<Activity>) {
    val numDays = Statistic.countUniqueDays(activities)
    val numRuns = activities.size
    val totalDuration = Statistic.calculateTotalDuration(activities)
    val totalDistance = Statistic.calculateTotalDistance(activities)
    val shortestRun = Statistic.findShortestRun(activities)
    val longestRun = Statistic.findLongestRun(activities)

    println("----------------")
    println("# of Days: $numDays")
    println("# of Runs: $numRuns")
    println("Total Duration: $totalDuration")
    println("Total Distance: ${formatDistance(totalDistance)}")
    println("Shortest Run: ${formatDistance(shortestRun?.distance ?: 0.0)}")
    println("Longest Run: ${formatDistance(longestRun?.distance ?: 0.0)}")
}

fun formatDistance(distanceKilometers: Double): String {
    return "%.2fkm".format(distanceKilometers)
}

fun mapActivities(entities: List<ActivityEntity>): List<Activity> {
    return entities.map { entity ->
        Activity(
            sport = entity.type ?: "unknown",
            date = LocalDateTime.ofEpochSecond(entity.startDateLocalRaw, 0, ZoneOffset.UTC),
            title = entity.name,
            time = Duration.ofSeconds(entity.elapsedTimeRaw.toLong()),
            distance = entity.distanceRaw / 1000,
            elevation = entity.elevationGainRaw.toInt(),
            effort = entity.sufferScore.toInt()
        )
    }
}
