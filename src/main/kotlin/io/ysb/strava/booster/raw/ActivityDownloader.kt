package io.ysb.strava.booster.raw

import tools.jackson.databind.ObjectMapper
import org.springframework.http.HttpEntity
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpMethod
import org.springframework.stereotype.Service
import org.springframework.web.client.RestTemplate
import java.time.Instant
import java.time.LocalDate
import java.time.ZoneOffset
import java.time.format.DateTimeFormatter
import java.util.Locale

@Service
class ActivityDownloader(
    private val restTemplate: RestTemplate,
    private val mapper: ObjectMapper
) {
    fun downloadActivities(from: LocalDate, auth: String): List<ActivityEntity> {
        val activities = mutableListOf<ActivityEntity>()
        var hasData = true
        var currentPage = 1

        println("Downloading activities since ${from.format(DateTimeFormatter.ofPattern("yyyy-MM-dd"))}...")
        while (hasData) {
            val page = downloadPage(from, currentPage, auth)

            val totalPages = (page.total / page.perPage.toDouble()).toInt()
            println("Pages downloaded $currentPage of $totalPages")

            activities.addAll(page.activities)
            page.activities.forEach { logActivity(it) }

            hasData = currentPage < totalPages && activities.none { notBefore(it, from) }
            currentPage++
        }

        return activities.filter { !notBefore(it, from) }
    }

    private fun logActivity(it: ActivityEntity) {
        val zdt = Instant.ofEpochSecond(it.startDateLocalRaw).atZone(ZoneOffset.UTC)
        val date = zdt.format(DateTimeFormatter.ofPattern("yyyy-MM-dd"))
        val day = zdt.format(DateTimeFormatter.ofPattern("EEEE", Locale.ENGLISH))
        println(listOf(
            it.activityTypeDisplayName,
            it.name,
            day,
            date,
            zdt.format(DateTimeFormatter.ofPattern("HH:mm")),
            it.distance,
            it.elapsedTimeRaw,
            it.movingTimeRaw,
            it.elevationGain,
            it.activityUrl
        ).joinToString("\t"))
    }

    private fun notBefore(it: ActivityEntity, from: LocalDate) =
        it.startDateLocalRaw < from.atStartOfDay().toEpochSecond(ZoneOffset.UTC)

    private fun downloadPage(from: LocalDate, page: Int, auth: String): ActivitiesPage {
        val url = buildUrl(from, page)
        val headers = HttpHeaders().apply {
            set("X-Requested-With", "XMLHttpRequest")
            set("Cookie", "_strava4_session=$auth")
        }

        val entity = HttpEntity<String>(headers)
        val response = restTemplate.exchange(url, HttpMethod.GET, entity, String::class.java)

        return mapper.readValue(response.body, ActivitiesPage::class.java)
    }

    private fun buildUrl(from: LocalDate, page: Int): String {
        val stravaUrl = "https://www.strava.com/athlete/training_activities"

        return "$stravaUrl?start_date=${from.format(DateTimeFormatter.ofPattern("MM/dd/yyyy"))}&activity_type=Run&per_page=20&page=$page"
    }
}
