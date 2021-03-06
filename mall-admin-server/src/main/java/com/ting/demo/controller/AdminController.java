package com.ting.demo.controller;

import com.ting.demo.bean.Admin;
import com.ting.demo.bean.RespBean;
import com.ting.demo.mapper.AdminMapper;
import com.ting.demo.service.AdminService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@Slf4j
@Controller
@RestController
@RequestMapping("/admin")
public class AdminController {
    @Autowired
    Admin admin;
    @Autowired
    AdminService adminService;

    @GetMapping("/info")
    public RespBean getRole() {
        Map<String , Object> userMap = new HashMap<>();
        userMap.put("userid", admin.getAdminId());
        userMap.put("username", admin.getUsername());
        userMap.put("roles", admin.getRoles());
        return new RespBean("200", "success", userMap);
    }

    // 修改用户信息
    @PostMapping("/{id}")
    public RespBean updateAdmin(@PathVariable("id") Integer adminId, @RequestParam(required = false, value = "username", defaultValue = "") String username, @RequestParam(required = false, value = "password", defaultValue = "") String password, @RequestParam(required = false, value = "roleId", defaultValue = "") Integer roleId) {
        int updateRes = adminService.updateAdmin(adminId, username, password, roleId);
        if (updateRes == 1) {
            return new RespBean("200", "success");
        } else {
            return new RespBean("400", "failed");
        }
    }
}
