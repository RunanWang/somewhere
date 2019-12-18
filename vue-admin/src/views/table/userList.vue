<template>
  <div id="app" class="app-container">
    <div class="filter-container" align="right">
      <el-button class="filter-item" type="primary" icon="edit" @click="handleCreate">新建用户</el-button>
      <p />
    </div>
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
        <el-table-column align="center" label="操作" width="200">
          <template scope="scope">
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增弹窗 -->
    <el-dialog title="新建用户" :visible.sync="dialogFormVisible">
      <el-form class="small-space" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="用户账号">
          <el-input v-model="temp.name" />
        </el-form-item>
        
        <el-form-item label="用户名称">
          <el-input v-model="temp.user_name" />
        </el-form-item>

        <el-form-item label="用户年龄">
          <el-input v-model.number="temp.user_age" type="number" />
        </el-form-item>

        <el-form-item label="用户性别">
          <el-select v-model="temp.user_gender" clearable class="filter-item" placeholder="-请选择-">
            <el-option v-for="item in genderOptions" :key="item.key" :label="item.display_name" :value="item.key" />
          </el-select>
        </el-form-item>

        <el-form-item label="所在城市">
          <el-input v-model="temp.user_city" />
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleCreateSubmit(temp)">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 修改弹窗 -->
    <el-dialog title="编辑用户" :visible.sync="dialogFormVisible2">
      <el-form class="small-space" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="用户名称">
          <el-input v-model="temp.user_name" />
        </el-form-item>

        <el-form-item label="用户年龄">
          <el-input v-model.number="temp.user_age" type="number" />
        </el-form-item>

        <el-form-item label="用户性别">
          <el-select v-model="temp.user_gender" clearable class="filter-item" placeholder="-请选择-">
            <el-option v-for="item in genderOptions" :key="item.key" :label="item.display_name" :value="item.key" />
          </el-select>
        </el-form-item>

        <el-form-item label="所在城市">
          <el-input v-model="temp.user_city" />
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible2 = false">取 消</el-button>
        <el-button type="primary" @click="handleEditSubmit(temp)">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getUserList, postUserList, deleteUserList, putUserList } from '@/api/table'
import { parseTime } from '@/utils/index.js'

export default {
  inject: ['reload'],
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
      listLoading: true,
      temp: {
        'user_id': '',
        'name': '',
        'user_name': '',
        'user_age': 18,
        'user_gender': 0,
        'user_city': ''
      },
      listQuery: {
        user_gender: 0
      },
      genderOptions: [
        { key: 0, display_name: '-请选择-' },
        { key: 1, display_name: '男' },
        { key: 2, display_name: '女' }
      ],
      dialogFormVisible: false,
      dialogFormVisible2: false
    }
  },

  created() {
    this.fetchData()
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

    handleEdit(index, row) {
      this.dialogFormVisible2 = true
      // const vm = this
      console.log('编辑的row：', index, '-----', row)
      this.temp.user_id = row.user_id
      this.temp.user_gender = row.user_gender
      this.temp.user_city = row.user_city
      this.temp.user_age = row.user_age
      this.temp.user_name = row.user_name
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
      deleteUserList(row).then(response => {
        vm.list.splice(index, 1)
      })
    },

    handleCreateSubmit(temp) {
      const vm = this
      console.log('新增用户：', vm.temp)
      // this.isD
      postUserList(temp).then(response => {
        this.reload()
      })
      this.dialogFormVisible = false
    },

    handleEditSubmit(temp) {
      const vm = this
      console.log('修改用户：', vm.temp)
      // this.isD
      putUserList(temp).then(response => {
        this.reload()
      })
      this.dialogFormVisible2 = false
    },

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
