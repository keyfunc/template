import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

/** 鉴权状态 */
type AuthState = {
	/** 登录 token */
	token: string;
	/** 设置登录 token */
	setToken: (token: string) => void;
	/** 清空登录 token */
	clearToken: () => void;
};

export const useAuthStore = create<AuthState>()(
	persist(
		(set) => ({
			token: "",

			setToken: (token) => {
				set({ token });
			},

			clearToken: () => {
				set({ token: "" });
			},
		}),
		{
			name: "auth-storage",
			storage: createJSONStorage(() => localStorage),
		},
	),
);
