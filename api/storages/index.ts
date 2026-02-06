import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Storage,
  StorageForm,
  StoragesQueryParam,
} from './model';

/**
 * 查询存储列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Storage>>
 */
export async function queryStorages(params?: StoragesQueryParam) {
  const res = await request.get<ApiResult<PageResult<Storage>>>('/api/v1/storages', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询存储列表失败'));
}


/**
 * 创建存储
 * @param data - 请求数据
 * @returns Promise<Storage>
 */
export async function createStorage(data: StorageForm) {
  const res = await request.post<ApiResult<Storage>>('/api/v1/storages', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建存储失败'));
}


/**
 * 根据ID获取存储
 * @param id - 唯一 ID
 * @returns Promise<Storage>
 */
export async function getStorage(id: string) {
  const res = await request.get<ApiResult<Storage>>(`/api/v1/storages/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取存储失败'));
}


/**
 * 根据ID更新存储配置
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateStorage(id: string, data: StorageForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/storages/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新存储配置失败'));
}


/**
 * 根据ID删除存储
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteStorage(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/storages/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除存储失败'));
}

