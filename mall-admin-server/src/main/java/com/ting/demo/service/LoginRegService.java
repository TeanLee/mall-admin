package com.ting.demo.service;

import com.ting.demo.bean.Admin;
import com.ting.demo.bean.Role;
import com.ting.demo.mapper.AdminMapper;
import com.ting.demo.mapper.RolesMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.List;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

@Slf4j
@Service
public class LoginRegService implements UserDetailsService {
    @Autowired
    AdminMapper adminMapper;
    @Autowired
    RolesMapper rolesMapper;
    @Autowired
    Admin admin;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        Admin user = adminMapper.loadUserByUsername(username);
        log.info("loadUserByUsername" + username);
        if (user == null) {
            //避免返回null，这里返回一个不含有任何值的User对象，在后期的密码比对过程中一样会验证失败
            return new Admin();
        }
//        String password = user.getPassword();
//        //查询用户的角色信息，并返回存入user中
//        List<Role> roles = rolesMapper.getRolesByUid(user.getId());
//        user.setRoles(roles);
//        return user;
//        Admin user = adminMapper.loadUserByUsername(username);
//        if (user == null) {
//            return new Admin();
//        }
//        log.info("loadUserByUsername" + username);
//        if (username.equalsIgnoreCase("1")) {
//            throw new BadCredentialsException("帐号不存在，请重新输入");
//        }
        UserDetails userDetails =null;
//        List<GrantedAuthority> authList = new ArrayList<GrantedAuthority>();
        List<Role> roles = rolesMapper.getRolesByUid(user.getAdminId());
        log.info("List<Role> roles：" + user.getAdminId());
        log.info("List<Role> roles：" + roles);
        user.setRoles(roles);
        admin.setRoles(roles);
        String password = user.getPassword();
        userDetails = new Admin(username,password,roles);
        return userDetails;
    }

    /**
     * @param admin
     * @return 0表示成功
     * 1表示用户名重复
     * 2表示失败
     */
    public int reg(Admin admin) {
        Admin loadUserByUsername = adminMapper.loadUserByUsername(admin.getUsername());
        if (loadUserByUsername != null) {
            return 1;
        }
        BCryptPasswordEncoder encoder = new BCryptPasswordEncoder();
        //插入用户,插入之前先对密码进行加密
//        admin.setPassword(encoder.encode(admin.getPassword()));
        admin.setPassword(admin.getPassword());
        long result = adminMapper.reg(admin);
        //配置用户的角色，默认都是普通用户
        String[] roles = new String[]{"-1"};
        int i = rolesMapper.addRoles(roles, admin.getAdminId());
        boolean b = i == roles.length && result == 1;
        if (b) {
            return 0;
        } else {
            return 2;
        }
    }
}
