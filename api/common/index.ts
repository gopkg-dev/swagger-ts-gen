import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  CommonDict,
  DeptTreeQueryParam,
  OptionDictsQueryParam,
  Tree,
  UploadCommonFileBody,
  UploadResult,
} from './model';

/**
 * 查询参数字典
 * @param params - 查询参数
 * @returns Promise<CommonDict[]>
 */
export async function queryOptionDicts(params?: OptionDictsQueryParam) {
  const res = await request.get<ApiResult<CommonDict[]>>('/api/v1/common/dict/option', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询参数字典失败'));
}


/**
 * 查询系统配置参数
 * @returns Promise<CommonDict[]>
 */
export async function querySiteOptionDicts() {
  const res = await request.get<ApiResult<CommonDict[]>>('/api/v1/common/dict/option/site');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询系统配置参数失败'));
}


/**
 * 查询角色字典
 * @returns Promise<CommonDict[]>
 */
export async function queryRoleDicts() {
  const res = await request.get<ApiResult<CommonDict[]>>('/api/v1/common/dict/role');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询角色字典失败'));
}


/**
 * 查询用户字典
 * @returns Promise<CommonDict[]>
 */
export async function queryUserDicts() {
  const res = await request.get<ApiResult<CommonDict[]>>('/api/v1/common/dict/user');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询用户字典失败'));
}


/**
 * 字典查询
 * @param code - 字典代码
 * @returns Promise<CommonDict[]>
 */
export async function queryCommonDicts(code: string) {
  const res = await request.get<ApiResult<CommonDict[]>>(`/api/v1/common/dict/${code}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '字典查询失败'));
}


/**
 * 上传文件
 * @param data - 请求数据
 * @returns Promise<UploadResult>
 */
export async function uploadCommonFile(data?: UploadCommonFileBody) {
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
  const res = await request.post<ApiResult<UploadResult>>('/api/v1/common/file', formData, { headers: { 'Content-Type': 'multipart/form-data' } });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '上传文件失败'));
}


/**
 * 查询部门树
 * @param params - 查询参数
 * @returns Promise<Tree[]>
 */
export async function queryDeptTree(params?: DeptTreeQueryParam) {
  const res = await request.get<ApiResult<Tree[]>>('/api/v1/common/tree/dept', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询部门树失败'));
}

