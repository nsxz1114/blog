import axios from "axios";
import { message } from "antd";

export const useAxios = axios.create({
  baseURL: "http://127.0.0.1:8888",
});

export interface baseResponse<T> {
  code: number;
  data: T;
  msg: string;
}

export interface listDataType<T> {
  count: number;
  list: T[];
}

export interface paramsType {
  page?: number;
  page_size?: number;
  keyword?: string;
  sort?: string;
}

// 请求拦截器
useAxios.interceptors.request.use((config) => {
  return config;
});

// 响应拦截器
useAxios.interceptors.response.use(
  async (response) => {
    if (response.status !== 200) {
      console.log("服务失败", response.status);
      message.warning(response.statusText);
      return Promise.reject(response.statusText);
    }
    return response.data;
  },
  (err) => {
    console.log("服务错误", err);
    message.error(err.message);
    return Promise.reject(err.message);
  }
);
