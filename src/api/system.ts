import type { baseResponse, listDataType, paramsType } from "@/api/index";
import { useAxios } from "@/api/index";

export interface navListType {
  item: string;
  title: string;
  href: string;
  src: string;
  description: string;
}

export function navList(): Promise<baseResponse<listDataType<navListType>>> {
  return useAxios.get("/api/navlist");
}

export interface circleListType {
  icon: () => JSX.Element;
  className: string;
  duration: number;
  delay: number;
  radius: number;
  reverse: boolean;
}

export function circleList(): Promise<
  baseResponse<listDataType<circleListType>>
> {
  return useAxios.get("/api/circlelist");
}
