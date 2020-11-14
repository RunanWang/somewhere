<template>
  <div id="app" class="app-container">
    <div class="filter-container" align="right">
      <el-button class="filter-item" type="primary" icon="edit" @click="handleCreate">新建广告主</el-button>
      <p />
    </div>
    <div>
      <el-table
        v-loading="listLoading"
        :data="data"
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
        <el-table-column label="广告主ID" width="250" align="center">
          <template slot-scope="scope">
            {{ scope.row.store_id }}
          </template>
        </el-table-column>
        <el-table-column label="广告主名称" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.store_name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="广告主城市" width="100" align="center">
          <template slot-scope="scope">
            {{ scope.row.store_city }}
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="广告主评分" width="130" align="center">
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
        <el-table-column align="center" label="操作" width="200">
          <template scope="scope">
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增弹窗 -->
    <el-dialog title="新建广告主" :visible.sync="dialogFormVisible">
      <el-form class="small-space" :model="temp" label-position="left" label-width="90px" style="width: 400px; margin-left:50px;">
        <el-form-item label="广告主账号">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="广告主名称">
          <el-input v-model="temp.store_name" />
        </el-form-item>

        <el-form-item label="广告主评分">
          <el-input v-model.number="temp.store_level" type="number" step="0.01" max=5.0 />
        </el-form-item>

        <el-form-item label="所在城市">
          <el-input v-model="temp.store_city" />
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleCreateSubmit(temp)">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 修改弹窗 -->
    <el-dialog title="编辑广告主" :visible.sync="dialogFormVisible2">
      <el-form class="small-space" :model="temp" label-position="left" label-width="90px" style="width: 400px; margin-left:50px;">
        <el-form-item label="广告主名称">
          <el-input v-model="temp.store_name" />
        </el-form-item>

        <el-form-item label="广告主评分">
          <el-input v-model.number="temp.store_level" type="number" step="0.01" max=5.0 />
        </el-form-item>

        <el-form-item label="所在城市">
          <el-input v-model="temp.store_city" />
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible2 = false">取 消</el-button>
        <el-button type="primary" @click="handleEditSubmit(temp)">确 定</el-button>
      </div>
    </el-dialog>
    <el-pagination
      :current-page="pageInfo.page_num"
      :page-size="pageInfo.page_size"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      @current-change="handlePageChange"
      @size-change="handleSizeChange"
    />
  </div>
</template>

<script>
import { getStoreList, getStoreListByPage, postStoreList, deleteStoreList, putStoreList } from '@/api/store'
import { parseTime } from '@/utils/index.js'
import { getBasic } from '@/api/basic.js'

export default {
  inject: ['reload'],
  filters: {
    statusFilter(Status) {
      if (Status <= 2) {
        return 'danger'
      }
      if (Status <= 3) {
        return 'warning'
      }
      if (Status <= 4) {
        return 'primary'
      }
      return 'success'
    },
    formatDate(time) {
      return parseTime(time)
    }
  },
  data() {
    return {
      data: null,
      listLoading: true,
      temp: {
        'store_id': '',
        'name': '',
        'store_name': '',
        'store_level': 5.0,
        'store_city': ''
      },
      dialogFormVisible: false,
      dialogFormVisible2: false,
      pageInfo:{
        'page_num':1,
        'page_size':10
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
        'store_id': '',
        'store_name': '',
        'store_level': 5,
        'store_city': ''
      }
    },

    handleEdit(index, row) {
      this.dialogFormVisible2 = true
      // const vm = this
      console.log('编辑的row：', index, '-----', row)
      this.temp.store_id = row.store_id
      this.temp.store_city = row.store_city
      this.temp.store_level = row.store_level
      this.temp.store_name = row.store_name
    },

    handleCreate() {
      // 每次点击新增时 重置下新增表单数据，避免双向绑定的影响
      this.initTemp()
      this.dialogFormVisible = true
    },

    handleDelete(index, row) {
      const vm = this
      console.log('单个删除选择的row：', index, '-----', row)
      // 前端删除。
      deleteStoreList(row).then(response => {
        this.total--
        vm.data.splice(index,1)
        // this.reload()
      })
    },

    handlePageChange(val) {
      const vm = this
      this.pageInfo.page_num = val
      console.log('页面改变：', this.pageInfo.page_num,this.pageInfo.page_size)
      // 前端删除。
      getStoreListByPage(vm.pageInfo).then(response => {
        this.data = response.data
        this.listLoading = false
      })
    },

    handleSizeChange(val) {
      const vm = this
      this.pageInfo.page_size = val
      console.log('页面改变：', this.pageInfo.page_num,this.pageInfo.page_size)
      // 前端删除。
      getStoreListByPage(vm.pageInfo).then(response => {
        this.data = response.data
        this.listLoading = false
      })
    },

    handleCreateSubmit(temp) {
      const vm = this
      console.log('新增广告主：', vm.temp)
      postStoreList(temp).then(response => {
        this.reload()
      })
      this.dialogFormVisible = false
    },

    handleEditSubmit(temp) {
      const vm = this
      console.log('修改用户：', vm.temp)
      // this.isD
      putStoreList(temp).then(response => {
        this.reload()
      })
      this.dialogFormVisible2 = false
    },

    fetchData(pageInfo) {
      this.listLoading = true
      getBasic().then(response => {
        this.total = response.shop_num
      })
      getStoreListByPage(pageInfo).then(response => {
        this.data = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
