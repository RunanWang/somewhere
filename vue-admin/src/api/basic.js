import request from '@/utils/request'

export function getBasic(params) {
  return request({
    url: '/basic',
    method: 'get',
    params
  })
}
