import axios from 'axios';
import { message } from 'ant-design-vue';

axios.defaults.baseURL = "http://localhost:80/api/v1"

const request = axios.create({
    // timeout: 5000,`
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})

export default request;