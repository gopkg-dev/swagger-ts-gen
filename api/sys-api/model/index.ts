import type { PageParam } from '@/api';

export interface Api {
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 是否废弃 */
  deprecated?: boolean;
  /** 唯一 ID */
  id?: string;
  /** HTTP方法 */
  method?: string;
  /** API路径 */
  path?: string;
  /** 接口描述 */
  summary?: string;
  /** 接口分组（Swagger tags） */
  tag?: string;
  /** 修改人 */
  updateUser?: string;
  /** 更新人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface ApisQueryParam extends PageParam {
  /** HTTP方法 */
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';
  /** 路径模糊查询 */
  path?: string;
  /** 结果类型 */
  resultType?: 'default' | 'select';
  /** 接口分组标签 */
  tag?: string;
}

export interface DeleteApiForm {
  /** 接口ID列表 */
  ids: Array<string>;
}

export interface SyncResult {
  /** 新增 */
  added?: number;
  /** 删除（保留字段，暂不使用） */
  deleted?: number;
  /** 总数 */
  total?: number;
  /** 更新 */
  updated?: number;
}
