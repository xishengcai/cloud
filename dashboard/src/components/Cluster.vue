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

  <a-table :dataSource="resp.data.items" :columns="columns" />

  <a-modal v-model:open="visible" :title="title" @ok="onSave" @cancel="onCancel" cancelText="取消" okText="保存"
           width="800px" :centered="true">
    <div style="height: 55vh; overflow-y: scroll; padding: 0 15px;">
      <a-form ref="clusterFormRef" :model="cluster" layout="vertical" name="cluster" :rules="rules">
        <a-row :gutter="20">
          <a-col auto-size>
            <a-form-item label="名称" name="name">
              <a-input v-model:value="cluster.name"/>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="20">
          <a-col auto-size>
            <a-form-item label="endPoint" name="endPoint">
              <a-input v-model:value="cluster.controlPlaneEndpoint"/>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="20">
          <a-col auto-size>
            <a-form-item label="service网段" name="serviceCidr">
              <a-input v-model:value="cluster.serviceCidr"/>
            </a-form-item>
          </a-col>

          <a-col auto-size>
            <a-form-item label="pod网段" name="podCidr">
              <a-input v-model:value="cluster.podCidr"/>
            </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="20">

          <a-col auto-size>
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

        <a-row :gutter="20">
          <a-col auto-size>
            <a-form-item label="version" name="版本">
              <a-input v-model:value="cluster.version"></a-input>
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="80">
          <a-col auto-size>
            <a-form-item label="添加控制节点" >
              <a-form-item>

              </a-form-item>
              <a-space
                  v-for="(host, index) in cluster.master"
                  :key="host.id"
                    style="display: flex; margin-bottom: 8px"
                    align="baseline">
                <!-- host ip -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing IP',
                }">
                  <a-input
                      v-model:value="host.ip" placeholder="IP" style="width: 100%; "
                  />
                </a-form-item>
                <!-- host port -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing Port',
                }">
                  <a-input
                      v-model:value="host.port" placeholder="port" value=22 style="width: 100%; "
                  />
                </a-form-item>
                <!-- password -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing Password',
                }">
                  <a-input
                      v-model:value="host.password" placeholder="password" style="width: 100%; "
                  />
                </a-form-item>
                <!-- 移除host -->
                <MinusCircleOutlined
                    v-if="cluster.master.length > 1"
                    class="dynamic-delete-button"
                    :disabled="cluster.master.length === 1"
                    @click="removeMaster(host)"
                />
              </a-space>


              </a-form-item>
                <a-form-item v-bind="formItemLayoutWithOutLabel">
                <a-button type="dashed" style="width: 60%"  @click="addMaster">
                  <PlusOutlined />
                  Add Controller Node
                </a-button>
              </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="80">
          <a-col auto-size>
            <a-form-item label="添加工作节点" >
              <a-form-item>

              </a-form-item>
              <a-space
                  v-for="(host, index) in cluster.workNodes"
                  :key="host.id"
                    style="display: flex; margin-bottom: 8px"
                    align="baseline">
                <!-- host ip -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing IP',
                }">
                  <a-input
                      v-model:value="host.ip" placeholder="IP" style="width: 100%; "
                  />
                </a-form-item>
                <!-- host port -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing Port',
                }">
                  <a-input
                      v-model:value="host.port" placeholder="port" value=22 style="width: 100%; "
                  />
                </a-form-item>
                <!-- password -->
                <a-form-item
                :rules="{
                  required: true,
                  message: 'Missing Password',
                }">
                  <a-input
                      v-model:value="host.password" placeholder="password" style="width: 100%; "
                  />
                </a-form-item>
                <!-- 移除host -->
                <MinusCircleOutlined
                    class="dynamic-delete-button"
                    :disabled="cluster.master.length === 1"
                    @click="removeSlave(host)"
                />

              </a-space>


              </a-form-item>
              <a-form-item v-bind="formItemLayoutWithOutLabel">
                <a-button type="dashed" style="width: 60%" @click="addSlave">
                  <PlusOutlined />
                  Add Work Node
                </a-button>
              </a-form-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col auto-size>
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

import { SearchOutlined, MinusCircleOutlined, PlusOutlined} from '@ant-design/icons-vue';
import {defineComponent} from 'vue';
import { queryCluster,createCluster } from '../api/cluster';
import { message } from 'ant-design-vue';
export default defineComponent({
  components:{
    SearchOutlined,
    MinusCircleOutlined,
    PlusOutlined
  },
  methods:{
    clusterList(){
      let param = {
        page: 1,
        pageSize: 10
      }
      queryCluster(param).then((res) => {
        this.resp = res.data
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
      let param = {
          name: this.cluster.name,
          controlPlaneEndpoint: this.cluster.controlPlaneEndpoint,
          master: this.cluster.master,
          networkPlug: this.cluster.networkPlug,
          podCidr: this.cluster.podCidr,
          registry: this.cluster.registry,
          serviceCidr: this.cluster.serviceCidr,
          version: this.cluster.version,
          workNodess: this.cluster.workNodess
      }
      createCluster(param).then((res) => {
        if (res.data.code == 201) {
            message.success('保存成功')
            this.data.defaultSelectedIds = []
            this.clusterList()
        }
      })
      // this.clusterFormRef.resetFields()
      this.visible=false
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
    removeSlave(item) {
      let index = this.cluster.workNodes.indexOf(item);
      if (index !== -1) {
        this.cluster.workNodes.splice(index, 1);
      }
    },
    addSlave(){
      this.cluster.workNodes.push({
        ip: "",
        password: "",
        port: 22}
      );
    },
    // 已选中的ID
    onSelectedIds(selectedRowKeys){
      this.data.clusterIDs = selectedRowKeys
      if (this.data.clusterIDs.length !== 0) {
        this.data.disabled.value = false
      } else {
        this.data.disabled.value = true
      }
    }
  },
  data() {
    return {
      resp: {
        code: 0,
        data: {
          total: 0,
          items: [],
        },
        message: "",
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
      cluster:{
        name: "",
        controlPlaneEndpoint:"",
        networkPlug:"cilium",
        podCidr:"10.244.0.0/16",
        registry:"registry.aliyuncs.com/google_containers",
        serviceCidr:"10.96.0.0/16",
        version:"1.23.5",
        master:[{
          ip: "",
          password: "test@123",
          port: 22,
        }],
        workNodes:[
          {
            ip: "",
            password: "test@123",
            port: 22,
          }
        ],
      },
      title: "",
      operation: 1,
      visible: false,
      disabled: false,
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
      defaultSelectedIds: [],
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