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

export interface Role {
  /** 角色代码 */
  code?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 数据权限（1：全部数据权限；2：本部门及以下数据权限；3：本部门数据权限；4：仅本人数据权限；5：自定义数据权限） */
  dataScope?: number;
  /** 角色部门列表 */
  departments?: Array<RoleDepartment>;
  /** 描述 */
  description?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否为系统内置数据 */
  isSystem?: boolean;
  /** 角色菜单列表 */
  menus?: Array<RoleMenu>;
  /** 角色名称 */
  name?: string;
  /** 排序（按降序） */
  sequence?: number;
  /** 状态 (启用，禁用) */
  status?: string;
  /** 修改人 */
  updateUser?: string;
  /** 更新人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface RoleDataScope {
  dataScope?: number;
  departments?: Array<Department>;
  selects?: Array<string>;
}

export interface RoleDataScopeForm {
  dataScope: number;
  selects?: Array<string>;
}

export interface RoleDepartment {
  /** 部门ID，来源于 Department.ID */
  deptId?: string;
  /** 部门名称，来源于 Department.Name */
  deptName?: string;
  /** 角色ID，来源于 Role.ID */
  roleId?: string;
}

export interface RoleForm {
  /** 角色代码 */
  code: string;
  /** 描述 */
  description?: string;
  /** 角色名称 */
  name: string;
  /** 排序 */
  sequence?: number;
  /** 状态 */
  status?: 'enable' | 'disable';
}

export interface RoleMenu {
  /** 菜单ID，来源于 Menu.ID */
  menuId?: string;
  /** 角色ID，来源于 Role.ID */
  roleId?: string;
}

export interface RoleMenuData {
  menus?: Array<RolePermission>;
  selects?: Array<string>;
}

export interface RoleMenuDataForm {
  selects?: Array<string>;
}

export interface RolePermission {
  /** 子按钮 */
  actions?: RolePermissions;
  /** 子菜单 */
  children?: RolePermissions;
  /** 菜单图标 */
  icon?: string;
  /** 唯一 ID */
  id?: string;
  /** 上级菜单ID（来源于 Menu.ID） */
  parentId?: string;
  /** 菜单标题 */
  title?: string;
}

export type RolePermissions = Array<RolePermission>;

export interface RoleUsersForm {
  /** 角色 ID */
  roleId: string;
  /** 用户 ID */
  userIds: Array<string>;
}

export interface RolesQueryParam extends PageParam {
  /** 角色名称 */
  name?: string;
  /** 结果类型 */
  resultType?: 'select';
  /** 状态 */
  status?: 'disable' | 'enable';
}
