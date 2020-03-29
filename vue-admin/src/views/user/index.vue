<template>
  <div>
    <div>
      <el-input
        v-model="listQuery.name"
        placeholder="请输入用户名称"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />
      <el-button class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >新建</el-button>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      style="width: 100%"
    >
      <el-table-column label="id" prop="id" sortable align="center" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" width="180" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="密码" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.password }}</span>
        </template>
      </el-table-column>
      <el-table-column label="头像" width="110" align="center">
        <template slot-scope="scope">
          <img :src="scope.row.avatar" style="width: 30px; height: 30px">
        </template>
      </el-table-column>
      <el-table-column
        prop="state"
        label="状态"
        width="110"
        align="center"
        :filters="[{ text: '正常', value: '正常' }, { text: '禁用', value: '禁用' }]"
        :filter-method="filterTag"
      >
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusFilter">{{ scope.row.state }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="user_type"
        label="用户类型"
        width="110"
        align="center"
        :filters="[{ text: '管理员', value: '管理员' }, { text: '测试用户', value: '测试用户' }]"
        :filter-method="filterType"
      >
        <template slot-scope="scope">
          <el-tag :type="scope.row.user_type | userTypeFilter">{{ scope.row.user_type }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="date" label="创建日期" sortable width="220" align="center">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span style="margin-left: 10px">{{ scope.row.created_on }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="fetchData"
    />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="temp.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="temp.password" />
        </el-form-item>
        <el-form-item label="权限" prop="user_type">
          <el-select v-model="temp.user_type" placeholder="请选择权限">
            <el-option label="管理员" :value="1" />
            <el-option label="测试用户" :value="2" />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getList, createUser, updateUser, deleteUser } from '@/api/user'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
export default {
  name: 'ComplexTable',
  components: { Pagination },
  filters: {
    statusFilter(Status) {
      const statusMap = {
        正常: 'success',
        禁用: 'danger'
      }
      return statusMap[Status]
    },
    userTypeFilter(userType) {
      const statusMap = {
        管理员: 'success',
        测试用户: 'primary'
      }
      return statusMap[userType]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑用户',
        create: '新建用户'
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        user_type: [{ required: true, message: '请选择权限', trigger: 'change' }]
      },
      temp: {
        id: undefined,
        username: '',
        password: '',
        user_type: undefined
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    handleEdit(index, row) {
      this.temp.id = row.id
      this.temp.username = row.username
      this.temp.password = row.password
      this.temp.user_type = row.user_type
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleDelete(index, row) {
      deleteUser(row.id).then(response => {
        this.fetchData()
        this.dialogFormVisible = false
        if (response.msg === 'fail') {
          this.$notify({
            title: 'Fail',
            message: response.detail,
            type: 'error',
            duration: 2000
          })
        } else {
          this.$notify({
            title: 'Success',
            message: '删除用户成功!',
            type: 'success',
            duration: 2000
          })
        }
      })
    },
    fetchData() {
      this.listLoading = true
      getList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.fetchData()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        username: '',
        password: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    filterTag(value, row) {
      return row.state === value
    },
    filterType(value, row) {
      return row.user_type === value
    },
    createData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          createUser(this.temp).then(() => {
            this.fetchData()
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '新建用户成功!',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    updateData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          if (tempData.user_type === '管理员') {
            tempData.user_type = 1
          } else if (tempData.user_type === '测试用户') {
            tempData.user_type = 2
          }
          updateUser(tempData).then(() => {
            this.fetchData()
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '更新数据成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    }
  }
}
</script>
