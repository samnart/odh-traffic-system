package com.samnart.gateway_service.filter;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.http.HttpStatus;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.util.List;

@Component
@Order(Ordered.HIGHEST_PRECEDENCE)
public class ApiKeyFilter implements GlobalFilter {
    
    @Value("${security.api-key}")
    private String apiKey;
    
    @Value("${security.api-key-header:X-API-Key}")
    private String apiKeyHeader;
    
    @Value("#{'${security.public-paths:/actuator/health,/actuator/info}'.split(',')}")
    private List<String> publicPaths;

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        ServerHttpRequest request = exchange.getRequest();
        String path = request.getURI().getPath();
        
        // Skip API key validation for public paths
        if (publicPaths.stream().anyMatch(path::startsWith)) {
            return chain.filter(exchange);
        }
        
        // Validate API key
        String providedKey = request.getHeaders().getFirst(apiKeyHeader);
        if (providedKey == null || !apiKey.equals(providedKey)) {
            exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
            return exchange.getResponse().setComplete();
        }
        
        return chain.filter(exchange);
    }
}