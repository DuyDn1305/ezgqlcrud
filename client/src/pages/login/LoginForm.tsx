import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { useLoginStore } from "../../stores/loginStore";
import { SignInSchema } from "../../schemas/loginSchema";
import { useUserStore } from "../../stores/userStore";

type UserInput = {
  email: string;
  password: string;
};

export default function LoginForm() {
  const setPhase = useLoginStore(state => state.setPhase);
  const setToken = useUserStore(state => state.setToken);
  const [remember, setRemenber] = useState(true);
  const navigate = useNavigate();

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const data = new FormData(form);

    const input: UserInput = {
      email: data.get("email") as string,
      password: data.get("password") as string
    };

    const valid = SignInSchema.safeParse(input);
    if (!valid.success) {
      const issue = valid.error.issues[0];
      toast.error(issue.path[0] + ": " + issue.message);
      return;
    }

    try {
      const resp = await fetch("http://localhost:8080/noauth/login", {
        method: "post",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(input)
      });

      const res = await resp.json();
      if (!resp.ok) throw new Error(res.message);
      console.log(res);
      localStorage["token"] = res.token;
      setToken();
      toast.success("Logged in");
      navigate("/");
    } catch (e) {
      toast.error((e as Error).message);
    }
  };

  return (
    <div className="gap-10 flex flex-col">
      <div className="text-center gap-10 flex flex-col">
        <div className="text-4xl text-[#273c75] font-bold">Login Account</div>
        <div className="text-sm text-[#dcdde1]">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic, quod?
          Sed ut doloremque sit! Modi, adipisci accusamus dolores, suscipit
          ullam fugit repudiandae possimus sint sequi error, dicta tempore rem
          animi!
        </div>
      </div>
      <form className="flex flex-col gap-5" onSubmit={onSubmit}>
        <input
          name="email"
          className="user-input"
          type="text"
          placeholder="Email"
        />
        <input
          name="password"
          className="user-input"
          type="password"
          placeholder="Password"
        />
        <div className="flex justify-between">
          <div
            className="flex gap-7 text-[#00a8ff]"
            onClick={() => setRemenber(!remember)}
          >
            <span>
              <div
                className={`remember-box ${
                  remember ? "scale-125 z-10" : "scale-0 z-0"
                }`}
              >
                <i className="fal fa-check"></i>
              </div>
              <div
                className={`remember-box ${
                  !remember ? "scale-125 z-10" : "scale-0 z-0"
                }`}
              >
                <i className="fal fa-square"></i>
              </div>
            </span>

            <span className="">Keep me signed in</span>
          </div>
          <button
            type="button"
            onClick={() => setPhase("signup")}
            className="login-switch-btn"
          >
            <i className="fal fa-arrow-up"></i>
            Sign up
          </button>
        </div>
        <button className="bg-[#00a8ff] text-white py-4 rounded-full mt-10 font-bold hover:opacity-70 transition-all duration-300">
          LOGIN
        </button>
      </form>
    </div>
  );
}
