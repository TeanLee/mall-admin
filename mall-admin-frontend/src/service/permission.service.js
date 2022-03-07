import api from "./api"

class PermissionService {
  constructor() {
    this.api = api;
  }

  getPermissionList() {
    return this.api.get(`/permission`);
  }

  GetRoles() {
    return this.api.get(`/permission/roles`);
  }
}

export default new PermissionService()