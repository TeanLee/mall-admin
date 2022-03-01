import api from "./api"

class AuthService {
  constructor() {
    this.api = api;
  }

  login(param) {
    return this.api.post(`/login`, param);
  }

  logout() {
    return this.api.post(`/logout`);
  }
}

export default new AuthService()