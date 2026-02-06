import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  MonitorData,
} from './model';

/**
 * 获取系统监控信息
 * @returns Promise<MonitorData>
 */
export async function getMonitorData() {
  const res = await request.get<ApiResult<MonitorData>>('/api/v1/monitor');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取系统监控信息失败'));
}

