<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="项目:"><el-input v-model.number="formData.projectId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="排序:"><el-input v-model.number="formData.sequence" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="里程碑名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="启动时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.planStart" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="完成时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.planEnd" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="验收时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.checktime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="状态:">
                <el-input v-model="formData.state" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="剩余天数:"><el-input v-model.number="formData.remainday" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="实际启动时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.practicalStart" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="实际完成时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.practicalEnd" clearable></el-date-picker>
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
    createDtProjectMilepost,
    updateDtProjectMilepost,
    findDtProjectMilepost
} from "@/api/dtProjectMilepost";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DtProjectMilepost",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            projectId:0,
            sequence:0,
            name:"",
            planStart:new Date(),
            planEnd:new Date(),
            checktime:new Date(),
            state:"",
            remainday:0,
            practicalStart:new Date(),
            practicalEnd:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDtProjectMilepost(this.formData);
          break;
        case "update":
          res = await updateDtProjectMilepost(this.formData);
          break;
        default:
          res = await createDtProjectMilepost(this.formData);
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
    const res = await findDtProjectMilepost({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.redtProjectMilepost
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