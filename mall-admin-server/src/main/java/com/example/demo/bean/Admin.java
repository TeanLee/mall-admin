package com.example.demo.bean;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.stereotype.Component;

@AllArgsConstructor //全参构造函数
@NoArgsConstructor  //无参构造函数
@Data
@Component
public class Admin {
    private int adminId;
    private String username;
    private String password;
}
