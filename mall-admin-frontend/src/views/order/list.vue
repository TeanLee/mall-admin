<template>
  <div class="order-list">
    <h1>订单列表</h1>
    <!-- <div class="header-search">
      <div class="tip">
        <div class="left">
          <i class="el-icon-search"></i>
          <span>筛选搜索</span>
        </div>
        <div class="right">
          <el-button size="small">重置</el-button>
          <el-button type="primary" size="small">查询结果</el-button>
        </div>
      </div>
      <div class="inputs">
        <el-row>
          <el-col :span="10"><div class="demo-input-suffix">商品名称：<el-input v-model="productName" placeholder="请输入内容" style="display:inline"></el-input></div></el-col>
          <el-col :span="4"><div class="grid-content bg-purple-dark"></div></el-col>
          <el-col :span="10">
            <div class="demo-input-suffix">商品分类：
              <el-select v-model="value" placeholder="请选择">
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </div>
          </el-col>
        </el-row>
      </div>
    </div> -->
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="order_time"
        label="下单时间">
        <template slot-scope="scope">
          {{ timestampToTime(scope.row.order_time) }}
        </template>
      </el-table-column>
      <el-table-column
        prop="user_name"
        label="用户名"
        width="80">
      </el-table-column>
      <el-table-column
        prop="receiver"
        label="收件人"
        width="80">
      </el-table-column>
      <el-table-column
        prop="address"
        label="收货地址">
      </el-table-column>
     <el-table-column prop="order_items" label="订单列表">
        <template slot-scope="scope">
          <el-table
            :data="scope.row.order_items"
            border
            style="width: 100%">
            <el-table-column
              prop="count"
              label="数量"
              width="30">
            </el-table-column>
            <el-table-column
              prop="product"
              label="商品名"
              width="300">
              <template slot-scope="scope">
               {{ scope.row.product.title }}
              </template>
            </el-table-column>
          </el-table>
          <!-- <span v-for="(item, index) in scope.row.order_items" :key="index">
            {{ item.count }}--{{ item.product.title }}
          </span> -->
        </template>
      </el-table-column>
      <el-table-column
        prop="address"
        label="地址">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-button @click="handleClick(scope.row)" size="small">编辑</el-button>
          <el-button type="danger" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import OrderService from "@/service/order.service.js"
export default {
  created() {
    this.getOrders()
  },
  methods: {
    handleClick(row) {
      console.log(row);
    },
    getOrders() {
      OrderService.getOrders(this.currentPage, this.pageSize).then(res => {
        const { data } = res
        this.tableData = data
      });
    },
    timestampToTime(timestamp) {
      var date = new Date(parseInt(timestamp));//时间戳为10位需*1000，时间戳为13位的话不需乘1000
      var Y = date.getFullYear() + '-';
      var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
      var D = date.getDate() + ' ';
      var h = date.getHours() + ':';
      var m = date.getMinutes();
      return Y + M + D + h + m;
    }
  },

  data() {
    return {
      pageSize: 6,
      currentPage: 0,
      productName: '',
      options: [{
        value: '选项1',
        label: '黄金糕'
      }, {
        value: '选项2',
        label: '双皮奶'
      }, {
        value: '选项3',
        label: '蚵仔煎'
      }, {
        value: '选项4',
        label: '龙须面'
      }, {
        value: '选项5',
        label: '北京烤鸭'
      }],
      value: '',
      tableData: []
    }
  }
}
</script>

<style lang="scss" scoped>
.order-list {
  padding: 0 20px;

  .header-search {
    width: 100%;
    height: 120px;
    border: 1px solid #EBEEF5;
    margin-bottom: 30px;
    padding: 15px 100px 15px 10px;

    .tip {
      width: 100%;
      height: 30px;
      .left {
        float: left;
      }
      .right {
        float: right;
      }
    }

    .inputs {
      padding-top: 20px;
    }
  }
}
.demo-input-suffix {
  display: flex;
  align-items: center;
  justify-content: center;
  .el-input {
    width: 70%;
  }
  .el-select {
    width: 70%;
  }
}
</style>