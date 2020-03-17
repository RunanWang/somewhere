import request from '@/utils/request'

export function trainModel(data) {
  return request({
    url: '/somewhere/algo/train',
    method: 'post',
    data
  })
}