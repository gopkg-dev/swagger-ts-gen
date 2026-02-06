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

export interface Menu {
  /** 子菜单 */
  children?: Array<Menu>;
  /** 权限代码（每个级别唯一） */
  code?: string;
  /** 组件路径 */
  component?: string;
  /** 创建者 */
  createUser?: string;
  /** 创建人用户名 */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 是否整页路由 */
  fullPage?: boolean;
  /** 是否隐藏 */
  hidden?: boolean;
  /** 隐藏面包屑 */
  hiddenBreadcrumb?: boolean;
  /** 菜单图标 */
  icon?: string;
  /** 唯一 ID */
  id?: string;
  /** 是否缓存 */
  keepAlive?: boolean;
  /** 组件名称 */
  name?: string;
  /** 上级菜单ID（来源于 Menu.ID） */
  parentId?: string;
  /** 上级菜单路径（以点分隔） */
  parentPath?: string;
  /** 路由路径 */
  path?: string;
  /** 菜单属性（JSON格式） */
  properties?: string;
  /** 重定向地址 */
  redirect?: string;
  /** 关联的接口列表 */
  resources?: Array<MenuApi>;
  /** 排序（按降序排序） */
  sequence?: number;
  /** 状态 (启用，禁用) */
  status?: string;
  /** 标题 */
  title?: string;
  /** 类型 (菜单 按钮 链接 框架) */
  type?: string;
  /** 修改者 */
  updateUser?: string;
  /** 更新人用户名 */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

export interface MenuApi {
  /** 接口ID，来源于 Api.ID */
  apiId?: string;
  /** API方法，来源于 Api.Method */
  apiMethod?: string;
  /** API路径，来源于 Api.Path */
  apiPath?: string;
  /** API描述，来源于 Api.Summary */
  apiSummary?: string;
  /** 菜单ID，来源于 Menu.ID */
  menuId?: string;
}

/** FormData */
export interface UpdateCurrentAvatarBody {
  /** 头像文件 */
  file: Blob;
}

export interface UpdateCurrentUser {
  /** 用户描述 */
  description?: string;
  /** 性别 */
  gender: string;
  /** 昵称 */
  nickName: string;
  /** 用户名 */
  userName: string;
}

export interface UpdatePassword {
  /** 新密码（md5） */
  newPassword: string;
  /** 旧密码（md5） */
  oldPassword: string;
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

export interface UserPermission {
  codes?: Array<string>;
  menus?: Array<Menu>;
  roles?: Array<string>;
  user?: User;
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

export interface UserSocial {
  /** 创建时间 */
  createdAt?: string;
  /** 唯一 ID */
  id?: string;
  /** 最后登录时间 */
  lastLoginTime?: string;
  /** 附加信息 */
  meta?: string;
  /** 开放 ID */
  openId?: string;
  /** 来源 */
  source?: string;
  /** 用户 ID，来源于 User.ID */
  userId?: string;
}
