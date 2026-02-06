import type { PageParam } from '@/api';

export interface Logger {
  /** IP 归属地 */
  address?: string;
  /** 浏览器 */
  browser?: string;
  /** 创建人用户真名（来源于 User.NickName） */
  createNickName?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名称（来源于 User.UserName） */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 日志描述 */
  description?: string;
  /** 错误信息 */
  errorMsg?: string;
  /** ID */
  id?: string;
  /** IP */
  ip?: string;
  /** 耗时（毫秒） */
  latency?: number;
  /** 所属模块 */
  module?: string;
  /** 操作系统 */
  os?: string;
  /** 请求体 */
  requestBody?: string;
  /** 请求头 */
  requestHeaders?: string;
  /** 请求方式 */
  requestMethod?: string;
  /** 请求URL */
  requestURL?: string;
  /** 响应体 */
  responseBody?: string;
  /** 响应头 */
  responseHeaders?: string;
  /** 状态（成功/失败） */
  status?: boolean;
  /** 响应状态码 */
  statusCode?: number;
  /** 链路ID */
  traceId?: string;
}

export interface LoggersQueryParam extends PageParam {
  /** 结束时间 */
  endTime?: string;
  /** 所属模块 */
  module?: string;
  /** 搜索字段 */
  searchField?: 'create_user' | 'trace_id' | 'description' | 'request_url' | 'error_msg' | 'ip';
  /** 搜索内容 */
  searchValue?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'latency' | 'status_code' | 'module' | 'status' | 'ip';
  /** 开始时间 */
  startTime?: string;
  /** 状态（成功/失败） */
  status?: boolean;
  /** 时间字段 */
  timeField?: 'created_at';
  /** 链路ID */
  traceId?: string;
  /** 创建人ID列表 */
  userIds?: Array<string>;
}

export interface LoggersQueryParam2 {
  /** 日志 ID 数组 */
  ids: Array<string>;
}

export interface LoggersQueryParam3 {
  /** 模块名称 */
  module: string;
}
