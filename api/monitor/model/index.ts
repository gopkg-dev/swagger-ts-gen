export interface CPUInfo {
  /** 逻辑核心数 */
  logicalCores?: number;
  /** 频率 */
  mhz?: number;
  /** CPU 型号 */
  modelName?: string;
  /** CPU 使用率 */
  percent?: number;
  /** 物理核心数 */
  physicalCores?: number;
}

export interface DiskInfo {
  /** 分区列表 */
  partitions?: Array<PartitionInfo>;
}

export interface HistoryData {
  /** CPU 使用率历史（包含系统和进程） */
  cpu?: Array<HistoryPoint>;
  /** 系统负载历史（1分钟平均负载） */
  load?: Array<HistoryPoint>;
  /** 内存使用历史（包含系统和进程） */
  memory?: Array<MemoryHistoryPoint>;
  /** 网络流量历史 */
  network?: Array<NetworkPoint>;
}

export interface HistoryPoint {
  /** 进程值 */
  processValue?: number;
  /** 系统值 */
  systemValue?: number;
  /** 时间戳 */
  timestamp?: string;
}

export interface HostInfo {
  /** 主机名 */
  hostname?: string;
  /** 内核架构 */
  kernelArch?: string;
  /** 内核版本 */
  kernelVersion?: string;
  /** 操作系统 */
  os?: string;
  /** 平台 */
  platform?: string;
  /** 平台版本 */
  platformVersion?: string;
  /** 运行时间 */
  uptime?: number;
}

export interface LoadInfo {
  /** 1分钟 */
  load1?: number;
  /** 15分钟 */
  load15?: number;
  /** 5分钟 */
  load5?: number;
}

export interface MemoryHistoryPoint {
  /** 进程内存使用率 (%) */
  processPercent?: number;
  /** 进程内存使用量 (bytes) */
  processUsed?: number;
  /** 系统内存使用率 (%) */
  systemPercent?: number;
  /** 系统内存使用量 (bytes) */
  systemUsed?: number;
  /** 时间戳 */
  timestamp?: string;
}

export interface MemoryInfo {
  /** 可用 */
  available?: number;
  /** 空闲 */
  free?: number;
  /** 总内存 */
  total?: number;
  /** 已使用 */
  used?: number;
  /** 使用率 */
  usedPercent?: number;
}

export interface MonitorData {
  /** 历史记录 */
  history?: HistoryData;
  /** 当前进程信息 */
  process?: ProcessInfo;
  /** 系统信息 */
  system?: SystemInfo;
  /** 时间戳 */
  timestamp?: string;
}

export interface NetInterfaceInfo {
  /** 接收字节数 */
  bytesRecv?: number;
  /** 发送字节数 */
  bytesSent?: number;
  /** 接口名称 */
  name?: string;
}

export interface NetworkInfo {
  /** 连接总数 */
  connectionCount?: number;
  /** 下载速度 (bytes/s) */
  downSpeed?: number;
  /** 网络接口列表 */
  interfaces?: Array<NetInterfaceInfo>;
  /** 上传速度 (bytes/s) */
  upSpeed?: number;
}

export interface NetworkPoint {
  /** 接收字节数（累计） */
  bytesRecv?: number;
  /** 发送字节数（累计） */
  bytesSent?: number;
  /** 下载速度 (bytes/s) */
  downSpeed?: number;
  /** 时间戳 */
  timestamp?: string;
  /** 上传速度 (bytes/s) */
  upSpeed?: number;
}

export interface PartitionInfo {
  /** 空闲 */
  free?: number;
  /** 挂载点 */
  mountPoint?: string;
  /** 总容量 */
  total?: number;
  /** 使用率 */
  usedPercent?: number;
}

export interface ProcessInfo {
  /** CPU 使用率 (%) */
  cpu?: number;
  /** Go 版本 */
  goVersion?: string;
  /** Goroutine 数量 */
  goroutines?: number;
  /** 内存使用 (bytes) */
  memory?: number;
  /** 内存使用率 (%) */
  memoryPercent?: number;
  /** 进程 ID */
  pid?: number;
  /** 启动时间戳 (毫秒) */
  startTime?: number;
  /** 线程数 */
  threads?: number;
  /** 运行时长 (秒) */
  uptime?: number;
  /** 工作目录 */
  workingDir?: string;
}

export interface SystemInfo {
  /** CPU 信息 */
  cpu?: CPUInfo;
  /** 磁盘信息 */
  disk?: DiskInfo;
  /** 主机信息 */
  host?: HostInfo;
  /** 系统负载 */
  load?: LoadInfo;
  /** 内存信息 */
  memory?: MemoryInfo;
  /** 网络信息 */
  network?: NetworkInfo;
}
