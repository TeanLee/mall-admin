package com.ting.demo.controller;

import com.ting.demo.bean.Admin;
import com.ting.demo.bean.RespBean;
import com.ting.demo.service.AdminService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

@Slf4j
@Controller
@RestController
public class LoginRegController {
    @Autowired
    AdminService adminService;

    @RequestMapping("/login_error")
    public RespBean loginError() {
        return new RespBean("error", "登录失败!");
    }

    @PostMapping("/login_success")
    public RespBean loginSuccess() {
        return new RespBean("success", "登录成功!");
    }

    /**
     * 如果自动跳转到这个页面，说明用户未登录，返回相应的提示即可
     * <p>
     * 如果要支持表单登录，可以在这个方法中判断请求的类型，进而决定返回JSON还是HTML页面
     *
     * @return
     */
    @RequestMapping("/login_page")
    public RespBean loginPage() {
        return new RespBean("error", "尚未登录，请登录!");
    }

//    @PostMapping("/login")
//    public RespBean loginPageError(@RequestParam("msg") String msg) {
//        log.info("loginPageError" + msg);
//        return new RespBean("error", "尚未登录，请登录!");
//    }
//
//    @GetMapping("/login")
//    public RespBean loginPageError1(@RequestParam("msg") String msg) {
//        log.info("get loginPageError" + msg);
//        return new RespBean("error", "尚未登录，请登录!");
//    }

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

    @RequestMapping("/")
    public String index(){
        return "login success";
    }

    @RequestMapping("/user")
    public String get(){
        return "user:xzw";
    }

    @RequestMapping("/admin")
    public String getAdmin(){
        return "admin:xzw";
    }
}
