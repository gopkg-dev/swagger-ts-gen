import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  BindSocialAccountResult,
  GetSocialAuthURLResult,
  LoginToken,
  SocialBindForm,
  SocialLoginForm,
  SocialProvider,
  SocialProviderForm,
  SocialProvidersQueryParam,
  UnbindSocialAccountResult,
} from './model';

/**
 * 第三方登录（公开接口）
 * @param data - 请求数据
 * @returns Promise<LoginToken>
 */
export async function socialLogin(data: SocialLoginForm) {
  const res = await request.post<ApiResult<LoginToken>>('/api/v1/social/login', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '第三方登录（公开接口）失败'));
}


/**
 * 创建第三方登录平台配置
 * @param data - 请求数据
 * @returns Promise<SocialProvider>
 */
export async function createSocialProvider(data: SocialProviderForm) {
  const res = await request.post<ApiResult<SocialProvider>>('/api/v1/social/providers', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建第三方登录平台配置失败'));
}


/**
 * 获取已启用的第三方登录平台列表（公开接口）
 * @returns Promise<SocialProvider[]>
 */
export async function getEnabledSocialProviders() {
  const res = await request.get<ApiResult<SocialProvider[]>>('/api/v1/social/providers');
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取已启用的第三方登录平台列表（公开接口）失败'));
}


/**
 * 查询第三方登录平台配置列表
 * @param params - 查询参数
 * @returns Promise<PageResult<SocialProvider>>
 */
export async function querySocialProviders(params?: SocialProvidersQueryParam) {
  const res = await request.get<ApiResult<PageResult<SocialProvider>>>('/api/v1/social/providers/page', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询第三方登录平台配置列表失败'));
}


/**
 * 根据ID获取第三方登录平台配置
 * @param id - 唯一 ID
 * @returns Promise<SocialProvider>
 */
export async function getSocialProvider(id: string) {
  const res = await request.get<ApiResult<SocialProvider>>(`/api/v1/social/providers/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取第三方登录平台配置失败'));
}


/**
 * 根据ID更新第三方登录平台配置
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateSocialProvider(id: string, data: SocialProviderForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/social/providers/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新第三方登录平台配置失败'));
}


/**
 * 根据ID删除第三方登录平台配置
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteSocialProvider(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/social/providers/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除第三方登录平台配置失败'));
}


/**
 * 禁用第三方登录平台配置
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function disableSocialProvider(id: string) {
  const res = await request.patch<ApiResult<void>>(`/api/v1/social/providers/${id}/disable`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '禁用第三方登录平台配置失败'));
}


/**
 * 启用第三方登录平台配置
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function enableSocialProvider(id: string) {
  const res = await request.patch<ApiResult<void>>(`/api/v1/social/providers/${id}/enable`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '启用第三方登录平台配置失败'));
}


/**
 * 获取授权链接（公开接口）
 * @param source - 平台编码（如：GitHub、TikTok）
 * @returns Promise<GetSocialAuthURLResult>
 */
export async function getSocialAuthURL(source: string) {
  const res = await request.get<ApiResult<GetSocialAuthURLResult>>(`/api/v1/social/${source}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '获取授权链接（公开接口）失败'));
}


/**
 * 绑定第三方账号（需登录）
 * @param source - 平台编码（如：GitHub、TikTok）
 * @param data - 请求数据
 * @returns Promise<BindSocialAccountResult>
 */
export async function bindSocialAccount(source: string, data: SocialBindForm) {
  const res = await request.post<ApiResult<BindSocialAccountResult>>(`/api/v1/social/${source}`, data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '绑定第三方账号（需登录）失败'));
}


/**
 * 解绑第三方账号（需登录）
 * @param source - 平台编码（如：GitHub、TikTok）
 * @returns Promise<UnbindSocialAccountResult>
 */
export async function unbindSocialAccount(source: string) {
  const res = await request.delete<ApiResult<UnbindSocialAccountResult>>(`/api/v1/social/${source}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '解绑第三方账号（需登录）失败'));
}

