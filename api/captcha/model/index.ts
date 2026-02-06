export interface Captcha {
  /** 调试模式下会返回验证码 */
  captchaCode?: string;
  /** 验证码ID */
  captchaId?: string;
  /** 验证码内容（Base64） */
  content?: string;
  /** 过期时间（秒） */
  expireTime?: number;
}

export interface CaptchaQueryParam {
  /** 验证码ID */
  id: string;
  /** 是否刷新验证码图片（reload=1） */
  reload?: number;
}

export interface SendEmailCodeForm {
  /** 邮箱地址 */
  email: string;
}
