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
      <el-table-column label="商铺ID" width="250" align="center">
        <template slot-scope="scope">
          {{ scope.row.store_id }}
        </template>
      </el-table-column>
      <el-table-column label="商铺名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.store_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="商铺城市" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.store_city }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="商铺评分" width="100" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.store_level | statusFilter">{{ scope.row.store_level }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="创建时间" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.timestamp | formatDate }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getList } from '@/api/table'
import { parseTime } from '@/utils/index.js'

export default {
  filters: {
    statusFilter(Status) {
      if (Status <= 2) {
        return 'danger'
      }
      if (Status <= 4) {
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
      getList().then(response => {
        this.list = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
