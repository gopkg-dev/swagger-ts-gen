import type { PageParam } from '@/api';

export interface DeletedAt {
  time?: string;
  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface Storage {
  /** Access Key */
  accessKey?: string;
  /** Bucket */
  bucketName?: string;
  /** 编码 */
  code?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名，来源于 User.UserName */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 删除时间 */
  deletedAt?: DeletedAt;
  /** 描述 */
  description?: string;
  /** 域名 */
  domain?: string;
  /** Endpoint */
  endpoint?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为默认存储 */
  isDefault?: boolean;
  /** 名称 */
  name?: string;
  /** 是否启用回收站 */
  recycleBinEnabled?: string;
  /** 回收站路径 */
  recycleBinPath?: string;
  /** Secret Key */
  secretKey?: string;
  /** 排序 */
  sequence?: number;
  /** 状态（disable/enable） */
  status?: string;
  /** 存储类型（1：本地存储，2：对象存储） */
  type?: number;
  /** 修改人 */
  updateUser?: string;
  /** 修改人用户名，来源于 User.UserName */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface StorageForm {
  /** 访问密钥 */
  accessKey: string;
  /** 桶名称 */
  bucketName: string;
  /** 编码 */
  code: string;
  /** 描述 */
  description?: string;
  /** 域名 */
  domain?: string;
  /** 终端节点 */
  endpoint: string;
  /** 是否为默认存储 */
  isDefault?: boolean;
  /** 名称 */
  name: string;
  /** 是否启用回收站 */
  recycleBinEnabled: string;
  /** 回收站路径 */
  recycleBinPath?: string;
  /** 私有密钥 */
  secretKey: string;
  /** 排序 */
  sequence?: number;
  /** 状态 */
  status?: 'enable' | 'disable';
  /** 存储类型（1：本地存储，2：对象存储） */
  type: number;
}

export interface StoragesQueryParam extends PageParam {
  /** 存储类型 */
  Type?: 1 | 2;
  /** 结果类型 */
  resultType?: 'select';
  /** 搜索内容 */
  searchValue?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'updated_at';
  /** 状态 */
  status?: 'enable' | 'disable';
}
