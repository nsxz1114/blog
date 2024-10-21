import axios from "axios";
import { useStore } from "@/stores";
import { Message } from "@arco-design/web-vue";
import Cookies from "js-cookie";
export const useAxios = axios.create({
  baseURL: "",
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
useAxios.interceptors.request.use(
  (config) => {
    const store = useStore();

    const token = store.userStoreInfo.token;
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
useAxios.interceptors.response.use(
  (response) => {
    if (response.status !== 200) {
      Message.error(response.statusText);
      return Promise.reject(response.statusText);
    }
    return response.data;
  },
  async (error) => {
    const store = useStore();
    const originalRequest = error.config;

    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true; // 防止无限循环

      try {
        const refreshToken = Cookies.get("refresh_token");
        if (!refreshToken) {
          Message.error("请重新登录");
          return Promise.reject("No refresh token");
        }

        // 请求刷新token的接口
        const res = await useAxios.post(
          "/api/refreshtoken",
          {},
          {
            withCredentials: true, // 确保请求会携带cookie
          }
        );
        const newAccessToken = res.data;

        // 更新本地存储的token
        store.userStoreInfo.token = newAccessToken;

        // 更新请求的 token 并重发原请求
        originalRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;
        return useAxios(originalRequest);
      } catch (err) {
        // 刷新 token 失败，跳转到登录页面
        Message.error("Token 刷新失败，请重新登录");
        store.logout();
        window.location.href = "/login";
        return Promise.reject(err);
      }
    }
    console.log("服务错误", error);
    Message.error(error.message);
    return Promise.reject(error.message);
  }
);
