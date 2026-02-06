import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Logger,
  LoggersQueryParam,
  LoggersQueryParam2,
  LoggersQueryParam3,
} from './model';

/**
 * 清空日志记录
 * @returns Promise<void>
 */
export async function deleteAllLoggers() {
  const res = await request.delete<ApiResult<void>>('/api/v1/loggers');
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '清空日志记录失败'));
}


/**
 * 查询日志列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Logger>>
 */
export async function queryLoggers(params?: LoggersQueryParam) {
  const res = await request.get<ApiResult<PageResult<Logger>>>('/api/v1/loggers', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询日志列表失败'));
}


/**
 * 批量删除日志
 * @param params - 查询参数
 * @returns Promise<void>
 */
export async function deleteLoggersByIds(params: LoggersQueryParam2) {
  const res = await request.delete<ApiResult<void>>('/api/v1/loggers/batchDelete', { params });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '批量删除日志失败'));
}


/**
 * 按模块删除日志
 * @param params - 查询参数
 * @returns Promise<void>
 */
export async function deleteLoggersByModule(params: LoggersQueryParam3) {
  const res = await request.delete<ApiResult<void>>('/api/v1/loggers/deleteByModule', { params });
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '按模块删除日志失败'));
}


/**
 * 获取日志模块列表
 * @returns Promise<string[]>
 */
export async function getLoggerModules() {
  const res = await request.get<ApiResult<string[]>>('/api/v1/loggers/modules');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取日志模块列表失败'));
}


/**
 * 按ID删除日志
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteLogger(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/loggers/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '按ID删除日志失败'));
}


/**
 * 通过ID获取日志详情
 * @param id - 唯一 ID
 * @returns Promise<Logger>
 */
export async function getLogger(id: string) {
  const res = await request.get<ApiResult<Logger>>(`/api/v1/loggers/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '通过ID获取日志详情失败'));
}

