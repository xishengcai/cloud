import request from "../router/request.js";
import {BasicPageParams} from "./model/baseModel";

// 查询集群列表
export function queryCluster(param:BasicPageParams) {
    return request({
        url: '/cluster',
        method: 'get',
        params: param,
    })
}

export function createCluster(param){
    return request({
        url:"/cluster",
        method:"post",
        data: param,
    })
}