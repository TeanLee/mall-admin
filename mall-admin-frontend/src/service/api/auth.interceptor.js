import Vue from 'vue';
import { get as getValue } from 'lodash';
import { Message } from "element-ui";

/**
 * auth service handle global authrization.
 */
export default {
  request(config) {
    return config;
  },

  // global ajax success handler
  response(res) {
    if (/^20\d/.test(res.status)) {
      return res.data;
    }
    return res;
  },

  // global ajax error handler
  responseError(error) {
    // error reponse
    const { response = {} } = error;
    switch (response.status) {
      case 502:
        Vue.noty.error('后端出问题了, 请联系管理员');
        break;
      case 401:
        Message.error("未授权，请先登录");
        if (Vue.$router) {
          Vue.$router.push({
            name: 'login',
          });
        }
        break;
      case 403:
        if (getValue(response, 'headers.Authorization')) {
          Vue.noty.error('权限不足');
        } else {
          Vue.noty.error(getValue(response, 'data.error_info'));
        }
        break;
      default:
        Vue.noty.error(getValue(response, 'data.error_info'));
    }
    return Promise.reject(response);
  },
};