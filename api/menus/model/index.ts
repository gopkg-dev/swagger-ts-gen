import type { PageParam } from '@/api';

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

export interface MenuForm {
  /** 菜单代码（每个级别唯一） */
  code?: string;
  /** 组件路径 */
  component?: string;
  /** 整页路由 */
  fullPage?: boolean;
  /** 是否隐藏 */
  hidden?: boolean;
  /** 隐藏面包屑 */
  hiddenBreadcrumb?: boolean;
  /** 菜单图标 */
  icon?: string;
  /** 是否缓存 */
  keepAlive?: boolean;
  /** 路由名称 */
  name?: string;
  /** 父级ID（来自Menu.ID） */
  parentId?: string;
  /** 路由路径 */
  path?: string;
  /** 菜单属性（JSON格式） */
  properties?: string;
  /** 重定向 */
  redirect?: string;
  /** 关联的接口列表 */
  resources?: Array<MenuApi>;
  /** 排序 */
  sequence?: number;
  /** 菜单状态（启用，禁用） */
  status?: 'enable' | 'disable';
  /** 菜单显示标题 */
  title: string;
  /** 菜单类型（目录 菜单 按钮 外链 内嵌） */
  type: 'directory' | 'menu' | 'button' | 'link' | 'iframe';
}

export interface MenusQueryParam extends PageParam {
  /** 菜单编码路径（如 xxx.xxx.xxx） */
  code?: string;
  /** 是否包含菜单资源 */
  includeResources?: boolean;
  /** 菜单名称 */
  name?: string;
  /** 路由地址 */
  path?: string;
  /** 菜单标题 */
  title?: string;
}
