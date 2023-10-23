import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        host: '0.0.0.0',
        port: 81,
        open: false, //自动打开
        // proxy: { // 本地开发环境通过代理实现跨域，生产环境使用 nginx 转发
        //     // 正则表达式写法
        //     '^/api': {
        //         target: 'http://localhost:80', // 后端服务实际地址
        //         changeOrigin: true, //开启代理
        //         rewrite: (path) => path.replace(/^\/api/, '')
        //     }
        // }
    },
    plugins: [vue()],
    css: {
        preprocessorOptions: {
            less: {
                modifyVars: {
                    'primary-color': '#476FFF',
                    'link-color': '#476FFF',
                    'border-radius-base': '5px'
                },
                javascriptEnabled: true
            }
        }
    },
})
