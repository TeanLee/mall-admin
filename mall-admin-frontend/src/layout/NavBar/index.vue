<template>
  <div class="navbar">
    <div class="username">
      <el-dropdown>
        <span class="el-dropdown-link">
          current user<i class="el-icon-arrow-down el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>退出登陆</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
export default {

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