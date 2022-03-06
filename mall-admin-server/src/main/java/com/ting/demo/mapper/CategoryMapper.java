package com.ting.demo.mapper;

import com.ting.demo.bean.Category;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

@Mapper
public interface CategoryMapper {
    List<Category> getAllCategories();
    int updateCategoryById(@Param("categoryId") Integer categoryId, @Param("category") String category, @Param("icon") String icon, @Param("color") String color, @Param("name") String name);
}
