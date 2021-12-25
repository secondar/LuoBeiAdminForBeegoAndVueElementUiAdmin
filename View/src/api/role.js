import request from '@/utils/request'
import Qs from 'qs'
export function Add(data) {
  return request({
    url: '/v1/role/add',
    method: 'POST',
    data: Qs.stringify(data)
  })
}
export function Edit(data) {
  return request({
    url: '/v1/role/edit',
    method: 'POST',
    data: Qs.stringify(data)
  })
}
export function GetList() {
  return request({
    url: '/v1/role/list',
    method: 'GET'
  })
}
export function Delete(data) {
  return request({
    url: '/v1/role/delete',
    method: 'POST',
    data: Qs.stringify(data)
  })
}