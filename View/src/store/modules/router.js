import { SettingRouter, GetRoleRouter } from '@/api/router'
import { getToken } from '@/utils/auth'

const state = {
  token: getToken()
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  }
}

const actions = {
  // 设置路由权限
  SettingRouter({ commit }, Info) {
    return new Promise((resolve, reject) => {
      SettingRouter(Info).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 获取角色权限路由
  GetRoleRouter({ commit }, Info) {
    return new Promise((resolve, reject) => {
      GetRoleRouter(Info).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

}
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
