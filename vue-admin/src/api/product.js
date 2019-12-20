import request from '@/utils/request'

export function getProductList(params) {
  return request({
    url: '/somewhere/products',
    method: 'get',
    params
  })
}

export function getProductListByPage(data) {
  return request({
    url: '/somewhere/productspage',
    method: 'post',
    data
  })
}

export function postProduct(data) {
  return request({
    url: '/somewhere/products',
    method: 'post',
    data
  })
}

export function deleteProduct(data) {
  return request({
    url: '/somewhere/products',
    method: 'delete',
    data
  })
}

export function putProduct(data) {
  return request({
    url: '/somewhere/products',
    method: 'put',
    data
  })
}
