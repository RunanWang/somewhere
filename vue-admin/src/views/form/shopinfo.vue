<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="商户名称">
        <el-input v-model="form.store_name" />
      </el-form-item>
      <!-- <el-form-item label="用户年龄">
        <el-input v-model.number="form.user_age" type="number" />
      </el-form-item>
      <el-form-item label="用户性别">
        <el-select v-model="form.user_gender" clearable class="filter-item" placeholder="-请选择-">
          <el-option v-for="item in genderOptions" :key="item.key" :label="item.display_name" :value="item.key" />
        </el-select>
      </el-form-item> -->
      <el-form-item label="所在城市">
          <el-input v-model="form.store_city" />
      </el-form-item>

      <!-- <el-form-item label="Activity zone">
        <el-select v-model="form.region" placeholder="please select your zone">
          <el-option label="Zone one" value="shanghai" />
          <el-option label="Zone two" value="beijing" />
        </el-select>
      </el-form-item> -->
      <!-- <el-form-item label="Activity time">
        <el-col :span="11">
          <el-date-picker v-model="form.date1" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
        <el-col :span="2" class="line">-</el-col>
        <el-col :span="11">
          <el-time-picker v-model="form.date2" type="fixed-time" placeholder="Pick a time" style="width: 100%;" />
        </el-col>
      </el-form-item> -->
      <!-- <el-form-item label="Instant delivery">
        <el-switch v-model="form.delivery" />
      </el-form-item>
      <el-form-item label="Activity type">
        <el-checkbox-group v-model="form.type">
          <el-checkbox label="Online activities" name="type" />
          <el-checkbox label="Promotion activities" name="type" />
          <el-checkbox label="Offline activities" name="type" />
          <el-checkbox label="Simple brand exposure" name="type" />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="Resources">
        <el-radio-group v-model="form.resource">
          <el-radio label="Sponsor" />
          <el-radio label="Venue" />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Activity form">
        <el-input v-model="form.desc" type="textarea" />
      </el-form-item> -->
      <el-form-item>
        <el-button type="primary" @click="onSubmit">修 改</el-button>
        <el-button @click="onCancel">取 消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getStoreInfo, putStoreList } from '@/api/store.js'
import { mapGetters } from 'vuex'

export default {
  inject: ['reload'],
  data() {
    return {
      form: {
        store_id: '',
        store_name: '',
        store_city: ''
      },
      pageInfo:{
        'store_id': '',
      }
    }
  },

  computed: {
    ...mapGetters([
      'name',
      'id',
      'roles'
    ])
  },

  created() {
    this.pageInfo.store_id = this.id
    this.form.store_id = this.id
    this.fetchData(this.pageInfo)
  },
  methods: {
    onSubmit() {
      const vm = this
      console.log('修改：', vm.form)
      this.$message('提交更改')
      putStoreList(this.form).then(response => {
        this.reload()
      })
    },
    onCancel() {
      this.reload()
      this.$message({
        message: '取消更改!'
      })
    },

    fetchData(pageInfo) {
      console.log(pageInfo)
      console.log(this.id)
      getStoreInfo(pageInfo).then(response => {
        console.log(pageInfo)
        this.form.store_name = response.store_name
        this.form.store_city = response.store_city
        // console.log(response)
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

