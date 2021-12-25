import { GetSort, AddSort, EditSort, DeleteSort, Add, GetList, Details, Edit, Delete } from '@/api/article'
const state = {
}
const mutations = {
}
const actions = {
  GetSort({ commit }, SearchData) {
    return new Promise((resolve, reject) => {
      GetSort(SearchData).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  AddSort({ commit }, Data) {
    return new Promise((resolve, reject) => {
      AddSort(Data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  EditSort({ commit }, Data) {
    return new Promise((resolve, reject) => {
      EditSort(Data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  DeleteSort({ commit }, id) {
    return new Promise((resolve, reject) => {
      DeleteSort({ id: id }).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Add({ commit }, Data) {
    return new Promise((resolve, reject) => {
      Add(Data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Edit({ commit }, Data) {
    return new Promise((resolve, reject) => {
      Edit(Data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Delete({ commit }, id) {
    return new Promise((resolve, reject) => {
      Delete({ id: id }).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  GetList({ commit }, Data) {
    return new Promise((resolve, reject) => {
      GetList(Data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Details({ commit }, id) {
    return new Promise((resolve, reject) => {
      Details(id).then(response => {
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
