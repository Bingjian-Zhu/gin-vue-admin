import { login, logout, getInfo, refreshToken } from '@/api/user'
import { getToken, setToken, removeToken, getTokenExpire, setTokenExpire, removeTokenExpire } from '@/utils/auth'
import { resetRouter } from '@/router'

const state = {
  token: getToken(),
  name: '',
  avatar: '',
  roles: [],
  tokenExpire: getTokenExpire()
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
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_TOKENEXPIRE: (state, token) => {
    state.tokenExpire = token
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  refreshToken({ commit }) {
    return new Promise((resolve, reject) => {
      refreshToken().then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit }) {
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
        const data = response.data
        if (!data) {
          reject('Verification failed, please Login again.')
        }
        const { Name, Avatar, Roles } = data

        // roles must be a non-empty array
        if (!Roles || Roles.length <= 0) {
          reject('getInfo: roles must be a non-null array!')
        }

        commit('SET_ROLES', Roles)
        commit('SET_NAME', Name)
        commit('SET_AVATAR', Avatar)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        commit('SET_TOKEN', '')
        commit('SET_TOKENEXPIRE', '')
        commit('SET_ROLES', '')
        commit('SET_NAME', '')
        commit('SET_AVATAR', [])
        removeToken()
        removeTokenExpire()
        resetRouter()
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_TOKENEXPIRE', '')
      removeToken()
      removeTokenExpire()
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

