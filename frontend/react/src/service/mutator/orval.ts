import type { AxiosError, AxiosRequestConfig } from "axios";

import { request } from "@/util/request";

/** Orval 生成接口统一使用的 Axios 请求适配器 */
export const customInstance = <T>(
	config: AxiosRequestConfig,
	options?: AxiosRequestConfig,
): Promise<T> => {
	return request({
		...config,
		...options,
		headers: {
			...config.headers,
			...options?.headers,
		},
	}) as Promise<T>;
};

/** Orval 生成的 React Query 错误类型 */
export type ErrorType<Error> = AxiosError<Error>;

/** Orval 生成的请求体类型 */
export type BodyType<BodyData> = BodyData;
