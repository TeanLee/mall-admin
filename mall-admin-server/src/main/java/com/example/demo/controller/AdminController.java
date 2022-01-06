package com.example.demo.controller;

import com.example.demo.bean.Admin;
import com.example.demo.bean.RespBean;
import com.example.demo.service.LoginRegController;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RestController
@RequestMapping("/admin")
public class AdminController {
    @Autowired
    LoginRegController loginRegController;

    @GetMapping("/get")
    public Admin getAdminByUsername(@RequestParam("username") String username) {
        Admin admin = loginRegController.getAdminByUsername(username);
        log.info(String.valueOf(admin));
        return admin;
    }

    @PostMapping("/register")
    public RespBean register(Admin admin) {
        int result = loginRegController.reg(admin);
        if (result == 0) {
            //成功
            return new RespBean("success", "注册成功!");
        } else if (result == 1) {
            return new RespBean("error", "用户名重复，注册失败!");
        } else {
            //失败
            return new RespBean("error", "注册失败!");
        }
    }
}
