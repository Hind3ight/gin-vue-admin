<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createDtProject,
    updateDtProject,
    findDtProject
} from "@/api/dtProject";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DtProject",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findDtProject({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.redtProject
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>