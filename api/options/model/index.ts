export interface Option {
  /** 键 */
  code?: string;
  /** 描述 */
  description?: string;
  /** 唯一 ID */
  id?: string;
  /** 名称 */
  name?: string;
  /** 值 */
  value?: string;
}

export interface OptionsQueryParam {
  /** 参数类别 */
  category?: 'PASSWORD' | 'SITE' | 'MAIL';
  /** 参数键数组 */
  codes?: Array<string>;
}

export interface RestOptionValueForm {
  /** 类别 */
  category: string;
  /** 键（可选，不传则重置整个分类） */
  codes?: Array<string>;
}

export interface UpdateOption {
  /** 唯一 ID */
  id: string;
  /** 参数值 */
  value?: string;
}

export interface UpdateOptionsForm {
  /** 参数列表 */
  options: Array<UpdateOption>;
}
