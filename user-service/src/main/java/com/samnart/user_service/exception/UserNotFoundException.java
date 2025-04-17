package com.samnart.user_service.exception;

public class UserNotFoundException extends RuntimeException {
    public UserNotFoundException(){}
    public UserNotFoundException(String message) {
        super(message);
    }
}
