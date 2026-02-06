import type { PageParam } from '@/api';

export interface Dict {
  /** 编码 */
  code?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 描述 */
  description?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为系统内置数据 */
  isSystem?: boolean;
  /** 字典项 */
  items?: Array<DictItem>;
  /** 名称 */
  name?: string;
  /** 排序 */
  sequence?: number;
  /** 修改人 */
  updateUser?: string;
  /** 更新人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface DictForm {
  /** 编码 */
  code: string;
  /** 描述 */
  description?: string;
  /** 名称 */
  name: string;
  /** 排序 */
  sequence?: number;
}

export interface DictItem {
  /** 标签颜色 */
  color?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 描述 */
  description?: string;
  /** 字典 ID */
  dictId?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为系统内置数据 */
  isSystem?: boolean;
  /** 标签 */
  label?: string;
  /** 排序 */
  sequence?: number;
  /** 状态（disable/enable） */
  status?: string;
  /** 修改人 */
  updateUser?: string;
  /** 更新人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
  /** 值 */
  value?: string;
}

export interface DictsQueryParam extends PageParam {
  /** 字典编码 */
  code?: string;
  /** 字典名称 */
  name?: string;
  /** 结果类型 */
  resultType?: 'select';
}
