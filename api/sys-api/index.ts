import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Api,
  ApisQueryParam,
  DeleteApiForm,
  SyncResult,
} from './model';

/**
 * 查询接口列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Api>>
 */
export async function queryApis(params?: ApisQueryParam) {
  const res = await request.get<ApiResult<PageResult<Api>>>('/api/v1/sys-api', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询接口列表失败'));
}


/**
 * 批量删除接口
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function deleteApis(data: DeleteApiForm) {
  const res = await request.delete<ApiResult<void>>('/api/v1/sys-api', { data });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量删除接口失败'));
}


/**
 * 同步 Swagger 接口
 * @returns Promise<SyncResult>
 */
export async function syncApis() {
  const res = await request.post<ApiResult<SyncResult>>('/api/v1/sys-api/sync');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '同步 Swagger 接口失败'));
}


/**
 * 获取所有接口分组标签
 * @returns Promise<string[]>
 */
export async function getApiTags() {
  const res = await request.get<ApiResult<string[]>>('/api/v1/sys-api/tags');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取所有接口分组标签失败'));
}


/**
 * 根据ID获取接口
 * @param id - 唯一 ID
 * @returns Promise<Api>
 */
export async function getApi(id: string) {
  const res = await request.get<ApiResult<Api>>(`/api/v1/sys-api/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取接口失败'));
}

