import request from '@/utils/request'
import Qs from 'qs'
export function Login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function GetList(data) {
  return request({
    url: '/v1/admin/getlist',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Add(data) {
  return request({
    url: '/v1/admin/add',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Edit(data) {
  return request({
    url: '/v1/admin/edit',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Delete(data) {
  return request({
    url: '/v1/admin/delete',
    method: 'post',
    data: Qs.stringify(data)
  })
}