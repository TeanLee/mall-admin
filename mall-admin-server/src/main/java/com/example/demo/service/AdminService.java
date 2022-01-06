package com.example.demo.service;

import com.example.demo.bean.Admin;
import com.example.demo.mapper.AdminMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AdminService {
    @Autowired
    AdminMapper adminMapper;

    public Admin getAdminByUsername(String username) {
        return adminMapper.getUserByUsername(username);
    }
}
