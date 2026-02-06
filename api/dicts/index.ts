import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Dict,
  DictForm,
  DictsQueryParam,
} from './model';

/**
 * 查询字典列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Dict>>
 */
export async function queryDicts(params?: DictsQueryParam) {
  const res = await request.get<ApiResult<PageResult<Dict>>>('/api/v1/dicts', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询字典列表失败'));
}


/**
 * 创建字典
 * @param data - 请求数据
 * @returns Promise<Dict>
 */
export async function createDict(data: DictForm) {
  const res = await request.post<ApiResult<Dict>>('/api/v1/dicts', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建字典失败'));
}


/**
 * 根据ID删除字典
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteDict(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/dicts/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除字典失败'));
}


/**
 * 根据ID获取字典
 * @param id - 唯一 ID
 * @returns Promise<Dict>
 */
export async function getDict(id: string) {
  const res = await request.get<ApiResult<Dict>>(`/api/v1/dicts/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取字典失败'));
}


/**
 * 根据ID更新字典
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateDict(id: string, data: DictForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/dicts/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新字典失败'));
}

