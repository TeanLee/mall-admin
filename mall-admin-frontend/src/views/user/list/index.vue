<template>
  <div class="user-list">
    <h1>用户列表</h1>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="username"
        label="用户名">
      </el-table-column>
      <el-table-column
        prop="address"
        label="收货地址">
      </el-table-column>
      <el-table-column
        prop="phone"
        label="联系电话">
      </el-table-column>
      <el-table-column
        prop="receiver"
        label="收件人">
      </el-table-column>
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
      title="修改用户信息"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
      <el-form :model="dialogData">
        <el-form-item label="用户名" :label-width="formLabelWidth">
          <el-input v-model="dialogData.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="收货地址" :label-width="formLabelWidth">
          <el-input v-model="dialogData.adderss" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="联系电话" :label-width="formLabelWidth">
          <el-input v-model="dialogData.phone" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="收件人" :label-width="formLabelWidth">
          <el-input v-model="dialogData.receiver" autocomplete="off"></el-input>
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
import UserService from "@/service/user.service.js"

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
      UserService.getUsers().then(res => {
        const { data } = res
        this.tableData = data
      });
    },
    handleClose() {
      // this.dialogVisible = false
      // CategoryService.updateCategory(this.dialogData.category_id, this.dialogData).then(() => {
      //   this.getCategories()
      // })
    },
    handleClick(row) {
      this.dialogVisible = true
      this.dialogData = cloneDeep(row)
    }
  }
}
</script>

<style>
.user-list {
  padding: 0 20px;
}
</style>