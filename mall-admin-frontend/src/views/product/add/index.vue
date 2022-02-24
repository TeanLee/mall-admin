<template>
  <div class="product-add">
    <h1>上架商品</h1>
    <div class="add-form">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="商品名" prop="title">
          <el-input v-model="ruleForm.title"></el-input>
        </el-form-item>
        <el-form-item label="商品描述" prop="sub_title">
          <el-input v-model="ruleForm.sub_title"></el-input>
        </el-form-item>
        <el-form-item label="商品价格" prop="price">
          <el-input v-model.number="ruleForm.price"></el-input>
        </el-form-item>
        <el-form-item label="商品原价" prop="old_price">
          <el-input v-model.number="ruleForm.old_price"></el-input>
        </el-form-item>
        <el-form-item label="商品单位" prop="unit">
          <el-input v-model="ruleForm.unit"></el-input>
        </el-form-item>
        <el-form-item label="商品主图" prop="banner">
          <el-input v-model="ruleForm.banner"></el-input>
        </el-form-item>
        <el-form-item label="商品分类" prop="category_id">
          <el-select v-model="ruleForm.category_id" placeholder="请选择">
            <el-option
                  v-for="item in categories"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
          </el-select>
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
import ProductService from "@/service/product.service.js"
import CategoryService from "@/service/category.service.js"

export default {
  data() {
    return {
      ruleForm: {
        title: '',
        sub_title: '',
        price: '',
        old_price: '',
        unit: '',
        banner: '',
        category_id: '',
      },
      categories: [],
      rules: {
        title: [
          { required: true, message: '请输入商品名', trigger: 'blur' },
          // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        price: [
          { required: true, message: '请输入商品价格', trigger: 'change' }
        ],
        old_price: [
          { required: true, message: '请输入商品价格', trigger: 'change' }
        ],
        unit: [
          { required: true, message: '请输入计量单位', trigger: 'blur' },
          { max: 1, message: '长度为 1 个字符', trigger: 'blur' }
        ],
        banner: [
          { required: true, message: '请输入商品主图', trigger: 'change' }
        ],
        category_id: [
          { required: true, message: '请选择商品分类', trigger: 'change' }
        ]
      }
    };
  },
  created() {
    this.getCategories()
  },
  methods: {
    getCategories() {
      CategoryService.getCategories().then(res => {
        const { data } = res
        this.categories = data.map(item => {
          return {
            "label": item.category,
            "value": item.category_id
          }
        })
      })
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          ProductService.addProduct(this.ruleForm)
        } else {
          console.log('error submit!!');
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
.product-add {
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