import request from '@/utils/request'
import Qs from 'qs'
export function SettingRouter(data) {
  return request({
    url: '/v1/router/setting',
    method: 'POST',
    data: Qs.stringify(data)
  })
}
export function GetRoleRouter(data) {
  return request({
    url: '/v1/router/getrouter',
    method: 'POST',
    data: Qs.stringify(data)
  })
}