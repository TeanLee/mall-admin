package com.ting.demo.service;

import com.ting.demo.bean.Admin;
import com.ting.demo.mapper.AdminMapper;
import com.ting.demo.mapper.RolesMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class AdminService {
    @Autowired
    AdminMapper adminMapper;
    @Autowired
    RolesMapper rolesMapper;

    public int updateAdmin(Integer adminId, String username, String password, Integer roleId) {
        int flag = 1, flag1 = 1, flag2 = 1;
        if (!username.equals("") || !password.equals("")) {
            flag1 = adminMapper.updateAdmin(adminId, username, password);
        }
        if (roleId != null) {
            flag2 = rolesMapper.updataRoleByAdminId(adminId, roleId);
        }
        if (flag1 == 1 && flag2 == 1) flag = 1;
        else flag = 0;
        // 修改账号密码 或者 权限 全部没有问题才返回正确，否则显示异常
        return flag;
    }
}
