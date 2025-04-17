package com.samnart.user_service.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import org.springframework.stereotype.Service;

import com.samnart.user_service.exception.UserNotFoundException;
import com.samnart.user_service.model.User;
import com.samnart.user_service.repository.UserRepository;

@Service
public class UserService {
    
    private final UserRepository userRepository;
    
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public List<User> getAllUsers() {
        return userRepository.findAll();
    }

    public Optional<User> getUserById(String id) {
        return userRepository.findById(id);
    }

    public User addUser(String name, String email) {
        User user = new User(UUID.randomUUID().toString(), name, email, "ACTIVE");
        return userRepository.save(user);
    }
    
    public User updateUser(String id, String name, String email) {
        return userRepository.findById(id)
            .map(existingUser -> {
                existingUser.setName(name);
                existingUser.setEmail(email);
                return userRepository.save(existingUser);
            })
            .orElseThrow(() -> new UserNotFoundException("User not found with ID: " + id));
    }
    
    public void deleteUser(String id) {
        User user = userRepository.findById(id)
            .orElseThrow(() -> new UserNotFoundException("User not found with ID: " + id));
        
        userRepository.delete(user);
    }
}
