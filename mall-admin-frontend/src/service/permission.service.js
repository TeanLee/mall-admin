import api from "./api"

class PermissionService {
  constructor() {
    this.api = api;
  }

  getPermissionList() {
    return this.api.get(`/permission`);
  }

  getRoles() {
    return this.api.get(`/permission/roles`);
  }

  updateRole(adminId, params) {
    return this.api.post(`/permission/update-role/${adminId}`, params)
  }

  deleteAdmin(adminId) {
    return this.api.delete(`/permission/${adminId}`)
  }
}

export default new PermissionService()