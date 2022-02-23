<template>
  <div class="product-list">
    <h1>商品列表</h1>
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
      style="width: 100%"
      max-height="100%">
      <el-table-column
        prop="product_id"
        label="ID">
      </el-table-column>
      <el-table-column
        label="商品名">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="light" :content="scope.row.title" placement="top">
            <span style="margin-left: 10px; overflow: hidden; text-overflow: ellipsis; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical;">{{ scope.row.title }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        prop="category_id"
        label="商品分类">
      </el-table-column>
      <el-table-column
        prop="price"
        label="现价">
      </el-table-column>
      <el-table-column
        prop="old_price"
        label="原价">
      </el-table-column>
      <el-table-column
        prop="sub_title"
        label="次标题">
      </el-table-column>
      <el-table-column
        label="封面图">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="light" :content="scope.row.banner" placement="top">
            <span style="margin-left: 10px; overflow: hidden; text-overflow: ellipsis; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical;">{{ scope.row.banner }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        prop="unit"
        label="计量单位">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-button @click="handleUpdate(scope.row)" size="small">编辑</el-button>
          <el-button @click="handleDelete(scope.row)" type="danger" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      style="margin-top: 15px;"
      background
      layout="prev, pager, next"
      @current-change="handleCurrentChange"
      @prev-click="prevClick"
      @next-click="nextClick"
      :page-size="pageSize"
      :total="total">
    </el-pagination>

    <el-dialog
      title="修改分类"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
      <el-form :model="dialogData">
        <el-form-item label="商品名" :label-width="formLabelWidth">
          <el-input v-model="dialogData.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="商品分类" :label-width="formLabelWidth">
          <el-input v-model.number="dialogData.category_id" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="现价" :label-width="formLabelWidth">
          <el-input v-model.number="dialogData.price" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="原价" :label-width="formLabelWidth">
          <el-input v-model.number="dialogData.old_price" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="次标题" :label-width="formLabelWidth">
          <el-input v-model="dialogData.sub_title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="封面图" :label-width="formLabelWidth">
          <el-input v-model="dialogData.banner" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="计量单位" :label-width="formLabelWidth">
          <el-input v-model="dialogData.unit" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleClose">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { cloneDeep } from "lodash"
import ProductService from "@/service/product.service.js"

export default {
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
      tableData: [],
      dialogVisible: false,
      dialogData: {},
      formLabelWidth: '120px',
      pageSize: 6,
      currentPage: 0,
      total: 0,
    }
  },
  created() {
    this.getProducts()
  },
  methods: {
    handleUpdate(row) {
      this.dialogVisible = true
      this.dialogData = cloneDeep(row)
    },
    getProducts() {
      ProductService.getProducts(this.currentPage, this.pageSize).then(res => {
        const { data, total } = res
        this.tableData = data
        this.total = total
      })
    },
    handleClose() {
      this.dialogVisible = false
      ProductService.updateProduct(this.dialogData.product_id, this.dialogData).then(() => {
        this.getProducts()
      })
    },
    prevClick() {
      this.currentPage --
      this.getProducts()
    },
    nextClick() {
      this.currentPage ++
      this.getProducts()
    },
    handleCurrentChange(val) {
      this.currentPage = val - 1
      this.getProducts()
    }
  }
}
</script>

<style lang="scss" scoped>
.product-list {
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