import api from "./api"

class AuthService {
  constructor() {
    this.api = api;
  }

  login(param) {
    return this.api.post(`/login`, param);
  }
}

export default new AuthService()