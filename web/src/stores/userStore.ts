import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";
import { IUser } from "../types/user";

interface IUserStore {
  user: IUser | null;

  getDataFromCookies: () => void;
  isLoggedIn: () => boolean;
  login: ({ user }: { user: IUser }) => Promise<void>;
  logout: () => void;
}

const userValue: IUser = {
  email: "example@example.com",
  login: "adminlogin",
  token: "token",
  refreshToken: "refreshToken",
  tokenExpiresMilliseconds: 1000,
  role: 1,
};

export const useUserStore = create<IUserStore>()(
  persist(
    (set, get) => ({
      user: userValue,

      getDataFromCookies: () => { },

      isLoggedIn: () => {
        if (
          !get().user &&
          get().user?.email === "" &&
          get().user?.token === ""
        ) {
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
          user: {
            email: "",
            login: "",
            refreshToken: "",
            role: -1,
            token: "",
            tokenExpiresMilliseconds: -1,
          },
        });
      },
    }),
    {
      name: "user-store",
      storage: createJSONStorage(() => sessionStorage),
      partialize: (state) => ({}),
    }
  )
);
