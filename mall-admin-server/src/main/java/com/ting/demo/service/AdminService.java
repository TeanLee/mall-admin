package com.ting.demo.service;

import com.ting.demo.bean.Admin;
import com.ting.demo.mapper.AdminMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class AdminService {
    @Autowired
    AdminMapper adminMapper;

    private int updateAdmin(String adminId, String username, String password) {
        log.info("AdminService adminId：" + adminId + "username：" + username);
        int flag = adminMapper.updateAdmin(adminId, username, password);
        log.info("flag：" + flag);
        return flag;
    }
}
