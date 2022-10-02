import React, { FormEvent, useCallback, useState } from "react";
import { redirect } from "react-router-dom";
import { useLoginStore } from "../../stores/loginStore";
import { SignUpSchema } from "../../schemas/loginSchema";
import { toast } from "react-toastify";

type UserInput = {
  email: string;
  password: string;
  name: string;
};

export default function SignupForm() {
  const setPhase = useLoginStore(useCallback(state => state.setPhase, []));

  async function onSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const form = e.currentTarget;
    const data = new FormData(form);

    const input: UserInput = {
      email: data.get("email") as string,
      password: data.get("password") as string,
      name: data.get("name") as string
    };

    const valid = SignUpSchema.safeParse(input);
    if (!valid.success) {
      const issue = valid.error.issues[0];
      toast.error(issue.path[0] + ": " + issue.message);
      return;
    }

    try {
      const resp = await fetch("http://localhost:8080/noauth/register", {
        method: "post",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(input)
      });

      const res = await resp.json();
      if (!resp.ok) throw new Error(res.message);
      console.log(res);
      setPhase("login");
    } catch (e) {
      toast.error((e as Error).message);
    }
  }

  return (
    <div className="gap-10 flex flex-col">
      <div className="text-center gap-10 flex flex-col">
        <div className="text-4xl text-[#273c75] font-bold">Create Account</div>
        <div className="text-sm text-[#dcdde1]">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic, quod?
          Sed ut doloremque sit! Modi, adipisci accusamus dolores, suscipit
          ullam fugit repudiandae possimus sint sequi error, dicta tempore rem
          animi!
        </div>
      </div>
      <form className="flex flex-col gap-5" onSubmit={onSubmit}>
        <input
          name="name"
          className="user-input"
          type="text"
          placeholder="Display name"
        />
        <input
          name="email"
          className="user-input"
          type="text"
          placeholder="duydn1305@gmail.com"
        />
        <input
          name="password"
          className="user-input"
          type="password"
          placeholder="Password"
        />
        <button
          type="button"
          onClick={() => setPhase("login")}
          className="self-end login-switch-btn"
        >
          <i className="fal fa-arrow-down"></i>
          Sign in
        </button>
        <button
          type="submit"
          className="bg-[#00a8ff] text-white py-4 rounded-full mt-10 font-bold hover:opacity-70 transition-all duration-300"
        >
          REGISTER
        </button>
      </form>
    </div>
  );
}
