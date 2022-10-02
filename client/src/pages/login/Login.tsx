import bgImg from "../../assets/img/login-bg2.jpg";
import LoginForm from "./LoginForm";
import "./login.css";
import SignupForm from "./SignupForm";
import { useCallback, useState } from "react";
import { useLoginStore } from "../../stores/loginStore";

type Phase = "login" | "signup";

export default function Login() {
  const phase = useLoginStore(state => state.phase);

  return (
    <div className="min-h-screen flex justify-center items-center bg-login font-[Montserrat]">
      <div className="w-[1200px] flex shadow1 bg-white rounded-lg overflow-hidden">
        <div className="w-1/2 relative">
          <img src={bgImg} className="object-cover min-h-[74vh] h-full" />
          <div className="absolute inset-0 text-[#f5f6fa] flex flex-col p-10 justify-center items-center gap-3">
            <div className="text-[50px] uppercase font-bold">proj: EzEntGo</div>
            <div className="w-10 h-1.5 bg-white"></div>
            <div className="text-sm">
              Lorem ipsum dolor sit, amet consectetur adipisicing elit. Beatae
              architecto totam exercitationem illum rerum eos, minima voluptate
              nemo sunt quia voluptates, odio ipsam ex soluta animi! Provident
              voluptatibus quas dicta!
            </div>
          </div>
        </div>
        <div className={`w-1/2 relative phase-${phase}`}>
          <div className="login-panel login-panel-up">
            <SignupForm />
          </div>
          <div className="login-panel login-panel-down">
            <LoginForm />
          </div>
        </div>
      </div>
    </div>
  );
}
