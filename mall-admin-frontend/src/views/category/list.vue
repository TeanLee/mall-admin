<template>
  <div class="category-list">
    <h1>分类列表</h1>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="category_id"
        label="ID">
      </el-table-column>
      <el-table-column
        prop="category"
        label="分类名">
      </el-table-column>
      <el-table-column
        prop="color"
        label="颜色">
      </el-table-column>
      <el-table-column
        prop="icon"
        label="icon">
      </el-table-column>
      <el-table-column
        prop="name"
        label="name">
      </el-table-column>
      <!-- <el-table-column label="banner">
        <template slot-scope="{row}">
          <el-link type="primary" :href="row.banner" target="_blank">{{ row.banner }}</el-link>
        </template>
      </el-table-column> -->
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-button @click="handleClick(scope.row)" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
// import axios from "axios"
import CategoryService from "@/service/category.service.js"

export default {
  name: 'category-list',
  data() {
    return {
      tableData: [],
    }
  },
  created() {
    CategoryService.getCategories().then(res => {
      const { data } = res
      this.tableData = data
    });
  },
  methods: {
    
  }
}
</script>

<style scoped lang="scss">
.category-list {
  padding: 0 20px;
}
</style>