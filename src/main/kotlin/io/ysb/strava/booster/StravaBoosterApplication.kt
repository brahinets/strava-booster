package io.ysb.strava.booster

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class StravaBoosterApplication

fun main(args: Array<String>) {
	runApplication<StravaBoosterApplication>(*args)
}
