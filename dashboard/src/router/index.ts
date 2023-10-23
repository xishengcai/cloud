import Home from '../components/Home.vue'
import Cluster from '../components/Cluster.vue'
import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
    {
        path: "/",
        component: Home,
        redirect: "/cluster",
        children: [
            {
                path: "cluster",
                name:"cluster",
                component: Cluster,
            }
        ],
    }
];

const router = createRouter({
    history: createWebHashHistory(), routes
})

export default router