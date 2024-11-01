import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";
import { IUser } from "../types/user";

interface IUserStore {
  user: IUser;

  getDataFromCookies: () => void;
  isLoggedIn: () => boolean;
  login: ({ user }: { user: IUser }) => Promise<void>;
  logout: () => void;
}

const userValue: IUser = {
  email: "example@example.com",
  login: "login",
  token: "token",
  refreshToken: "refreshToken",
  tokenExpiresMilliseconds: 1000,
  role: 1,
};

const userValueLogout: IUser = {
  email: "",
  login: "",
  refreshToken: "",
  role: -1,
  token: "",
  tokenExpiresMilliseconds: -1,
};

export const useUserStore = create<IUserStore>()(
  persist(
    (set, get) => ({
      user: userValueLogout,

      getDataFromCookies: () => {},

      isLoggedIn: () => {
        if (get().user?.email === "" && get().user?.token === "") {
          return false;
        } else {
          return true;
        }
      },

      login: async ({ user }: { user: IUser }) => {
        set({ user: user });
      },

      logout: () => {
        set({
          user: userValueLogout,
        });
      },
    }),
    {
      name: "user-store",
      storage: createJSONStorage(() => sessionStorage),
      partialize: (state) => ({
        user: state.user,
      }),
    }
  )
);
