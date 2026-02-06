import type { PageParam } from '@/api';

export interface PassWordResult {
  /** 新密码 */
  passWord?: string;
}

export interface UpdatePassWordForm {
  /** 用户密码MD5 */
  passWord: string;
}

export interface User {
  /** 用户头像（base64） */
  avatar?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名，来源于 User.UserName */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 部门 ID */
  deptId?: string;
  /** 部门名称，来源于 Department.Name */
  deptName?: string;
  /** 用户描述 */
  description?: string;
  /** 邮箱账号 */
  email?: string;
  /** 用户性别（male/female/other） */
  gender?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为系统内置数据 */
  isSystem?: boolean;
  /** 用户昵称 */
  nickName?: string;
  /** 用户密码 */
  passWord?: string;
  /** 手机号码 */
  phone?: string;
  /** 最后一次修改密码时间 */
  pwdResetTime?: string;
  /** 用户角色 */
  roles?: Array<UserRole>;
  /** 用户状态（enable/disable） */
  status?: string;
  /** 修改人 */
  updateUser?: string;
  /** 修改人用户名，来源于 User.UserName */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
  /** 用户名称 */
  userName?: string;
}

export interface UserForm {
  /** 部门 ID */
  deptId: string;
  /** 描述 */
  description?: string;
  /** 邮箱 */
  email?: string;
  /** 性别 */
  gender?: 'male' | 'female' | 'other';
  /** 真实姓名 */
  nickName: string;
  /** 用户密码MD5 */
  passWord?: string;
  /** 手机 */
  phone?: string;
  /** 角色 */
  roles: Array<UserRole>;
  /** 排序 */
  sequence?: number;
  /** 状态 */
  status?: 'enable' | 'disable';
  /** 用户名称 */
  userName: string;
}

export interface UserRole {
  /** 角色代码，来源于 Role.Code */
  roleCode?: string;
  /** 角色ID，来源于 Role.ID */
  roleId?: string;
  /** 角色名称，来源于 Role.Name */
  roleName?: string;
  /** 用户ID，来源于 User.ID */
  userId?: string;
}

export interface UsersQueryParam extends PageParam {
  /** 部门ID */
  deptId?: string;
  /** 结束时间 */
  endTime?: string;
  /** 真实姓名 */
  nickname?: string;
  /** 手机号码 */
  phone?: string;
  /** 结果类型 */
  resultType?: 'select';
  /** 角色ID数组 */
  roleIds?: Array<string>;
  /** 搜索字段 */
  searchField?: 'username' | 'nickname' | 'phone' | 'email';
  /** 搜索内容 */
  searchValue?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'updated_at';
  /** 开始时间 */
  startTime?: string;
  /** 状态 */
  status?: 'enable' | 'disable';
  /** 时间字段 */
  timeField?: 'created_at' | 'updated_at';
}
