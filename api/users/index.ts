import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  PassWordResult,
  UpdatePassWordForm,
  User,
  UserForm,
  UsersQueryParam,
} from './model';

/**
 * 查询用户列表
 * @param params - 查询参数
 * @returns Promise<PageResult<User>>
 */
export async function queryUsers(params?: UsersQueryParam) {
  const res = await request.get<ApiResult<PageResult<User>>>('/api/v1/users', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询用户列表失败'));
}


/**
 * 创建用户
 * @param data - 请求数据
 * @returns Promise<User>
 */
export async function createUser(data: UserForm) {
  const res = await request.post<ApiResult<User>>('/api/v1/users', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建用户失败'));
}


/**
 * 更新用户信息
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateUser(id: string, data: UserForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/users/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '更新用户信息失败'));
}


/**
 * 根据ID删除用户
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteUser(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/users/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除用户失败'));
}


/**
 * 根据ID获取用户
 * @param id - 唯一 ID
 * @returns Promise<User>
 */
export async function getUser(id: string) {
  const res = await request.get<ApiResult<User>>(`/api/v1/users/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取用户失败'));
}


/**
 * 禁用账号
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function disableUser(id: string) {
  const res = await request.patch<ApiResult<void>>(`/api/v1/users/${id}/disable`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '禁用账号失败'));
}


/**
 * 启用账号
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function enableUser(id: string) {
  const res = await request.patch<ApiResult<void>>(`/api/v1/users/${id}/enable`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '启用账号失败'));
}


/**
 * 修改用户密码
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateUserPassWord(id: string, data: UpdatePassWordForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/users/${id}/password`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '修改用户密码失败'));
}


/**
 * 重置用户密码
 * @param id - 唯一 ID
 * @returns Promise<PassWordResult>
 */
export async function resetUserPassWord(id: string) {
  const res = await request.patch<ApiResult<PassWordResult>>(`/api/v1/users/${id}/reset-pwd`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '重置用户密码失败'));
}

