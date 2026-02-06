import request from '@/utils/request';
import type { ApiResult, PageResult } from '@/api';
import type {
  Department,
  DepartmentForm,
  DepartmentsQueryParam,
} from './model';

/**
 * 创建部门
 * @param data - 请求数据
 * @returns Promise<Department>
 */
export async function createDepartment(data: DepartmentForm) {
  const res = await request.post<ApiResult<Department>>('/api/v1/departments', data);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '创建部门失败'));
}


/**
 * 查询部门列表
 * @param params - 查询参数
 * @returns Promise<PageResult<Department>>
 */
export async function queryDepartments(params?: DepartmentsQueryParam) {
  const res = await request.get<ApiResult<PageResult<Department>>>('/api/v1/departments', { params });
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '查询部门列表失败'));
}


/**
 * 根据ID更新部门
 * @param id - 唯一 ID
 * @param data - 请求数据
 * @returns Promise<void>
 */
export async function updateDepartment(id: string, data: DepartmentForm) {
  const res = await request.put<ApiResult<void>>(`/api/v1/departments/${id}`, data);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID更新部门失败'));
}


/**
 * 根据ID删除部门
 * @param id - 唯一 ID
 * @returns Promise<void>
 */
export async function deleteDepartment(id: string) {
  const res = await request.delete<ApiResult<void>>(`/api/v1/departments/${id}`);
  if (res.data.success) {
    return;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID删除部门失败'));
}


/**
 * 根据ID获取部门
 * @param id - 唯一 ID
 * @returns Promise<Department>
 */
export async function getDepartment(id: string) {
  const res = await request.get<ApiResult<Department>>(`/api/v1/departments/${id}`);
  if (res.data.success && res.data.data !== undefined) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.error?.message ?? '根据ID获取部门失败'));
}

