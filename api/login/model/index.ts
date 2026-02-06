export interface LoginForm {
  /** 验证码 */
  captchaCode?: string;
  /** 验证码 ID */
  captchaId?: string;
  /** 登录密码（md5） */
  password: string;
  /** 登录账号 */
  username: string;
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
