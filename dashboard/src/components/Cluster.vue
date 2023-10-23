<template>
  <div style="display: flex;justify-content: space-between;margin-bottom: 20px;">
    <a-space>
      <a-input v-model:value="query.name" placeholder="名称" style="width: 250px; margin-right: 10px;">
        <template #suffix>
          <SearchOutlined style="color: rgba(0, 0, 0, 0.45)" @click="onSearch" />
        </template>
      </a-input>
      <a-button type="primary" @click="onDelete" :disabled="disabled" danger>删除</a-button>
    </a-space>
    <div>
      <a-space size="middle">
        <a-button type="primary" @click="onCreate">新建</a-button>
      </a-space>
    </div>
  </div>
<a-table rowKey="id"
         :columns="columns"
         :data-source="data.data.items" :scroll="{ y: '59vh' }" class="ant-table-striped" >
  <template #bodyCell="{ column, text, record }">

  </template>
</a-table>

  <a-modal v-model:visible="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
           width="800px" :centered="true">
    <div style="height: 55vh; overflow-y: scroll;padding: 0 15px;">
      <a-form ref="clusterFormRef" :model="cluster" layout="vertical" name="cluster" :rules="rules">
        <a-row :gutter="16">
          <a-col auto-size>
            <a-form-item label="名称" name="name">
              <a-input v-model:value="cluster.name"/>
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item label="endPoint" name="endPoint">
              <a-input v-model:value="cluster.controlPlaneEndpoint"/>
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item label="service网段" name="serviceCidr">
              <a-input v-model:value="cluster.serviceCidr"/>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="8">
            <a-form-item label="pod网段" name="podCidr">
              <a-input v-model:value="cluster.podCidr"/>
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item label="网络插件" name="networkPlug">
              <a-select
                  v-model:value="cluster.networkPlug"
                  show-search
                  style="width: 100%"
                  placeholder="请选择"
              >
                <a-select-option value="cilium">cilium</a-select-option>
                <!--                <a-select-option value="flannel">flannel</a-select-option>-->
                <!--                <a-select-option value="calico">calico</a-select-option>-->
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="16">
            <a-form-item label="version" name="版本">
              <a-input v-model:value="cluster.version"></a-input>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="16">
          <a-col :span="16">
            <a-form-item label="master node">
              <a-form-item
                  v-for="(host, index) in cluster.master"
                  :key="host"
                  v-bind="index === 0 ? formItemLayout : {}"
                  :label="index === 0 ? 'host' : ''">
                <a-form-item label="IP" name="IP">
                  <a-input
                      v-model:value="host.ip"
                      placeholder="please input host IP"
                      style="width: 30%; margin-right: 8px"
                  />
                </a-form-item>
                <MinusCircleOutlined
                    v-if="cluster.master.length > 1"
                    class="dynamic-delete-button"
                    :disabled="cluster.master.length === 1"
                    @click="removeMaster(host)"
                />
              </a-form-item>
              <a-form-item v-bind="formItemLayoutWithOutLabel">
                <a-button type="dashed" style="width: 60%" @click="addMaster">
                  <PlusOutlined />
                  Add field
                </a-button>
              </a-form-item>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="8">
            <a-form-item label="registry" name="registry">
              <a-select
                  v-model:value="cluster.registry"
                  show-search
                  style="width: 100%"
                  placeholder="请选择"
              >
                <a-select-option value="registry.aliyuncs.com/google_containers">阿里云</a-select-option>
                <a-select-option value="k8s.gcr.io">k8s.gcr.io</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </div>
  </a-modal>
</template>

<script >
import { SearchOutlined, ExclamationCircleOutlined, MinusCircleOutlined, PlusOutlined} from '@ant-design/icons-vue';
import {defineComponent, reactive, ref} from 'vue';
import { queryCluster } from '../api/cluster';
export default defineComponent({
  components:{
    SearchOutlined
  },
  methods:{
    clusterList(){
      let param = {
        pageNum: 1,
        pageSize: 10,
      }
      queryCluster(param).then((res) => {
        console.log(res.data)
        this.data = res.data
      })
    },
    onSearch(){},
    // 表单校验
    rules() {},
    onCreate() {
      this.title = '创建集群'
      this.operation = 1
      this.visible = true
    },
    onSave(){

    },
    onCancel(){},
    onDelete(){},
    removeMaster(item) {
      let index = this.cluster.master.indexOf(item);
      if (index !== -1) {
        this.cluster.master.splice(index, 1);
      }
    },
    addMaster(){
      this.cluster.master.push({
        ip: "",
        password: "",
        port: 22}
      );
    },
  },
  data() {
    return {
      data: {
        code: 0,
        data: {
          total: 0,
          items: [],
        },
        resMsg: "",
      },
      query: {
        name: "",
      },

      // 表格分页
      pagination: {
        current: 1,
        pageSize: 10,
        total: undefined,
      },
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        },
        {
          title: '名称',
          dataIndex: 'name',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        }, {
          title: 'endPoint',
          dataIndex: 'controlPlaneEndpoint',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        },{
          title: '网络插件',
          dataIndex: 'networkPlug',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        },{
          title: 'pod网段',
          dataIndex: 'podCidr',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        }, {
          title: 'service网段',
          dataIndex: 'serviceCidr',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        },{
          title: '版本',
          dataIndex: 'version',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        },{
          title: '状态',
          dataIndex: 'status',
          width: 100,
          fixed: 'left',
          ellipsis: true,
        }

      ],
      collapsed: false,
      selectedKeys: 1,
      host: {
        ip: "",
        password: "",
        port: 22,
      },
      cluster:{
        name: "",
        controlPlaneEndpoint:"",
        networkPlug:"cilium",
        podCidr:"10.244.0.0/16",
        registry:"registry.aliyuncs.com/google_containers",
        serviceCidr:"10.96.0.0/16",
        version:"1.22.15",
        master:[{
          ip: "1.2.2.2",
          password: "",
          port: 22,
        }],
        slaveNode:[{
          ip: "3.3.3.3",
          password: "",
          port: 22,
        }],
      },
      title: "",
      operation: 1,
      visible: false,
      disabled:false,
      clusterFormRef: undefined,
      formItemLayout: {
        labelCol: {
          xs: { span: 24 },
          sm: { span: 4 },
        },
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 20 },
        },
      },
      formItemLayoutWithOutLabel: {
        wrapperCol: {
          xs: { span: 24, offset: 0 },
          sm: { span: 20, offset: 4 },
        },
      },
    };
  },
  mounted(){
    this.clusterList()
  }

});


</script>
<style>
.dynamic-delete-button {
  cursor: pointer;
  position: relative;
  top: 4px;
  font-size: 24px;
  color: #999;
  transition: all 0.3s;
}
.dynamic-delete-button:hover {
  color: #777;
}
.dynamic-delete-button[disabled] {
  cursor: not-allowed;
  opacity: 0.5;
}
</style>