package com.example.demo.controller;

import com.example.demo.bean.Admin;
import com.example.demo.service.AdminService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RestController
@RequestMapping("/admin")
public class AdminController {
    @Autowired
    AdminService adminService;

    @GetMapping("/get")
    public Admin getAdminByUsername(@RequestParam("username") String username) {
        Admin admin = adminService.getAdminByUsername(username);
        log.info(String.valueOf(admin));
        return admin;
    }
}
