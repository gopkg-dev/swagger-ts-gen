export interface CommonDict {
  disabled?: boolean;
  extend?: string;
  label?: string;
  value?: string;
}

export interface DeptTreeQueryParam {
  /** 显示名称 */
  likeName?: string;
  /** 排序方向 */
  sortDirection?: 'desc' | 'asc';
  /** 排序字段 */
  sortField?: 'created_at' | 'updated_at' | 'sequence';
  /** 部门状态 */
  status?: 'enable' | 'disable';
}

export interface OptionDictsQueryParam {
  /** 类别 */
  category?: string;
}

export interface Tree {
  children?: Trees;
  disabled?: boolean;
  id?: string;
  label?: string;
}

export type Trees = Array<Tree>;

/** FormData */
export interface UploadCommonFileBody {
  /** 文件 */
  file: Blob;
}

export interface UploadResult {
  id?: string;
  metadata?: string;
  thumbnailUrl?: string;
  url?: string;
}
