import { GetSystem, SaveSystem } from '@/api/system'
const state = {
}
const mutations = {
}
const actions = {
  GetSystem({ commit }) {
    return new Promise((resolve, reject) => {
      GetSystem().then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  SaveSystem({ commit }, Data) {
    return new Promise((resolve, reject) => {
      SaveSystem(Data).then(response => {
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
