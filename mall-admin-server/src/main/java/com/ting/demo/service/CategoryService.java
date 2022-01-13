package com.ting.demo.service;

import com.ting.demo.bean.Category;
import com.ting.demo.mapper.CategoryMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CategoryService {
    @Autowired
    CategoryMapper categoryMapper;

    public List<Category> getAllCategories() {
        List<Category> categoryList = categoryMapper.getAllCategories();
        return categoryList;
    }
}
