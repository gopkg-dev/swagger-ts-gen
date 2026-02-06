import type { PageParam } from '@/api';

export interface GenConfig {
  author?: string;
  businessName?: string;
  createdAt?: string;
  /** 关联字段配置 */
  fields?: Array<GenFieldConfig>;
  isOverride?: boolean;
  moduleName?: string;
  packageName?: string;
  tableName?: string;
  tablePrefix?: string;
  updateUser?: string;
  updateUserName?: string;
  updatedAt?: string;
}

export interface GenConfigForm {
  author?: string;
  businessName: string;
  fields?: Array<GenFieldConfig>;
  isOverride?: boolean;
  moduleName: string;
  packageName: string;
  tableName: string;
  tablePrefix?: string;
}

export interface GenConfigUpdateForm {
  fieldConfigs?: Array<GenFieldConfig>;
  genConfig: GenConfigForm;
}

export interface GenConfigsQueryParam extends PageParam {
  /** 结束时间 */
  endTime?: string;
  /** 结果类型 */
  resultType?: 'default' | 'select';
  /** 搜索字段 */
  searchField?: 'table_name' | 'module_name' | 'business_name' | 'author';
  /** 搜索值 */
  searchValue?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'updated_at' | 'table_name' | 'module_name';
  /** 开始时间 */
  startTime?: string;
  /** 时间字段 */
  timeField?: 'created_at' | 'updated_at';
}

export interface GenFieldConfig {
  column?: string;
  comment?: string;
  createdAt?: string;
  dictType?: string;
  formType?: string;
  goName?: string;
  goType?: string;
  id?: string;
  isForm?: boolean;
  isList?: boolean;
  isPrimaryKey?: boolean;
  isQuery?: boolean;
  isRequired?: boolean;
  jsonName?: string;
  queryType?: string;
  sort?: number;
  table?: string;
  updateUser?: string;
  updateUserName?: string;
  updatedAt?: string;
}

export interface GenFieldConfigsQueryParam extends PageParam {
  /** 结束时间 */
  endTime?: string;
  /** 是否表单字段 */
  isForm?: boolean;
  /** 是否列表字段 */
  isList?: boolean;
  /** 是否查询字段 */
  isQuery?: boolean;
  /** 搜索字段 */
  searchField?: 'column' | 'go_name' | 'json_name' | 'comment';
  /** 搜索值 */
  searchValue?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'sort' | 'id' | 'created_at' | 'updated_at';
  /** 开始时间 */
  startTime?: string;
  /** 表名 */
  table?: string;
  /** 时间字段 */
  timeField?: 'created_at' | 'updated_at';
}

export interface GenerateCodeResult {
  files?: Array<GeneratedFile>;
}

export interface GeneratedFile {
  content?: string;
  path?: string;
}
