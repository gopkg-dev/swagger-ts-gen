import type { PageParam } from '@/api';

export interface DeleteDictItemForm {
  /** 字典项ID 列表 */
  ids: Array<string>;
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

export interface DictItemForm {
  /** 标签颜色 */
  color?: string;
  /** 描述 */
  description?: string;
  /** 字典 ID */
  dictId: string;
  /** 标签 */
  label: string;
  /** 排序 */
  sequence?: number;
  /** 状态（disable/enable） */
  status?: 'enable' | 'disable';
  /** 值 */
  value: string;
}

export interface DictItemsQueryParam extends PageParam {
  /** 字典项描述 */
  desc?: string;
  /** 字典ID */
  dictId?: string;
  /** 字典项标签 */
  label?: string;
  /** 状态 */
  status?: 'disable' | 'enable';
}
