import { Add, GetList, Edit, Delete } from '@/api/menus'
import { getToken, setToken, removeToken } from '@/utils/auth'

const state = {
  token: getToken(),
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
}

const actions = {
  // 添加
  Add({ commit }, Info) {
    return new Promise((resolve, reject) => {
      Add(Info).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 编辑
  Edit({ commit }, Info) {
    return new Promise((resolve, reject) => {
      Edit(Info).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 删除
  Delete({ commit }, Id) {
    return new Promise((resolve, reject) => {
      Delete({ id: Id }).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // 获取列表
  GetList({ commit }) {
    return new Promise((resolve, reject) => {
      GetList().then(response => {
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
