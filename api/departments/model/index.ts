import type { PageParam } from '@/api';

export interface Department {
  /** 子部门 */
  children?: Array<Department>;
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
  /** 部门领导 */
  leaders?: Array<DepartmentLeader>;
  /** 名称 */
  name?: string;
  /** 上级部门 ID（来源于 Department.ID） */
  parentId?: string;
  /** 祖级列表（以 . 分隔） */
  parentPath?: string;
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
}

export interface DepartmentForm {
  /** 描述 */
  description?: string;
  /** 部门领导 */
  leaders?: Array<DepartmentLeader>;
  /** 部门名称 */
  name: string;
  /** 父级 ID */
  parentId?: string;
  /** 排序 */
  sequence?: number;
  /** 状态 */
  status?: 'enable' | 'disable';
}

export interface DepartmentLeader {
  /** 来源于 Department.ID */
  deptId?: string;
  /** 真实姓名 */
  nickName?: string;
  /** 来源于 User.ID */
  userId?: string;
  /** 用户名称 */
  userName?: string;
}

export interface DepartmentsQueryParam extends PageParam {
  /** 部门名称 */
  name?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'updated_at' | 'sequence';
  /** 状态 */
  status?: 'enable' | 'disable';
}
