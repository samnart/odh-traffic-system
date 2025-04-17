package com.samnart.gateway_service.filter;

import io.github.bucket4j.Bandwidth;
import io.github.bucket4j.Bucket;
import io.github.bucket4j.Bucket4j;
import io.github.bucket4j.Refill;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Component
@Order(Ordered.HIGHEST_PRECEDENCE + 2)
public class RateLimiterFilter implements GlobalFilter {

    private final Map<String, Bucket> buckets = new ConcurrentHashMap<>();
    
    @Value("${rate-limiter.capacity:100}")
    private int capacity;
    
    @Value("${rate-limiter.refill-tokens:10}")
    private int refillTokens;
    
    @Value("${rate-limiter.refill-duration:1}")
    private int refillDuration;

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        String clientIpAddress = exchange.getRequest().getRemoteAddress().getAddress().getHostAddress();
        
        // Get or create a bucket for this client
        Bucket bucket = buckets.computeIfAbsent(clientIpAddress, this::createNewBucket);
        
        // Try to consume a token from the bucket
        if (bucket.tryConsume(1)) {
            return chain.filter(exchange);
        } else {
            // If no tokens available, return 429 Too Many Requests
            exchange.getResponse().setStatusCode(HttpStatus.TOO_MANY_REQUESTS);
            exchange.getResponse().getHeaders().add("X-Rate-Limit-Retry-After-Seconds", 
                String.valueOf(refillDuration));
            return exchange.getResponse().setComplete();
        }
    }
    
    private Bucket createNewBucket(String key) {
        Refill refill = Refill.intervally(refillTokens, Duration.ofSeconds(refillDuration));
        Bandwidth limit = Bandwidth.classic(capacity, refill);
        return Bucket4j.builder().addLimit(limit).build();
    }
}