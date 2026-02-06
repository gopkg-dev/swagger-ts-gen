export interface DeletedAt {
  time?: string;
  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface File {
  /** 内容类型 */
  contentType?: string;
  /** 创建人 */
  createUser?: string;
  /** 创建人用户名，来源于 User.UserName */
  createUserName?: string;
  /** 创建时间 */
  createdAt?: string;
  /** 删除时间 */
  deletedAt?: DeletedAt;
  /** 扩展名 */
  extension?: string;
  /** 唯一 ID */
  id?: string;
  /** 元数据 */
  metadata?: string;
  /** 名称 */
  name?: string;
  /** 原始名称 */
  originalName?: string;
  /** 父路径 */
  parentPath?: string;
  /** 路径 */
  path?: string;
  /** SHA256值 */
  sha256?: string;
  /** 大小（字节） */
  size?: number;
  /** 存储 ID, 来源于 Storage.ID */
  storageId?: string;
  /** 缩略图元数据 */
  thumbnailMetadata?: string;
  /** 缩略图名称 */
  thumbnailName?: string;
  /** 缩略图大小（字节） */
  thumbnailSize?: number;
  /** 类型（0：目录；1：其他；2：图片；3：文档；4：视频；5：音频） */
  type?: number;
  /** 修改人 */
  updateUser?: string;
  /** 修改人用户名，来源于 User.UserName */
  updateUserName?: string;
  /** 修改时间 */
  updatedAt?: string;
}

/** FormData */
export interface UploadFileBody {
  /** 文件 */
  file: Blob;
}
