<template>
  <div class="login">
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="ruleForm.username"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="ruleForm.password"></el-input>
      </el-form-item>
      <el-button @click="login" type="primary">登录</el-button>
      <p>没有账号？<el-link href="https://element.eleme.io" target="_blank" type="primary">立即注册</el-link></p>
    </el-form>
  </div>
</template>

<script>
import AuthService from "@/service/auth.service.js"

export default {
  data() {
    return {
      ruleForm: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
        ],
      }
    }
  },
  methods: {
    login() {
      AuthService.login({
        username: this.ruleForm.username,
        password: this.ruleForm.password
      }).then(res => {
        console.log("res：", res)
        this.$router.push('/category/list')
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.login {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  .el-form {
    width: 40%;
    p {
      font-size: 14px;
    }
  }
}
</style>