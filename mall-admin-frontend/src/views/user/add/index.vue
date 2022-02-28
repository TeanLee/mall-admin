<template>
  <div class="user-add">
    <h1>新增用户</h1>
    <div class="add-form">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="ruleForm.username"></el-input>
        </el-form-item>
        <el-form-item label="收货地址" prop="address">
          <el-input v-model="ruleForm.address"></el-input>
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="ruleForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="收件人" prop="receiver">
          <el-input v-model="ruleForm.receiver"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import UserService from "@/service/user.service.js"

export default {
  data() {
    return {
      ruleForm: {
        username: '',
        address: '',
        phone: '',
        receiver: '',
      },
      categories: [],
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        address: [
          { required: true, message: '请输入收货地址', trigger: 'change' }
        ],
        phone: [
          { required: true, message: '请输入联系电话', trigger: 'change' }
        ],
        receiver: [
          { required: true, message: '请输入收件人', trigger: 'change' }
        ],
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          UserService.addUser(this.ruleForm).then(() => {
            this.resetForm('ruleForm')
            this.$message({
              showClose: true,
              message: '用户新增成功！',
              type: 'success'
            });
          })
        } else {
          this.$message({
            showClose: true,
            message: 'error submit!!',
            type: 'error'
          });
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>

<style lang="scss" scoped>
.user-add {
  padding: 0 20px;
  .add-form {
    width: 50%;
    margin: 0 auto;
  }
  .el-select {
    width: 100%;
  }
}
</style>