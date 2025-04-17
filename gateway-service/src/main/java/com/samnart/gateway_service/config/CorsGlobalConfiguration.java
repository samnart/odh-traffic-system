package com.samnart.gateway_service.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsWebFilter;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;
import java.util.List;

@Configuration
public class CorsGlobalConfiguration {
    
    @Value("${cors.allowed-origins:*}")
    private List<String> allowedOrigins;
    
    @Value("${cors.allowed-methods:GET,POST,PUT,DELETE,OPTIONS}")
    private List<String> allowedMethods;
    
    @Value("${cors.allowed-headers:*}")
    private List<String> allowedHeaders;
    
    @Value("${cors.max-age:3600}")
    private long maxAge;

    @Bean
    public CorsWebFilter corsWebFilter() {
        CorsConfiguration corsConfig = new CorsConfiguration();
        corsConfig.setAllowedOrigins(allowedOrigins);
        corsConfig.setAllowedMethods(allowedMethods);
        corsConfig.setAllowedHeaders(allowedHeaders);
        corsConfig.setMaxAge(maxAge);
        
        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/api/**", corsConfig);
        
        return new CorsWebFilter(source);
    }
}