package com.example.demo.bean;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor //全参构造函数
@NoArgsConstructor  //无参构造函数
@Data
public class RespBean {
    private String status;
    private String msg;
}
