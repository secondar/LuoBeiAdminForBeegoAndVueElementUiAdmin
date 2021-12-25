import request from '@/utils/request'
import Qs from 'qs'
export function Add(data) {
  return request({
    url: '/v1/menu/add',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Edit(data) {
  return request({
    url: '/v1/menu/edit',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Delete(data) {
  return request({
    url: '/v1/menu/delete',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function GetList() {
  return request({
    url: '/v1/menu/list',
    method: 'get',
  })
}
export function GetMenuList() {
  return request({
    url: '/v1/menu/getadminmenurouter',
    method: 'get',
  })
}