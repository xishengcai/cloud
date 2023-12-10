import Home from '../components/Home.vue'
import Cluster from '../components/Cluster.vue'
import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
    {
        path: "/",
        component: Home,
        children: [
            {
                path: "home",
                name:"home",
                component: Home,
            },
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