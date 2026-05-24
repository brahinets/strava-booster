package io.ysb.strava.booster.config

import org.springframework.boot.jackson.autoconfigure.JsonMapperBuilderCustomizer
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import tools.jackson.databind.DeserializationFeature
import tools.jackson.databind.SerializationFeature
import tools.jackson.databind.ext.javatime.ser.LocalDateSerializer
import tools.jackson.databind.module.SimpleModule
import tools.jackson.databind.ser.std.ToStringSerializer
import java.time.Duration
import java.time.LocalDate
import java.time.format.DateTimeFormatter

private val LOCAL_DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd")

@Configuration
class SerializationConfig {

    @Bean
    fun jsonMapperCustomizer() = JsonMapperBuilderCustomizer { builder ->
        builder
            .addModule(
                SimpleModule("custom-serializers")
                    .addSerializer(LocalDate::class.java, LocalDateSerializer(LOCAL_DATE_FORMATTER))
                    .addSerializer(Duration::class.java, ToStringSerializer(Duration::class.java))
            )
            .disable(SerializationFeature.FAIL_ON_EMPTY_BEANS)
            .disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES)
            .disable(DeserializationFeature.FAIL_ON_NULL_FOR_PRIMITIVES)
    }
}
