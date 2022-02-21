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
    <el-dialog
      title="修改分类"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
      <el-form :model="dialogData">
        <el-form-item label="分类名" :label-width="formLabelWidth">
          <el-input v-model="dialogData.category" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="颜色" :label-width="formLabelWidth">
          <el-input v-model="dialogData.color" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="icon" :label-width="formLabelWidth">
          <el-input v-model="dialogData.icon" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleClose">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { cloneDeep } from "lodash"
import CategoryService from "@/service/category.service.js"

export default {
  name: 'category-list',
  data() {
    return {
      tableData: [],
      dialogVisible: false,
      dialogData: {},
      formLabelWidth: '120px'
    }
  },
  created() {
    this.getCategories()
  },
  methods: {
    getCategories() {
      CategoryService.getCategories().then(res => {
        const { data } = res
        this.tableData = data
      });
    },
    handleClose() {
      this.dialogVisible = false
      CategoryService.updateCategory(this.dialogData.category_id, this.dialogData).then(() => {
        this.getCategories()
      })
    },
    handleClick(row) {
      this.dialogVisible = true
      this.dialogData = cloneDeep(row)
    }
  }
}
</script>

<style scoped lang="scss">
.category-list {
  padding: 0 20px;
}
</style>