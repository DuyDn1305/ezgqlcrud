import create from "zustand";

type LoginPhase = "login" | "signup";
type LoginStore = {
  phase: LoginPhase;
  setPhase: (newPhase: LoginPhase) => void;
};

export const useLoginStore = create<LoginStore>()(set => ({
  phase: "login",
  setPhase: newPhase => set(() => ({ phase: newPhase }))
}));
