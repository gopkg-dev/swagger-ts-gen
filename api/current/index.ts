import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  LoginToken,
  Menu,
  UpdateCurrentAvatarBody,
  UpdateCurrentUser,
  UpdatePassword,
  User,
  UserPermission,
  UserSocial,
} from './model';

/**
 * 修改头像
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateCurrentAvatar(data?: UpdateCurrentAvatarBody) {
  const formData = new FormData();
  if (data) {
    for (const [key, value] of Object.entries(data)) {
      if (value === undefined || value === null) {
        continue;
      }
      if (Array.isArray(value)) {
        for (const item of value) {
          if (item !== undefined && item !== null) {
            formData.append(key, item as any);
          }
        }
      } else {
        formData.append(key, value as any);
      }
    }
  }
  const res = await request.put<ApiResult<void>>('/api/v1/current/avatar', formData, { headers: { 'Content-Type': 'multipart/form-data' } });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '修改头像失败'));
}


/**
 * 登出
 * @returns Promise<void>
 */
export async function logout() {
  const res = await request.post<ApiResult<void>>('/api/v1/current/logout');
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '登出失败'));
}


/**
 * 获取路由信息
 * @returns Promise<Menu[]>
 */
export async function queryCurrentMenus() {
  const res = await request.get<ApiResult<Menu[]>>('/api/v1/current/menus');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取路由信息失败'));
}


/**
 * 修改密码
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateCurrentPassword(data: UpdatePassword) {
  const res = await request.put<ApiResult<void>>('/api/v1/current/password', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '修改密码失败'));
}


/**
 * 查询权限数据
 * @returns Promise<UserPermission>
 */
export async function queryCurrentPermissions() {
  const res = await request.get<ApiResult<UserPermission>>('/api/v1/current/permissions');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询权限数据失败'));
}


/**
 * 刷新当前访问令牌
 * @returns Promise<LoginToken>
 */
export async function refreshToken() {
  const res = await request.post<ApiResult<LoginToken>>('/api/v1/current/refresh-token');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '刷新当前访问令牌失败'));
}


/**
 * 查询当前用户已绑定的第三方账号（需登录）
 * @returns Promise<UserSocial[]>
 */
export async function getCurrentSocialBindings() {
  const res = await request.get<ApiResult<UserSocial[]>>('/api/v1/current/social');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询当前用户已绑定的第三方账号（需登录）失败'));
}


/**
 * 获取用户信息
 * @returns Promise<User>
 */
export async function getUserInfo() {
  const res = await request.get<ApiResult<User>>('/api/v1/current/user');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取用户信息失败'));
}


/**
 * 修改基础信息
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateCurrentUser(data: UpdateCurrentUser) {
  const res = await request.put<ApiResult<void>>('/api/v1/current/user', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '修改基础信息失败'));
}

