package com.samnart.gateway_service.filter;

import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.buffer.DataBuffer;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.server.reactive.ServerHttpResponse;
import reactor.core.publisher.Mono;

@Configuration
public class ErrorHandlingFilter {
    
    @Bean
    public GlobalFilter errorHanling() {
        return (exchange, chain) -> chain.filter(exchange)
            .onErrorResume(e -> {
                ServerHttpResponse response = exchange.getResponse();
                response.setStatusCode(HttpStatus.SERVICE_UNAVAILABLE);
                response.getHeaders().setContentType(MediaType.APPLICATION_JSON);
                String error = "{\"error\": \"" + e.getMessage() + "\"}";
                DataBuffer buffer = response.bufferFactory().wrap(error.getBytes());
                return response.writeWith(Mono.just(buffer));
            });
    }
}
