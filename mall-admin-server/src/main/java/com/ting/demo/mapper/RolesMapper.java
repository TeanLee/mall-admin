package com.ting.demo.mapper;

import com.ting.demo.bean.Role;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;

@Mapper
public interface RolesMapper {
    List<Role> getRolesByUid(Long uid);
    int addRoles(@Param("roles") String[] roles, @Param("adminId") Long uid);
    int updataRoleByAdminId(@Param("adminId") Integer adminId, @Param("roleId") Integer roleId);
}
