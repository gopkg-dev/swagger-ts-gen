import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Menu,
  MenuForm,
  MenusQueryParam,
} from './model';

/**
 * 查询菜单列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Menu>>
 */
export async function queryMenus(params?: MenusQueryParam) {
  const res = await request.get<ApiResult<PageResult<Menu>>>('/api/v1/menus', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询菜单列表失败'));
}


/**
 * 创建菜单
 * @param data - 请求数据
 * @returns Promise<Menu>
 */
export async function createMenu(data: MenuForm) {
  const res = await request.post<ApiResult<Menu>>('/api/v1/menus', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建菜单失败'));
}


/**
 * 根据ID获取菜单
 * @param id - 唯一 ID
 * @returns Promise<Menu>
 */
export async function getMenu(id: string) {
  const res = await request.get<ApiResult<Menu>>(`/api/v1/menus/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取菜单失败'));
}


/**
 * 根据ID更新菜单
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateMenu(id: string, data: MenuForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/menus/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新菜单失败'));
}


/**
 * 根据ID删除菜单
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteMenu(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/menus/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除菜单失败'));
}

