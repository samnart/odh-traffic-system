package com.samnart.gateway_service.filter;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.annotation.Order;
import org.springframework.core.io.buffer.DataBuffer;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.server.reactive.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;

@Component
@Order(-1) // Ensure this runs after other filters but before response is sent
public class ErrorHandlingFilter implements GlobalFilter {
    
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        return chain.filter(exchange)
            .onErrorResume(e -> {
                ServerHttpResponse response = exchange.getResponse();
                HttpStatus status = determineHttpStatus(e);
                response.setStatusCode(status);
                response.getHeaders().setContentType(MediaType.APPLICATION_JSON);
                
                Map<String, Object> errorResponse = new HashMap<>();
                errorResponse.put("timestamp", LocalDateTime.now().toString());
                errorResponse.put("status", status.value());
                errorResponse.put("error", status.getReasonPhrase());
                errorResponse.put("message", e.getMessage());
                errorResponse.put("path", exchange.getRequest().getURI().getPath());
                
                try {
                    String errorJson = objectMapper.writeValueAsString(errorResponse);
                    DataBuffer buffer = response.bufferFactory().wrap(errorJson.getBytes());
                    return response.writeWith(Mono.just(buffer));
                } catch (JsonProcessingException jsonException) {
                    // Fallback to simple error message if JSON processing fails
                    String error = "{\"error\": \"" + e.getMessage() + "\"}";
                    DataBuffer buffer = response.bufferFactory().wrap(error.getBytes());
                    return response.writeWith(Mono.just(buffer));
                }
            });
    }
    
    private HttpStatus determineHttpStatus(Throwable error) {
        // Map different exceptions to appropriate HTTP status codes
        if (error instanceof java.net.ConnectException) {
            return HttpStatus.SERVICE_UNAVAILABLE;
        } else if (error instanceof java.util.concurrent.TimeoutException) {
            return HttpStatus.GATEWAY_TIMEOUT;
        } else if (error instanceof org.springframework.web.server.ResponseStatusException) {
            return ((org.springframework.web.server.ResponseStatusException) error).getStatusCode().is4xxClientError() 
                ? HttpStatus.BAD_REQUEST 
                : HttpStatus.INTERNAL_SERVER_ERROR;
        }
        return HttpStatus.INTERNAL_SERVER_ERROR;
    }
}