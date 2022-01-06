package com.example.demo;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;

@SpringBootApplication
@MapperScan("com.example.demo.mapper")  // 设置 mapper 文件的扫描路径，这样就不需要在每个 Mapper 文件上添加 @Mapper 注解了
public class MallAdminServerApplication {

	public static void main(String[] args) {
		SpringApplication.run(MallAdminServerApplication.class, args);
	}

}
