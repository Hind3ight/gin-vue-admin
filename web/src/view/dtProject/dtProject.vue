<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                                                                                        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增dtcloudProject表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="项目名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="项目简介" prop="profile" width="120"></el-table-column> 
    
    <el-table-column label="合同备案编号" prop="projectNumber" width="120"></el-table-column> 
    
    <el-table-column label="项目所在区县" prop="county" width="120"></el-table-column> 
    
    <el-table-column label="项目属地" prop="ssdq" width="120"></el-table-column> 
    
    <el-table-column label="项目地址" prop="projectAddress" width="120"></el-table-column> 
    
    <el-table-column label="建筑面积(平方米)" prop="floorage" width="120"></el-table-column> 
    
    <el-table-column label="工程造价(万元)" prop="constructionCost" width="120"></el-table-column> 
    
    <el-table-column label="建设单位名称" prop="jsdw" width="120"></el-table-column> 
    
    <el-table-column label="建设单位社会信用代码" prop="constructionCode" width="120"></el-table-column> 
    
    <el-table-column label="建设单位联系人" prop="jdwlxr" width="120"></el-table-column> 
    
    <el-table-column label="建设单位联系电话" prop="jsdwlxdh" width="120"></el-table-column> 
    
    <el-table-column label="经度" prop="lng" width="120"></el-table-column> 
    
    <el-table-column label="纬度" prop="lat" width="120"></el-table-column> 
    
    <el-table-column label="开工日期" prop="startdate" width="120"></el-table-column> 
    
    <el-table-column label="完工日期" prop="enddate" width="120"></el-table-column> 
    
    <el-table-column label="公司" prop="companyId" width="120"></el-table-column> 
    
    <el-table-column label="国家" prop="countryId" width="120"></el-table-column> 
    
    <el-table-column label="省" prop="stateId" width="120"></el-table-column> 
    
    <el-table-column label="市" prop="cityId" width="120"></el-table-column> 
    
    <el-table-column label="区" prop="areaId" width="120"></el-table-column> 
    
    <el-table-column label="城市" prop="city" width="120"></el-table-column> 
    
    <el-table-column label="备注" prop="note" width="120"></el-table-column> 
    
    <el-table-column label="大屏地址" prop="url" width="120"></el-table-column> 
    
    <el-table-column label="数据来源" prop="source" width="120"></el-table-column> 
    
    <el-table-column label="是否归档" prop="active" width="120"></el-table-column> 
    
    <el-table-column label="审核状态" prop="auditStatus" width="120"></el-table-column> 
    
    <el-table-column label="是否同步-工人" prop="isSynchroWorker" width="120"></el-table-column> 
    
    <el-table-column label="项目阶段" prop="pstate" width="120"></el-table-column> 
    
    <el-table-column label="项目类型" prop="pcate" width="120"></el-table-column> 
    
    <el-table-column label="安全指数" prop="safeindex" width="120"></el-table-column> 
    
    <el-table-column label="总承包商" prop="contractorId" width="120"></el-table-column> 
    
    <el-table-column label="地上面积(平方米)" prop="abovegroundArea" width="120"></el-table-column> 
    
    <el-table-column label="地下面积(平方米)" prop="undergroundArea" width="120"></el-table-column> 
    
    <el-table-column label="结构层数" prop="structureStoreys" width="120"></el-table-column> 
    
    <el-table-column label="结构类型" prop="structureType" width="120"></el-table-column> 
    
    <el-table-column label="前台名称" prop="frontName" width="120"></el-table-column> 
    
    <el-table-column label="获取的全部信息" prop="extraInfo" width="120"></el-table-column> 
    
    <el-table-column label="组织编号" prop="orgno" width="120"></el-table-column> 
    
    <el-table-column label="组织名称" prop="orgname" width="120"></el-table-column> 
    
    <el-table-column label="合同编号" prop="ctno" width="120"></el-table-column> 
    
    <el-table-column label="合同名称" prop="ctname" width="120"></el-table-column> 
    
    <el-table-column label="组织全称" prop="completeName" width="120"></el-table-column> 
    
    <el-table-column label="编号全称" prop="completeNo" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDtProject(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="项目名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目简介:">
            <el-input v-model="formData.profile" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="合同备案编号:">
            <el-input v-model="formData.projectNumber" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目所在区县:">
            <el-input v-model="formData.county" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目属地:">
            <el-input v-model="formData.ssdq" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目地址:">
            <el-input v-model="formData.projectAddress" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="建筑面积(平方米):">
            <el-input v-model="formData.floorage" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="工程造价(万元):">
            <el-input v-model="formData.constructionCost" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="建设单位名称:">
            <el-input v-model="formData.jsdw" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="建设单位社会信用代码:">
            <el-input v-model="formData.constructionCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="建设单位联系人:">
            <el-input v-model="formData.jdwlxr" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="建设单位联系电话:">
            <el-input v-model="formData.jsdwlxdh" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="经度:">
            <el-input v-model="formData.lng" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="纬度:">
            <el-input v-model="formData.lat" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="开工日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.startdate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="完工日期:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.enddate" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="公司:"><el-input v-model.number="formData.companyId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="国家:"><el-input v-model.number="formData.countryId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="省:"><el-input v-model.number="formData.stateId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="市:"><el-input v-model.number="formData.cityId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="区:"><el-input v-model.number="formData.areaId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="城市:">
            <el-input v-model="formData.city" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="备注:">
            <el-input v-model="formData.note" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="大屏地址:">
            <el-input v-model="formData.url" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="数据来源:">
            <el-input v-model="formData.source" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否归档:">
            <el-input v-model="formData.active" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="审核状态:">
            <el-input v-model="formData.auditStatus" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否同步-工人:">
            <el-input v-model="formData.isSynchroWorker" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目阶段:">
            <el-input v-model="formData.pstate" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="项目类型:">
            <el-input v-model="formData.pcate" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="安全指数:"><el-input v-model.number="formData.safeindex" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="总承包商:"><el-input v-model.number="formData.contractorId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="地上面积(平方米):">
            <el-input v-model="formData.abovegroundArea" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="地下面积(平方米):">
            <el-input v-model="formData.undergroundArea" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="结构层数:">
            <el-input v-model="formData.structureStoreys" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="结构类型:">
            <el-input v-model="formData.structureType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="前台名称:">
            <el-input v-model="formData.frontName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="获取的全部信息:">
            <el-input v-model="formData.extraInfo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="组织编号:">
            <el-input v-model="formData.orgno" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="组织名称:">
            <el-input v-model="formData.orgname" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="合同编号:">
            <el-input v-model="formData.ctno" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="合同名称:">
            <el-input v-model="formData.ctname" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="组织全称:">
            <el-input v-model="formData.completeName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="编号全称:">
            <el-input v-model="formData.completeNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createDtProject,
    deleteDtProject,
    deleteDtProjectByIds,
    updateDtProject,
    findDtProject,
    getDtProjectList
} from "@/api/dtProject";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "DtProject",
  mixins: [infoList],
  data() {
    return {
      listApi: getDtProjectList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            profile:"",
            projectNumber:"",
            county:"",
            ssdq:"",
            projectAddress:"",
            floorage:"",
            constructionCost:"",
            jsdw:"",
            constructionCode:"",
            jdwlxr:"",
            jsdwlxdh:"",
            lng:"",
            lat:"",
            startdate:new Date(),
            enddate:new Date(),
            companyId:0,
            countryId:0,
            stateId:0,
            cityId:0,
            areaId:0,
            city:"",
            note:"",
            url:"",
            source:"",
            active:"",
            auditStatus:"",
            isSynchroWorker:"",
            pstate:"",
            pcate:"",
            safeindex:0,
            contractorId:0,
            abovegroundArea:"",
            undergroundArea:"",
            structureStoreys:"",
            structureType:"",
            frontName:"",
            extraInfo:"",
            orgno:"",
            orgname:"",
            ctno:"",
            ctname:"",
            completeName:"",
            completeNo:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10                                                
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteDtProject(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteDtProjectByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length && this.page > 1) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateDtProject(row) {
      const res = await findDtProject({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redtProject;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          profile:"",
          projectNumber:"",
          county:"",
          ssdq:"",
          projectAddress:"",
          floorage:"",
          constructionCost:"",
          jsdw:"",
          constructionCode:"",
          jdwlxr:"",
          jsdwlxdh:"",
          lng:"",
          lat:"",
          startdate:new Date(),
          enddate:new Date(),
          companyId:0,
          countryId:0,
          stateId:0,
          cityId:0,
          areaId:0,
          city:"",
          note:"",
          url:"",
          source:"",
          active:"",
          auditStatus:"",
          isSynchroWorker:"",
          pstate:"",
          pcate:"",
          safeindex:0,
          contractorId:0,
          abovegroundArea:"",
          undergroundArea:"",
          structureStoreys:"",
          structureType:"",
          frontName:"",
          extraInfo:"",
          orgno:"",
          orgname:"",
          ctno:"",
          ctname:"",
          completeName:"",
          completeNo:"",
          
      };
    },
    async deleteDtProject(row) {
      const res = await deleteDtProject({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1 && this.page > 1 ) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDtProject(this.formData);
          break;
        case "update":
          res = await updateDtProject(this.formData);
          break;
        default:
          res = await createDtProject(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
