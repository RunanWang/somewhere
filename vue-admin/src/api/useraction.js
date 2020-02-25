import request from '@/utils/request'

export function getRecommendListByPage(data) {
  return request({
    url: '/somewhere/recommend',
    method: 'post',
    data
  })
}
