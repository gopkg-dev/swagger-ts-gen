import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Role,
  RoleDataScope,
  RoleDataScopeForm,
  RoleForm,
  RoleMenuData,
  RoleMenuDataForm,
  RoleUsersForm,
  RolesQueryParam,
} from './model';

/**
 * 查询角色列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Role>>
 */
export async function queryRoles(params?: RolesQueryParam) {
  const res = await request.get<ApiResult<PageResult<Role>>>('/api/v1/roles', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询角色列表失败'));
}


/**
 * 创建角色
 * @param data - 请求数据
 * @returns Promise<Role>
 */
export async function createRole(data: RoleForm) {
  const res = await request.post<ApiResult<Role>>('/api/v1/roles', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建角色失败'));
}


/**
 * 批量添加用户到角色
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function addRoleUsers(data: RoleUsersForm) {
  const res = await request.post<ApiResult<void>>('/api/v1/roles/users', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量添加用户到角色失败'));
}


/**
 * 批量移除角色下的用户
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function deleteRoleUsers(data: RoleUsersForm) {
  const res = await request.delete<ApiResult<void>>('/api/v1/roles/users', { data });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量移除角色下的用户失败'));
}


/**
 * 根据ID获取角色
 * @param id - 唯一 ID
 * @returns Promise<Role>
 */
export async function getRole(id: string) {
  const res = await request.get<ApiResult<Role>>(`/api/v1/roles/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取角色失败'));
}


/**
 * 根据ID更新角色
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateRole(id: string, data: RoleForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/roles/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新角色失败'));
}


/**
 * 根据ID删除角色
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteRole(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/roles/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除角色失败'));
}


/**
 * 获取角色的数据权限
 * @param id - 唯一 ID
 * @returns Promise<RoleDataScope>
 */
export async function getRoleDataScope(id: string) {
  const res = await request.get<ApiResult<RoleDataScope>>(`/api/v1/roles/${id}/data`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取角色的数据权限失败'));
}


/**
 * 保存角色的数据权限
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function saveRoleDataScope(id: string, data: RoleDataScopeForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/roles/${id}/data`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '保存角色的数据权限失败'));
}


/**
 * 获取角色的菜单权限数据
 * @param id - 唯一 ID
 * @returns Promise<RoleMenuData>
 */
export async function getRoleMenuData(id: string) {
  const res = await request.get<ApiResult<RoleMenuData>>(`/api/v1/roles/${id}/menus`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取角色的菜单权限数据失败'));
}


/**
 * 保存角色的菜单权限数据
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function saveRoleMenuData(id: string, data: RoleMenuDataForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/roles/${id}/menus`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '保存角色的菜单权限数据失败'));
}

