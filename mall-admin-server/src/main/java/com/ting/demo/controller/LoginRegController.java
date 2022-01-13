package com.ting.demo.controller;

import com.ting.demo.bean.Admin;
import com.ting.demo.bean.RespBean;
import com.ting.demo.service.LoginRegService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

@Slf4j
@Controller
@RestController
public class LoginRegController {
    @Autowired
    LoginRegService loginRegService;
    @Autowired
    Admin admin;

    @PostMapping("/reg")
    public RespBean reg(Admin admin) {
        log.info("controller" + admin.getUsername());
        int result = loginRegService.reg(admin);
        if (result == 0) {
            //成功
            return new RespBean("200", "success", "注册成功!");
        } else if (result == 1) {
            return new RespBean("400","error", "用户名重复，注册失败!");
        } else {
            //失败
            return new RespBean("400","error", "注册失败!");
        }
    }
}
