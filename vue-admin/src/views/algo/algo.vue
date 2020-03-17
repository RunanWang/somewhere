<template>
  <div class="app-container">
    <el-row>
         当前模型状态：{{ status }}
     </el-row>
     <el-row >
          <p></p>
          <p></p>
     </el-row>
    <el-row>
        <el-button type="warning" @click="handleTrain()">训练模型</el-button>
        <el-button type="primary" @click="handleRefresh()">更新状态</el-button>
     </el-row>
  </div>
</template>

<script>
import { trainModel } from '@/api/algo'

export default {
  data() {
    return {
      status: '模型正在服务中'
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
    },
    handleTrain() {
      console.log('训练！')
      trainModel().then(response => {
        // this.reload()
        console.log(response)
      })
    },
    handleRefresh() {
      console.log('更新！')
    }
  }
}
</script>
