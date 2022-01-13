package com.ting.demo.mapper;

import com.ting.demo.bean.Category;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;

@Mapper
public interface CategoryMapper {
    List<Category> getAllCategories();
}
