import api from "./api"

class OrderService {
  constructor() {
    this.api = api;
  }

  getOrders(param) {
    return this.api.get('/order', param);
  }
}

export default new OrderService()