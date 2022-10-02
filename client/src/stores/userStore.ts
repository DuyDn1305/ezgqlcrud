import create from "zustand";

type UserStore = {
  user?: {
    name: string;
    email: string;
  };
  _token: {};
  setToken: () => void;
  setUser: (newUser: UserStore["user"]) => void;
};

export const useUserStore = create<UserStore>()(set => ({
  user: undefined,
  _token: {},
  setToken: () => set(() => ({ _token: {} })),
  setUser: newUser => set(prevState => ({ user: newUser }))
}));
