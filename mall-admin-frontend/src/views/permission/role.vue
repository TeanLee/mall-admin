<template>
  <div class="user-list">
    <h1>权限管理</h1>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="role"
        label="权限">
      </el-table-column>
      <el-table-column
        prop="username"
        label="用户名">
      </el-table-column>
      <el-table-column
        prop="password"
        label="密码">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-button @click="editRow(scope.row)" size="small">编辑</el-button>
          <el-button @click="handleClick(scope.row)" size="small" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button class="add-user" type="primary" @click.native="dialogFormVisible=true">新增用户</el-button>

    <el-dialog :title="dialogType === 'edit' ? '编辑用户' : '新增用户'" :visible.sync="dialogFormVisible" :close-on-click-modal="false" destroy-on-close ref="dialogForm" @close="closeDialog">
      <el-form :model="form" :rules="rules">
        <el-form-item label="用户角色" label-width="120px">
          <el-select v-model="form.role" placeholder="请选择活动区域">
            <el-option label="普通商户" value="1"></el-option>
            <el-option label="管理员" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="用户名" label-width="120px">
          <el-input v-model="form.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="120px">
          <el-input v-model="form.password" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'category-list',
  data() {
    return {
      tableData: [
        {
          role: '1',
          username: '11',
          password: '111'
        },
        {
          role: '2',
          username: '22',
          password: '222'
        }
      ],
      dialogFormVisible: false,
      form: {
        role: '',
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'change' }
        ]
      },
      dialogType: ''
    }
  },
  created() {
    
  },
  methods: {
    handleClick(row) {
      console.log(row);
    },
    editRow(row) {
      this.form = row;
      this.dialogType = 'edit';
      this.dialogFormVisible = true;
    },
    closeDialog() {
      this.form = {};
      this.dialogFormVisible = false;
      this.dialogType = '';
    },
  }
}
</script>

<style scoped lang="scss">
.user-list {
  padding: 0 20px;
  .add-user {
    float: left;
    margin-top: 20px;
  }
  .el-select {
    width: 100%;
  }
}
</style>