import request from '@/utils/request';
import type { ApiResult } from '@/api';
import type {
  Captcha,
  CaptchaQueryParam,
  SendEmailCodeForm,
} from './model';

/**
 * 发送邮箱验证码
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function sendMailCaptcha(data: SendEmailCodeForm) {
  const res = await request.post<ApiResult<void>>('/api/v1/captcha/email', data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '发送邮箱验证码失败'));
}


/**
 * 获取图片验证码ID
 * @returns Promise<Captcha>
 */
export async function getCaptcha() {
  const res = await request.get<ApiResult<Captcha>>('/api/v1/captcha/id');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取图片验证码ID失败'));
}


/**
 * 返回验证码图片
 * @param params - 查询参数
 * @returns Promise<Captcha>
 */
export async function getCaptchaContent(params: CaptchaQueryParam) {
  const res = await request.get<ApiResult<Captcha>>('/api/v1/captcha/image', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '返回验证码图片失败'));
}

