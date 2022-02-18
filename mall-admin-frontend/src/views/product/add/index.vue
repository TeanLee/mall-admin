<template>
  <div class="product-add">
    <h1>上架商品</h1>
    <div class="add-form">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="商品名" prop="title">
          <el-input v-model="ruleForm.title"></el-input>
        </el-form-item>
        <el-form-item label="商品描述" prop="subTitle">
          <el-input v-model="ruleForm.subTitle"></el-input>
        </el-form-item>
        <el-form-item label="商品价格" prop="price">
          <el-input v-model="ruleForm.price"></el-input>
        </el-form-item>
        <el-form-item label="商品单位" prop="unit">
          <el-input v-model="ruleForm.banner"></el-input>
        </el-form-item>
        <el-form-item label="商品主图" prop="banner">
          <el-input v-model="ruleForm.banner"></el-input>
        </el-form-item>
        <el-form-item label="商品分类" prop="category">
          <el-select v-model="ruleForm.region" placeholder="请选择商品分类">
            <el-option label="区域一" value="shanghai"></el-option>
            <el-option label="区域二" value="beijing"></el-option>
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
export default {
  data() {
    return {
      ruleForm: {
        title: '',
        subTitle: '',
        price: '',
        unit: '',
        banner: '',
        category: ''
      },
      rules: {
        title: [
          { required: true, message: '请输入商品名', trigger: 'blur' },
          // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        price: [
          { required: true, message: '请输入商品价格', trigger: 'change' }
        ],
        unit: [
          { required: true, message: '请输入计量单位', trigger: 'blur' },
          { max: 1, message: '长度为 1 个字符', trigger: 'blur' }
        ],
        banner: [
          { required: true, message: '请输入商品主图', trigger: 'change' }
        ],
        category: [
          { required: true, message: '请选择活动区域', trigger: 'change' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!');
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