<template>
  <div class="order-list">
    <h1>订单列表</h1>
    <div class="header-search">
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
    </div>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        prop="date"
        label="日期">
      </el-table-column>
      <el-table-column
        prop="name"
        label="姓名">
      </el-table-column>
      <el-table-column
        prop="province"
        label="省份">
      </el-table-column>
      <el-table-column
        prop="city"
        label="市区">
      </el-table-column>
      <el-table-column
        prop="address"
        label="地址">
      </el-table-column>
      <el-table-column
        prop="zip"
        label="邮编">
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
  export default {
    methods: {
      handleClick(row) {
        console.log(row);
      }
    },

    data() {
      return {
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
        tableData: [{
          date: '2016-05-02',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-04',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1517 弄',
          zip: 200333
        }, {
          date: '2016-05-01',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1519 弄',
          zip: 200333
        }, {
          date: '2016-05-03',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1516 弄',
          zip: 200333
        }]
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