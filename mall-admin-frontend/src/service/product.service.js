import api from "./api"

class ProductService {
  constructor() {
    this.api = api;
  }

  getProducts(page = 0, pageSize = 10) {
    return this.api.get('/product', {page, pageSize});
  }

  updateProduct(productId, product) {
    return this.api.post(`/product/${productId}`, product);
  }

  getAvailableServices() {
    return this.getServices()
      .then(list => list.filter(x => x.available));
  }

  getServicesByServicesetId(superserviceId) {
    return this.api.get(`/superservices/${superserviceId}/services`);
  }

  getServicesets() {
    return this.api.get('/superservices');
  }

  getServicesetsById(superserviceId) {
    return this.api.get(`/superservices/${superserviceId}`);
  }

  getAvailableServicesets() {
    return this.getServicesets()
      .then(list => list.filter(x => x.available));
  }

  /**
   * create a service
   * @param {Object} service
   *  service.name  {String} name of servcie;
   *  service.description {String} desc of service;
   *  service.service_type {String} type of service;
   */
  createService(service) {
    return this.api.post('/services', service);
  }

  /**
   * update service info
   * @param {Object} serviceId id of service;
   * @param {Object} service service
   */
  updateService(serviceId, service) {
    return this.api.patch(`/services/${serviceId}`, service);
  }

  /**
   * crate a servcieset
   * @param {Object} params
   *  params.name {String} name of servcieset;
   *  params.description {String} description of serviceset;
   */
  createServiceSet(params) {
    return this.api.post('/superservices', params);
  }

  updateServiceSet(superserviceId, params) {
    if (!params.logo_url) params.logo_url = '';
    if (!params.short_description) params.short_description = '';
    if (!params.description) params.description = '';
    return this.api.patch(`/superservices/${superserviceId}`, params);
  }

  deleteServiceSet(superserviceId) {
    return this.api.delete(`/superservices/${superserviceId}`);
  }

}

export default new ProductService()