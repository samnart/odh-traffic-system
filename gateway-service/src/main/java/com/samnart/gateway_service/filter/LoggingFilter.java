package com.samnart.gateway_service.filter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.UUID;

@Component
@Order(Ordered.HIGHEST_PRECEDENCE + 1)
public class LoggingFilter implements GlobalFilter {
    
    private static final Logger logger = LoggerFactory.getLogger(LoggingFilter.class);

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        ServerHttpRequest request = exchange.getRequest();
        String requestId = UUID.randomUUID().toString();
        
        // Add request ID to the exchange attributes for tracking
        exchange.getAttributes().put("requestId", requestId);
        
        long startTime = System.currentTimeMillis();
        
        // Log request details
        logger.info(
            "Request: [{}] {} {} from {} with headers {}",
            requestId,
            request.getMethod(),
            request.getURI(),
            request.getRemoteAddress(),
            request.getHeaders()
        );
        
        // Add the request ID as a header for downstream services
        ServerHttpRequest modifiedRequest = request.mutate()
            .header("X-Request-ID", requestId)
            .build();
        
        // Replace the request with the modified one
        ServerWebExchange modifiedExchange = exchange.mutate()
            .request(modifiedRequest)
            .build();
        
        // Process and log the response
        return chain.filter(modifiedExchange)
            .doOnSuccess(done -> {
                long duration = System.currentTimeMillis() - startTime;
                logger.info(
                    "Response: [{}] status {} in {} ms",
                    requestId,
                    exchange.getResponse().getStatusCode(),
                    duration
                );
            })
            .doOnError(error -> {
                long duration = System.currentTimeMillis() - startTime;
                logger.error(
                    "Error: [{}] {} in {} ms",
                    requestId,
                    error.getMessage(),
                    duration,
                    error
                );
            });
    }
}