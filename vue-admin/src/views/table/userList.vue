<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="序号" width="95">
        <template slot-scope="scope">
          {{ scope.$index+1 }}
        </template>
      </el-table-column>
      <el-table-column label="用户ID" width="250" align="center">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
      <el-table-column label="用户名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.user_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户城市" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.user_city }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="用户历史消费" width="130" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.user_historysum | statusFilter">{{ scope.row.user_historysum }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="创建时间" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.user_timestamp | formatDate }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getUserList } from '@/api/table'
import { parseTime } from '@/utils/index.js'

export default {
  filters: {
    statusFilter(Status) {
      if (Status <= 0) {
        return 'danger'
      }
      if (Status <= 1000) {
        return 'gray'
      }
      return 'success'
    },
    formatDate(time) {
      return parseTime(time)
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getUserList().then(response => {
        this.list = response.list
        this.listLoading = false
      })
    }
  }
}
</script>
