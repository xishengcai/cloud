import axios from 'axios';
import { message } from 'ant-design-vue';

axios.defaults.baseURL = "http://localhost:80/api/v1"

const request = axios.create({
    // timeout: 5000,`
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})

request.interceptors.response.use(response => {
    console.log(response)
    if (response.data.code != 0) {
        message.error(response.data.resMsg)
    }
    return response;
}, error => {
    console.log(error)
    return Promise.reject(error)
})

export default request;