<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider class="layout-side" width="150">
      <div class="logo">
        <div><img src="../assets/logo.png"
                  style="width: 32px;height: 32px;filter: drop-shadow(2px 2px 6px #79bbff);" /></div>
        <div v-if="collapsed == false" class="title"><b>Z</b>O<b style="color: #1283FF;">C</b>RM</div>
      </div>
      <a-menu style="border-right: none;width: 149px;" v-model:selectedKeys="selectedKeys" mode="inline">
        <a-menu-item :key="item.key" v-for="item in menuItem">
          <router-link :to="item.path" :key="item.path">
            <span>{{ item.name }}</span>
          </router-link>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout-content :style="{ padding: '0 24px', minHeight: '280px' }">
      <router-view></router-view>
    </a-layout-content>

  </a-layout>
</template>
<script lang="ts">
import {
  PieChartOutlined,
  DesktopOutlined,
  UserOutlined,
  TeamOutlined,
  FileOutlined,
} from '@ant-design/icons-vue';
import { reactive,defineComponent, ref } from 'vue';
import { DashboardOutlined } from '@ant-design/icons-vue';

export default defineComponent({
  components: {
    PieChartOutlined,
    DesktopOutlined,
    UserOutlined,
    TeamOutlined,
    FileOutlined,
  },

  data() {
    return {
      // 菜单选项
       menuItem: reactive([
         {
           key: "/",
           path: "/",
           icon: DashboardOutlined,
           name: "home"
         },{
        key: "cluster",
        path: "/cluster",
        icon: DashboardOutlined,
        name: "集群"
      }]),
      collapsed: ref<boolean>(false),
      selectedKeys: ref<string[]>(['1']),
    };
  },
});

</script>
<style>
#components-layout-demo-side .logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}
[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}

.layout-side {
  left: 0;
  top: 0;
  bottom: 0;
  overflow: auto;
  height: 100vh;
  position: fixed;
  background: #fff;
  border-right: 0.5px solid #F0F2F5;
}

</style>
