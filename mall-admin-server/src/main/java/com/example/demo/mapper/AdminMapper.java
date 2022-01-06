package com.example.demo.mapper;

import com.example.demo.bean.Admin;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface AdminMapper {
    public Admin getUserByUsername(String username);
    long reg(Admin admin);
}
