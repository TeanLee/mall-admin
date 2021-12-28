<template>
  <el-menu
    :default-active="currentRoute"
    class="el-menu-vertical-demo"
    @open="handleOpen"
    @close="handleClose"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b">

    <sidebar-item v-for="(route, index) in routes" :key="route.name" :item="route" :base-path="route.path" :index="index + ''" />

  </el-menu>
</template>

<script>
import router from '@/router'
import SidebarItem from './SidebarItem.vue'

export default {
  name: "siderbar",
  components: { SidebarItem },
  data() {
    return {
      router,
      routes: router.options.routes
    }
  },
  computed: {
    currentRoute() {
      return this.$route.path
    }
  },
  methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    }
}
</script>
