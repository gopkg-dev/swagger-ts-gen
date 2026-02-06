import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  GenConfig,
  GenConfigUpdateForm,
  GenConfigsQueryParam,
  GenFieldConfig,
  GenFieldConfigsQueryParam,
  GenerateCodeResult,
} from './model';

/**
 * 查询生成配置列表
 * @param params - 查询参数
 * @returns Promise<PageResult<GenConfig>>
 */
export async function queryGenConfigs(params?: GenConfigsQueryParam) {
  const res = await request.get<ApiResult<PageResult<GenConfig>>>('/api/v1/codegen/configs', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询生成配置列表失败'));
}


/**
 * 更新生成配置
 * @param tableName - 表名
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateGenConfig(tableName: string, data: GenConfigUpdateForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/codegen/configs/${tableName}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '更新生成配置失败'));
}


/**
 * 获取生成配置详情
 * @param tableName - 表名
 * @returns Promise<GenConfig>
 */
export async function getGenConfig(tableName: string) {
  const res = await request.get<ApiResult<GenConfig>>(`/api/v1/codegen/configs/${tableName}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取生成配置详情失败'));
}


/**
 * 查询字段配置列表
 * @param params - 查询参数
 * @returns Promise<PageResult<GenFieldConfig>>
 */
export async function queryGenFieldConfigs(params?: GenFieldConfigsQueryParam) {
  const res = await request.get<ApiResult<PageResult<GenFieldConfig>>>('/api/v1/codegen/fields', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询字段配置列表失败'));
}


/**
 * 生成代码
 * @param tableName - 表名
 * @returns Promise<GenerateCodeResult>
 */
export async function generateCode(tableName: string) {
  const res = await request.post<ApiResult<GenerateCodeResult>>(`/api/v1/codegen/tables/${tableName}/generate`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '生成代码失败'));
}

