import { Login, GetList as GetAdminList, Add, Edit, Delete } from '@/api/user'
import { getToken, setToken, removeToken, setUserInfo, GetUserInfo } from '@/utils/auth'
import router, { resetRouter, constantRoutes } from '@/router'
import { GetMenuList } from '@/api/menus'
import Layout from '@/layout'

export function handleRouterMenu(menu) {
  const ResData = [];
  for (const i in menu) {
    if (menu[i].type != 3) {
      const item = {};
      if (menu[i].type == 1) {
        item.component = Layout
        item.alwaysShow = true
        item.redirect = ''
        item.name = menu[i].router//因为是目录，没有组件名称，就给他路由名称
        item.path = "/" + menu[i].router;
      }
      if (menu[i].type == 2) {
        item.path = menu[i].router;
        item.component = (resolve) => require([`@/views${menu[i].path}`], resolve);
        item.name = menu[i].component;
      }
      if (menu[i].show == 0) {
        item.hidden = true
      }
      item.meta = {}
      item.meta.title = menu[i].title
      item.meta.icon = menu[i].icon
      if (menu[i].children.length > 0) {
        item.children = handleRouterMenu(menu[i].children)
      }
      ResData.push(item);
    }
  }
  return ResData;
}
const state = {
  token: getToken(),
  name: '',
  avatar: '',
  routes: [],
  menus: [],
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROUTES: (state, routes) => {
    state.routes = routes
  },
  SET_MEMUS: (state, menus) => {
    state.menus = menus
  }
}
const actions = {
  Login({ commit }, userInfo) {
    return new Promise((resolve, reject) => {
      Login(userInfo).then(response => {
        const { data } = response
        setUserInfo(data)
        commit('SET_TOKEN', data.token)
        commit('SET_NAME', data.account)
        commit('SET_AVATAR', 'https://www.bugquit.com/wp-content/uploads/2019/12/logo-1.png')
        setToken(data.token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  async GetMenuList({ commit }) {
    return new Promise(resolve => {
      GetMenuList().then(response => {

        let MenusList = handleRouterMenu(response.data)
        MenusList.push({ path: '*', redirect: '/404', hidden: true })
        MenusList = constantRoutes.concat(MenusList)
        let UserInfo = GetUserInfo()
        commit('SET_NAME', UserInfo.account)
        commit('SET_AVATAR', 'https://www.bugquit.com/wp-content/uploads/2019/12/logo-1.png')
        commit('SET_ROUTES', MenusList)
        commit('SET_MEMUS', response.data)
        resolve(MenusList)
      }).catch(error => {
        console.error(error);
      })
    })
  },
  logout({ commit, state, dispatch }) {
    return new Promise((resolve) => {
      commit('SET_TOKEN', '')
      commit('SET_ROUTES', [])
      removeToken()
      resetRouter()
      dispatch('tagsView/delAllViews', null, { root: true })
      resolve()
    })
  },
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },
  GetAdminList({ commit }, SearchData) {
    return new Promise((resolve, reject) => {
      GetAdminList(SearchData).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Add({ commit }, userInfo) {
    return new Promise((resolve, reject) => {
      Add(userInfo).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  Edit({ commit }, userInfo) {
    return new Promise((resolve, reject) => {
      Edit(userInfo).then(response => {
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
  async changeRoles({ commit, dispatch }) {
    const accessRoutes = await dispatch('GetMenuList')
    resetRouter()
    router.addRoutes(accessRoutes)
    dispatch('tagsView/delAllViews', null, { root: true })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
