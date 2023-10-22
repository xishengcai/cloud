<template>
<a-table rowKey="id"
         :columns="columns"
         :data-source="data.clusters" :scroll="{ y: '59vh' }" class="ant-table-striped" >
  <template #bodyCell="{ column, text, record }">

  </template>

</a-table>
</template>

<script >
import { defineComponent, ref , reactive } from 'vue';
import { queryCluster } from '../api/cluster';
export default defineComponent({
  components: {
    props:{
      pagination: reactive({
        current: 1,
        pageSize: 10,
        total: undefined,
      }),
    },
  },
  methods:{
    clusterList: ()=>{
      let param = {
        pageNum: 1,
        pageSize: 10,
      }
      queryCluster(param).then((res) => {
        if (res.data.code == 0) {
          this.data.clusters = res.data.data.list
        }
      })
    }
  },
  data() {
    return {
      data : reactive({
        clusters: [],
      }),
      // 表格分页
      pagination: reactive({
        current: 1,
        pageSize: 10,
        total: undefined,
      }),
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
      collapsed: ref(false),
      selectedKeys: ref(['1']),
    };
  },
  mounted: function (){
    this.clusterList()
  }
  // onMount: function (){
  //   this.clusterList()
  // }
});


</script>

<style>

</style>