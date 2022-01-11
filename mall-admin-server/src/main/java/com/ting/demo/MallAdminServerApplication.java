package com.ting.demo;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.ting.demo.mapper")
public class MallAdminServerApplication {

    public static void main(String[] args) {
        SpringApplication.run(MallAdminServerApplication.class, args);
    }

}
