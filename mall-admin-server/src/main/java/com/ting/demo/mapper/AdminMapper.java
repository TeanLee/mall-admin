package com.ting.demo.mapper;

import com.ting.demo.bean.Admin;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

@Mapper
public interface AdminMapper {
    Admin loadUserByUsername(@Param("username") String username);
    long reg(Admin admin);
    int updateAdmin(@Param("adminId") String adminId, @Param("username") String username, @Param("password") String password);
}
