package com.samnart.gateway_service.filter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.server.reactive.ServerHttpRequest;

@Configuration
public class LoggingFilter {
    
    private static final Logger logger = LoggerFactory.getLogger(LoggingFilter.class);

    @Bean
    public GlobalFilter logRequestResponse() {
        return (exchange, chain) -> {
            ServerHttpRequest request = exchange.getRequest();
            logger.info("{} {}", request.getMethod(), request.getURI());
            return chain.filter(exchange).doOnSuccess(done ->
                logger.info("Response sent for {}", request.getURI())
            );
        };
    }
}
