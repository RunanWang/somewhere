import request from '@/utils/request'

export function getRecordList(params) {
  return request({
    url: '/somewhere/records',
    method: 'get',
    params
  })
}

export function postRecord(data) {
  return request({
    url: '/somewhere/records',
    method: 'post',
    data
  })
}
