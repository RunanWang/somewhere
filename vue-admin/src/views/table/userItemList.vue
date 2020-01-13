<template>
  <div id="app" class="app-container">
    <div class="filter-container" align="right" >
      <el-select align="left" v-model="pageInfo.store_id" clearable class="filter-item" placeholder="-请选择商铺-">
            <el-option v-for="item in storeOptions" :key="item.store_id" :label="item.store_name" :value="item.store_id" />
      </el-select>
      <el-button  class="filter-item" type="primary" icon="edit" @click="handleRefresh">刷新列表</el-button>
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
        <el-table-column label="商品ID" width="250" align="center">
          <template slot-scope="scope">
            {{ scope.row.item_id }}
          </template>
        </el-table-column>
        <el-table-column label="商品名称" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.item_name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="商品品牌" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.item_brand }}</span>
          </template>
        </el-table-column>
        <el-table-column label="价格" width="100" align="center">
          <template slot-scope="scope">
            <span>{{ scope.row.item_price }}</span>
          </template>
        </el-table-column>
        <el-table-column label="销量" width="100" align="center">
          <template slot-scope="scope">
            {{ scope.row.item_salecount }}
          </template>
        </el-table-column>
        <el-table-column class-name="status-col" label="评价" width="100" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.item_score | statusFilter">{{ scope.row.item_score }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="created_at" label="创建时间" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.item_timestamp | formatDate }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="100">
          <template scope="scope">
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 查看弹窗 -->
    <el-dialog title="查看商品" :visible.sync="dialogFormVisible2">
      <el-form class="small-space" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="商品名称">
          <el-input readonly="true" v-model="temp.item_name" />
        </el-form-item>
        <el-form-item label="商品品牌">
          <el-input readonly="true" v-model="temp.item_brand" />
        </el-form-item>

        <el-form-item label="商品价格">
          <el-input readonly="true" v-model.number="temp.item_price" type="number" step="0.01" />
        </el-form-item>
        <el-form-item label="商品评价">
          <el-input readonly="true" v-model.number="temp.item_score" type="number" step="0.01" max=5.0 />
        </el-form-item>
        <el-form-item label="商品销量">
          <el-input readonly="true" v-model.number="temp.item_salecount" type="number" step="1" />
        </el-form-item>
        <el-form-item label="所属商铺">
          <el-select :disabled="true" v-model="temp.store_id" clearable class="filter-item" placeholder="-请选择-" style="width: 332px; margin-left:0px;">
            <el-option v-for="item in storeOptions" :key="item.store_id" :label="item.store_name" :value="item.store_id" />
          </el-select>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible2 = false">取 消</el-button>
        <el-button type="primary" @click="handleEditSubmit(temp)">购 买</el-button>
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
import { getProductList, getProductListByPage, putProduct } from '@/api/product'
import { getStoreList } from '@/api/store'
import { parseTime } from '@/utils/index.js'
import { getBasic } from '@/api/basic.js'
import { postRecord } from '@/api/record.js'
import { mapGetters } from 'vuex'

export default {
  inject: ['reload'],
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
      listLoading: true,
      temp: {
        'item_id': '',
        'item_name': '',
        'item_price': 0,
        'item_score': 0,
        'item_brand': '',
        'store_id':'',
        'item_salecount': 0
      },
      record: {
          'pro_id': '',
          'user_id': '',
          'query': '',
          'is_trade': 0
      },
      listQuery: {
        store_id: ""
      },
      storeOptions: [],
      dialogFormVisible2: false,
      pageInfo:{
        'store_id':'',
        'page_num':1,
        'page_size':10
      },
      total: 10
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
    this.fetchData(this.pageInfo)
    console.log('userid：',this.id,this.name)
  },
  methods: {
    initTemp() {
      const vm = this

      vm.temp = {
        'item_id': '',
        'item_name': '',
        'item_price': 100,
        'item_score': 5.0,
        'item_brand': '',
        'store_id':'',
        'item_salecount':0
      }
    },

    handleEdit(index, row) {
      this.dialogFormVisible2 = true
      //const vm = this
      this.record.pro_id = row.item_id
      this.record.user_id = this.id
      
      console.log('编辑的row：', index, '-----', this.record)
      this.temp.item_id = row.item_id
      this.temp.item_name = row.item_name
      this.temp.item_price = row.item_price
      this.temp.item_score = row.item_score
      this.temp.item_brand = row.item_brand
      this.temp.store_id = row.store_id
      this.temp.item_salecount = row.item_salecount
      postRecord(this.record).then(response => {
      })
    },

    handlePageChange(val) {
      const vm = this
      this.pageInfo.page_num = val
      console.log('页面改变：', this.pageInfo.page_num,this.pageInfo.page_size)
      // 前端删除。
      getProductListByPage(vm.pageInfo).then(response => {
        this.list = response.list
        this.listLoading = false
      })
    },

    handleRefresh() {
      this.listLoading = true
      const vm = this
      console.log('页面改变：', this.pageInfo.store_id,this.pageInfo.page_num,this.pageInfo.page_size)
      // 前端删除。
      getProductListByPage(vm.pageInfo).then(response => {
        this.list = response.list
        this.listLoading = false
      })
    },

    handleSizeChange(val) {
      const vm = this
      this.pageInfo.page_size = val
      console.log('页面改变：', this.pageInfo.page_num,this.pageInfo.page_size)
      // 前端删除。
      getProductListByPage(vm.pageInfo).then(response => {
        this.list = response.list
        this.listLoading = false
      })
    },

    handleEditSubmit(temp) {
      const vm = this
      console.log('购买：', vm.temp)
      this.record.is_trade = 1
      postRecord(this.record).then(response => {
      })
      this.dialogFormVisible2 = false
    },

    fetchData(pageInfo) {
      this.listLoading = true
      getBasic().then(response => {
        this.total = response.item_num
      })
      getProductListByPage(pageInfo).then(response => {
        this.list = response.list
      })
      getStoreList().then(response => {
        this.storeOptions = response.data
        this.listLoading = false
      })
    }
  }
}
</script>
