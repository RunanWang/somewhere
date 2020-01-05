<template>
  <div id="app" class="app-container">
    <div>
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
        <el-table-column label="记录ID" width="250" align="center">
          <template slot-scope="scope">
            {{ scope.row.record_id }}
          </template>
        </el-table-column>
        <el-table-column label="用户ID" width="200" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.user_id }}</span>
          </template>
        </el-table-column>
        <el-table-column label="商品ID" width="200" align="center">
          <template slot-scope="scope">
            {{ scope.row.item_id }}
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="购买情况" width="130" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.is_trade | statusFilter">{{ scope.row.is_trade }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="created_at" label="记录时间">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.timestamp | formatDate }}</span>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { getRecordList } from '@/api/record'
import { parseTime } from '@/utils/index.js'
import { getBasic } from '@/api/basic.js'

export default {
  inject: ['reload'],
  filters: {
    statusFilter(Status) {
      if (Status == 0) {
        return 'danger'
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
      listLoading: true,
      temp: {
        'user_id': '',
        'name': '',
        'user_name': '',
        'user_age': 18,
        'user_gender': 0,
        'user_city': ''
      },
      total: 10
    }
  },

  created() {
    this.fetchData(this.pageInfo)
  },
  methods: {
    initTemp() {
      const vm = this

      vm.temp = {
        'user_id': '',
        'user_name': '',
        'user_age': 18,
        'user_gender': 0,
        'user_city': ''
      }
    },

    handleDelete(index, row) {
      const vm = this
      console.log('单个删除选择的row：', index, '-----', row)
      // 前端删除。
      // deleteUserList(row).then(response => {
      //   this.total--
      //   vm.list.splice(index,1)
      //   // this.reload()
      // })
    },

    fetchData(pageInfo) {
      this.listLoading = true
      getBasic().then(response => {
        this.total = response.record_num
      })
      getRecordList(pageInfo).then(response => {
        this.list = response.list
        this.listLoading = false
      })
    }
  }
}
</script>
