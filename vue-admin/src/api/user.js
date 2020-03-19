import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/api/v1/user/list',
    method: 'get',
    params
  })
}

export function createUser(data) {
  return request({
    url: '/api/v1/user',
    method: 'post',
    data
  })
}

export function updateUser(data) {
  return request({
    url: '/api/v1/user',
    method: 'put',
    data
  })
}

export function deleteUser(data) {
  return request({
    url: '/api/v1/user/' + data,
    method: 'delete',
    data
  })
}

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/user/info',
    method: 'get'
  })
}

export function refreshToken() {
  return request({
    url: '/user/refresh_token',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
