package com.ting.demo.controller;

import com.ting.demo.bean.Admin;
import com.ting.demo.bean.RespBean;
import com.ting.demo.service.AdminService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.Collection;

@Slf4j
@Controller
@RestController
public class LoginRegController {
    @Autowired
    AdminService adminService;
    @Autowired
    Admin admin;

    @GetMapping("/")
    public void getRole() {
        UserDetails userDetails = (UserDetails) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        Collection<? extends GrantedAuthority> auths = userDetails.getAuthorities();
        log.info("GrantedAuthority" + String.valueOf(auths));
    }

    @PostMapping("/reg")
    public RespBean reg(Admin admin) {
        log.info("controller" + admin.getUsername());
        int result = adminService.reg(admin);
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
