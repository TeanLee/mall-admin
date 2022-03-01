import Vue from 'vue'
import Vuex from 'vuex'
import { setUserName, removeRole, removeUserName  } from '@/utils/storage'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    username: '',
    role: ''
  },
  getters: {
    username(state) {
      if(window.localStorage.getItem('username') != null) {
        return window.localStorage.getItem('username')
      }
      return state.username
    }
  },
  mutations: {
    saveUserName(state, username) {
      state.username = username
    },
    deleteUserInfo(state) {
      state.username = ''
    }
  },
  actions: {
    saveUserName(context, username) {
      setUserName(username)
      context.commit('saveUserName', username)
    },
    deleteUserInfo(context) {
      removeUserName();
      removeRole()
      context.commit('deleteUserInfo')
    }
  },
  modules: {
  }
})
