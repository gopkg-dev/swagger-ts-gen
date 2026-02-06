import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  Option,
  OptionsQueryParam,
  RestOptionValueForm,
  UpdateOptionsForm,
} from './model';

/**
 * 批量更新参数
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateOptions(data: UpdateOptionsForm) {
  const res = await request.put<ApiResult<void>>('/api/v1/options', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量更新参数失败'));
}


/**
 * 查询参数列表
 * @param params - 查询参数
 * @returns Promise<Option[]>
 */
export async function queryOptions(params?: OptionsQueryParam) {
  const res = await request.get<ApiResult<Option[]>>('/api/v1/options', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询参数列表失败'));
}


/**
 * 重置参数为默认值
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function resetOptions(data: RestOptionValueForm) {
  const res = await request.post<ApiResult<void>>('/api/v1/options/reset', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '重置参数为默认值失败'));
}

