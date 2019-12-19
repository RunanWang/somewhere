import request from '@/utils/request'

export function getStoreList(params) {
  return request({
    url: '/somewhere/stores',
    method: 'get',
    params
  })
}

export function getStoreListByPage(data) {
  return request({
    url: '/somewhere/storespage',
    method: 'post',
    data
  })
}

export function postStoreList(data) {
  return request({
    url: '/somewhere/stores',
    method: 'post',
    data
  })
}

export function deleteStoreList(data) {
  return request({
    url: '/somewhere/stores',
    method: 'delete',
    data
  })
}

export function putStoreList(data) {
  return request({
    url: '/somewhere/stores',
    method: 'put',
    data
  })
}
