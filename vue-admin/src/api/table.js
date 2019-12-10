import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/somewhere/stores',
    method: 'get',
    params
  })
}
