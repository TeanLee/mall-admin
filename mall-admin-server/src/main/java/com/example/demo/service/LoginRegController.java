package com.example.demo.service;

import com.example.demo.bean.Admin;
import com.example.demo.mapper.AdminMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@Transactional
public class LoginRegController implements UserDetailsService {
    @Autowired
    Admin admin;
    @Autowired
    AdminMapper adminMapper;

    public Admin getAdminByUsername(String username) {
        return adminMapper.getUserByUsername(username);
    }

    @Override
    public Admin loadUserByUsername(String username) throws UsernameNotFoundException {
        Admin admin = adminMapper.getUserByUsername(username);
        if (admin == null) {
            //避免返回null，这里返回一个不含有任何值的User对象，在后期的密码比对过程中一样会验证失败
            return new Admin();
        }
        //查询用户的角色信息，并返回存入user中
//        List<Role> roles = rolesMapper.getRolesByUid(user.getId());
//        user.setRoles(roles);
        return admin;
    }

    /**
     * @param admin
     * @return 0表示成功
     * 1表示用户名重复
     * 2表示失败
     */
    public int reg(Admin admin) {
        Admin loadAdminByUsername = adminMapper.getUserByUsername(admin.getUsername());
        if (loadAdminByUsername != null) {
            return 1;
        }
        admin.setPassword(admin.getPassword());
        long result = adminMapper.reg(admin);
        if (result == 1) {
            return 0;
        } else {
            return 2;
        }
    }
}
