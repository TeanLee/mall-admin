import REST from './rest';
import AuthInterceptor from './auth.interceptor';

// const API_URL = process.env.VUE_APP_API_URL;

const API_URL = "127.0.0.1";

class APIService extends REST {
  constructor() {
    super(`${API_URL}`);
    this.useInterceptor(AuthInterceptor);
  }

  getEndPointURL() {
    return this.endPointURL;
  }

  create(url = API_URL, config) {
    return new REST(url, config);
  }
}

export default new APIService();