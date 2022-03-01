<template>
  <div class="navbar">
    <div class="username">
      <el-dropdown @command="handleCommand">
        <span class="el-dropdown-link">
          {{ username }}<i class="el-icon-arrow-down el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="logout">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import AuthService from "@/service/auth.service.js"

export default {
  computed: {
    ...mapGetters(['username'])
  },
  methods: {
    handleCommand(command) {
      if(command != 'logout') return
      AuthService.logout().then(() => {
        this.$store.dispatch('deleteUserInfo');
        this.$router.push('/')
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  width: 100%;
  height: 40px;
  border-bottom: 1px solid #EBEEF5;
  .username {
    float: right;
    line-height: 40px;
    margin-right: 20px;
    cursor: pointer;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
}
</style>