import api from "./api"

class UserService {
  constructor() {
    this.api = api;
  }

  getUsers() {
    return this.api.get('/user');
  }

  updateCategory(categoryId, category) {
    return this.api.post(`/category/${categoryId}`, category);
  }

  addUser(params) {
    return this.api.post(`/user/add`, params);
  }
}

export default new UserService()