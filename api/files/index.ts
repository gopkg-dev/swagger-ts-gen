import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  File,
  UploadFileBody,
} from './model';

/**
 * 上传文件
 * @param data - 请求数据
 * @returns Promise<File>
 */
export async function uploadFile(data?: UploadFileBody) {
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
  const res = await request.post<ApiResult<File>>('/api/v1/files/upload', formData, { headers: { 'Content-Type': 'multipart/form-data' } });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '上传文件失败'));
}

