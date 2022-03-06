package com.ting.demo.service;

import com.ting.demo.bean.Category;
import com.ting.demo.mapper.CategoryMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@Slf4j
public class CategoryService {
    @Autowired
    CategoryMapper categoryMapper;

    public List<Category> getAllCategories() {
        return categoryMapper.getAllCategories();
    }

    public int updateCategoryById(int categoryId, String category, String icon, String color, String name) {
        log.info("updateCategoryById category：" + categoryId);
        return categoryMapper.updateCategoryById(categoryId, category, icon, color, name);
    }
}
