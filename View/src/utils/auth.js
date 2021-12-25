import Cookies from 'js-cookie'

const TokenKey = 'token'
const UserInfoKey = 'userinfo'

export function getToken() {
  return Cookies.get(TokenKey)
}
export function setToken(token) {
  return Cookies.set(TokenKey, token)
}
export function setUserInfo(info) {
  return Cookies.set(UserInfoKey, info)
}
export function GetUserInfo(info) {
  return Cookies.get(UserInfoKey)
}
export function removeToken() {
  Cookies.remove(UserInfoKey)
  return Cookies.remove(TokenKey)
}
