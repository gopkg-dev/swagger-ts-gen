import type { PageParam } from '@/api';

export interface BindSocialAccountResult {
  message?: string;
}

export interface GetSocialAuthURLResult {
  authUrl?: string;
}

export interface LoginToken {
  /** 访问令牌（JWT） */
  accessToken?: string;
  /** 过期时间（秒） */
  expiresAt?: number;
  /** 令牌类型（用法：Authorization=${token_type} ${access_token}） */
  tokenType?: string;
  /** 用户 ID */
  userId?: string;
}

export interface SocialBindForm {
  /** 授权码 */
  code: string;
  /** 状态码 */
  state: string;
}

export interface SocialLoginForm {
  /** 授权码 */
  code: string;
  /** 状态码 */
  state: string;
}

export interface SocialProvider {
  /** 客户端 ID / App ID */
  clientId?: string;
  /** 客户端密钥 / App Secret */
  clientSecret?: string;
  /** 平台编码（唯一标识，如：github、tiktok） */
  code?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 删除时间（软删除） */
  deletedAt?: string;
  /** 描述信息 */
  description?: string;
  /** 平台图标 URL */
  iconUrl?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为系统内置 */
  isSystem?: boolean;
  /** 平台名称（显示用，如：GitHub、TikTok） */
  name?: string;
  /** 回调地址（前端地址） */
  redirectUri?: string;
  /** 排序值 */
  sequence?: number;
  /** 状态（enable/disable） */
  status?: string;
  /** 修改人 */
  updateUser?: string;
  /** 修改人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface SocialProviderForm {
  /** 客户端 ID */
  clientId: string;
  /** 客户端密钥 */
  clientSecret: string;
  /** 平台编码（只允许字母和数字） */
  code: string;
  /** 描述信息 */
  description?: string;
  /** 平台图标 URL */
  iconUrl?: string;
  /** 平台名称 */
  name: string;
  /** 回调地址（前端地址） */
  redirectUri: string;
  /** 排序值 */
  sequence?: number;
  /** 状态 */
  status?: 'enable' | 'disable';
}

export interface SocialProvidersQueryParam extends PageParam {
  /** 平台编码 */
  code?: string;
  /** 平台名称（模糊查询） */
  name?: string;
  /** 提供商类型 */
  providerType?: 'oauth2' | 'oidc' | 'saml';
  /** 结果类型 */
  resultType?: 'select';
  /** 状态 */
  status?: 'disable' | 'enable';
}

export interface UnbindSocialAccountResult {
  message?: string;
}
