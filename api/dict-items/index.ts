import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  DeleteDictItemForm,
  DictItem,
  DictItemForm,
  DictItemsQueryParam,
} from './model';

/**
 * 批量删除字典项
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function deleteDictItems(data: DeleteDictItemForm) {
  const res = await request.delete<ApiResult<void>>('/api/v1/dict-items', { data });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量删除字典项失败'));
}


/**
 * 查询字典项列表
 * @param params - 查询参数
 * @returns Promise<PageResult<DictItem>>
 */
export async function queryDictItems(params?: DictItemsQueryParam) {
  const res = await request.get<ApiResult<PageResult<DictItem>>>('/api/v1/dict-items', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询字典项列表失败'));
}


/**
 * 创建字典项
 * @param data - 请求数据
 * @returns Promise<DictItem>
 */
export async function createDictItem(data: DictItemForm) {
  const res = await request.post<ApiResult<DictItem>>('/api/v1/dict-items', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建字典项失败'));
}


/**
 * 根据ID获取字典项
 * @param id - 唯一 ID
 * @returns Promise<DictItem>
 */
export async function getDictItem(id: string) {
  const res = await request.get<ApiResult<DictItem>>(`/api/v1/dict-items/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取字典项失败'));
}


/**
 * 根据ID更新字典项
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateDictItem(id: string, data: DictItemForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/dict-items/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新字典项失败'));
}

