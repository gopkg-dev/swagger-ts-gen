import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  LoginForm,
  LoginToken,
} from './model';

/**
 * 账号登录
 * @param data - 请求数据
 * @returns Promise<LoginToken>
 */
export async function login(data: LoginForm) {
  const res = await request.post<ApiResult<LoginToken>>('/api/v1/login', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '账号登录失败'));
}

