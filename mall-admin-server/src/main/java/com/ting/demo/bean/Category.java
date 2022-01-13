package com.ting.demo.bean;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Category {
    private int categoryId;
    private String category;
    private String icon;
    private String color;
    private String name;
}
