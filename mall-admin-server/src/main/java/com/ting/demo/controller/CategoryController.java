package com.ting.demo.controller;

import com.ting.demo.bean.Category;
import com.ting.demo.bean.RespBean;
import com.ting.demo.service.CategoryService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Slf4j
@Controller
@RestController
@RequestMapping("/category")
public class CategoryController {
    @Autowired
    CategoryService categoryService;

    @GetMapping("/list")
    public RespBean getAllCategories() {
        List<Category> categoryList = categoryService.getAllCategories();
        return new RespBean("200", "success", categoryList);
    }

    @PostMapping("/{categoryId}")
    public RespBean updateCategoryById(@PathVariable("categoryId") int categoryId, @RequestParam(required = false, value = "category", defaultValue = "") String category, @RequestParam(required = false, value = "icon", defaultValue = "") String icon, @RequestParam(required = false, value = "color", defaultValue = "") String color, @RequestParam(required = false, value = "name", defaultValue = "") String name) {
        int updateRes = categoryService.updateCategoryById(categoryId, category, icon, color, name);
        return new RespBean("200", "success", updateRes);
    }
}
