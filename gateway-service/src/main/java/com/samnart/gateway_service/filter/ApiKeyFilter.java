package com.samnart.gateway_service.filter;

import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpStatus;

@Configuration
public class ApiKeyFilter {
    
    private static final String API_KEY = "super-secret-key";

    @Bean
    public GlobalFilter apiKeyValidator() {
        return (exchange, chain) -> {
            String key = exchange.getRequest().getHeaders().getFirst("X-API-Key");
            if (!API_KEY.equals(key)) {
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }
            return chain.filter(exchange);
        };
    }
}
