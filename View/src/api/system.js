import request from '@/utils/request'
import Qs from 'qs'

export function GetSystem() {
    return request({
        url: '/v1/system/get',
        method: 'GET',
    })
}
export function SaveSystem(data) {
    return request({
        url: '/v1/system/save',
        method: 'post',
        data: Qs.stringify(data)
    })
}