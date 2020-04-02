import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/user/table/list',
    method: 'get',
    params
  })
}
