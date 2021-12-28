<template>
  <div v-if="!item.hidden">
    <template v-if="hasOneShowingChild(item.children,item) && (!onlyOneChild.children||onlyOneChild.noShowingChildren)">
        <router-link :to="resolvePath(onlyOneChild.path)">
        <el-menu-item :index="resolvePath(onlyOneChild.path)">
            <i :class="onlyOneChild.meta.icon"></i>
            <span slot="title">{{onlyOneChild.title}}</span>
        </el-menu-item>
        </router-link>
        
    </template>
    
    <el-submenu v-else ref="subMenu" :index="resolvePath(onlyOneChild.path)" popper-append-to-body>
      <template slot="title">
        <i class="el-icon-location"></i>
        <span>{{item.title}}</span>
      </template>
      <sidebar-item
        v-for="(child, subIndex) in item.children"
        :key="child.path"
        :is-nest="true"
        :item="child"
        :base-path="resolvePath(child.path)"
        class="nest-menu"
        :index="index + '' + subIndex"
      />
    </el-submenu>
  </div>
</template>

<script>
import path from 'path'

export default {
    name: 'sidebar-item',
    props: {
        item: {
            type: Object,
            required: true
        },
        basePath: {
            type: String,
            default: ''
        },
        // 主菜单栏的第几位，用户菜单定位
        index: {
            type: String,
            default: ''
        }
    },
    data() {
        this.onlyOneChild = null
        return {}
    },
    created() {
        // console.log('created', this.item)
    },
    methods: {
        hasOneShowingChild(children = [], parent) {
            const showingChildren = children.filter(item => {
                if (item.hidden) {
                    return false
                } else {
                // Temp set(will be used if only has one showing child)
                    this.onlyOneChild = item
                    return true
                }
            })

            // When there is only one child router, the child router is displayed by default
            if (showingChildren.length === 1) {
                console.log('showingChildren', this.onlyOneChild)
                return true
            }

            // Show parent if there are no child router to display
            if (showingChildren.length === 0) {
                this.onlyOneChild = { ... parent, path: '', noShowingChildren: true }
                console.log('showingChildren.length === 0', this.onlyOneChild)
                return true
            }

            return false
        },
        resolvePath(routePath) {
            return path.resolve(this.basePath, routePath)
        }
    }
}
</script>

<style lang="scss" scoped>
a {
    text-decoration: none;
}
</style>