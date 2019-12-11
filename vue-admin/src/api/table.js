import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/somewhere/stores',
    method: 'get',
    params
  })
}

export function getUserList(params) {
  return request({
    url: '/somewhere/users',
    method: 'get',
    params
  })
}
