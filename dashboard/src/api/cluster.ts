import request from "../router/request.js";

// 查询集群列表
export function queryCluster(param) {
    return request({
        url: '/cluster',
        method: 'get',
        params: param,
    })
}