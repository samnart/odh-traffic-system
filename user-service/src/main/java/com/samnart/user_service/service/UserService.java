package com.samnart.user_service.service;

import java.util.List;
import java.util.UUID;
import java.util.concurrent.CopyOnWriteArrayList;

import org.springframework.stereotype.Service;

import com.samnart.user_service.model.User;

@Service
public class UserService {
    
    private final List<User> users = new CopyOnWriteArrayList<>();

    public List<User> getAllUsers() {
        return users;
    }

    public User addUser(String name, String email) {
        User user = new User(UUID.randomUUID().toString(), name, email);
        users.add(user);
        return user;
    }
}
