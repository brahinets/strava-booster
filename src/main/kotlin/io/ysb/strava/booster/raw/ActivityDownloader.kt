package io.ysb.strava.booster.raw

import com.fasterxml.jackson.databind.ObjectMapper
import org.springframework.http.HttpEntity
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpMethod
import org.springframework.stereotype.Service
import org.springframework.web.client.RestTemplate
import java.time.LocalDate
import java.time.format.DateTimeFormatter

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

            hasData = currentPage < totalPages
            currentPage++
        }

        return activities
    }

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
        // TODO API does not filter by date anymore. Fetch only needed data

        return "$stravaUrl?start_date=${from.format(DateTimeFormatter.ofPattern("MM/dd/yyyy"))}&activity_type=Run&per_page=20&page=$page"
    }
}
