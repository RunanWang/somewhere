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

export function postUserList(data) {
  return request({
    url: '/somewhere/users',
    method: 'post',
    data
  })
}

export function deleteUserList(data) {
  return request({
    url: '/somewhere/users',
    method: 'delete',
    data
  })
}
