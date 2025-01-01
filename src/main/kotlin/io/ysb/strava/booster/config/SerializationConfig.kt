package io.ysb.strava.booster.config

import com.fasterxml.jackson.databind.ser.std.ToStringSerializer
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer
import com.fasterxml.jackson.module.kotlin.KotlinFeature
import com.fasterxml.jackson.module.kotlin.KotlinModule
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder
import java.time.Duration
import java.time.LocalDate
import java.time.format.DateTimeFormatter

private val LOCAL_DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd")

@Configuration
class SerializationConfig {

    @Bean
    fun objectMapper(): Jackson2ObjectMapperBuilder {
        return Jackson2ObjectMapperBuilder()
            .failOnUnknownProperties(false)
            .modules(
                KotlinModule.Builder()
                    .configure(KotlinFeature.NullToEmptyCollection, true)
                    .configure(KotlinFeature.NullToEmptyMap, true)
                    .build()
            )
            .modules(
                JavaTimeModule()
                    .addSerializer(LocalDate::class.java, LocalDateSerializer(LOCAL_DATE_FORMATTER))
                    .addSerializer(Duration::class.java, ToStringSerializer())
            )
    }
}
