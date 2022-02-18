<template>
  <div class="main-layout">
    <el-row style="height:100%">
      <el-col :span="3">
        <SiderBar />
      </el-col>
      <el-col :span="21">
        <NavBar />
        <AppMain />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import SiderBar from './SiderBar'
import AppMain from './AppMain'
import NavBar from './NavBar'
export default {
    name: 'layout',
    components: {
        SiderBar,
        AppMain,
        NavBar,
    }
}
</script>

<style lang="scss" scoped>
.main-layout {
    height: 100%;
}
.el-col {
  height: 100%;
  .el-menu {
    height: 100%;
  }
}
</style>