package com.ting.demo.bean;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.stereotype.Component;

@Data
@Component
@AllArgsConstructor
@NoArgsConstructor
public class RespBean<T> {
    private String code;
    private String status;
    private String msg;
    private T data;

    public RespBean(String code, String status, String msg) {
        this.code = code;
        this.status = status;
        this.msg = msg;
    }

    public RespBean(String code, String status, T data) {
        this.code = code;
        this.status = status;
        this.data = data;
    }

    public RespBean(String code, String status) {
        this.code = code;
        this.status = status;
    }
}
