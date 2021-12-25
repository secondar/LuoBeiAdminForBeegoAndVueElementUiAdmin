import request from '@/utils/request'
import Qs from 'qs'

export function AddSort(data) {
  return request({
    url: '/v1/article/sort/add',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function EditSort(data) {
  return request({
    url: '/v1/article/sort/edit',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function DeleteSort(data) {
  return request({
    url: '/v1/article/sort/delete',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function GetSort(data) {
  return request({
    url: '/v1/article/sort/getlist',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Add(data) {
  return request({
    url: '/v1/article/add',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Edit(data) {
  return request({
    url: '/v1/article/edit',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Delete(data) {
  return request({
    url: '/v1/article/delete',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function GetList(data) {
  return request({
    url: '/v1/article/getlist',
    method: 'post',
    data: Qs.stringify(data)
  })
}
export function Details(ID) {
  return request({
    url: '/v1/article/details?id=' + ID,
    method: 'GET',
  })
}