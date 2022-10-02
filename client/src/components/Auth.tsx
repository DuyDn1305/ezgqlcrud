import jwtDecode from "jwt-decode";
import { ReactNode, useEffect } from "react";
import { toast } from "react-toastify";
import { useUserStore } from "../stores/userStore";

export default function Auth() {
  const _token = useUserStore(state => state._token);
  const setUser = useUserStore(state => state.setUser);

  async function getUser(email: string, token: string) {
    const resp = await fetch("http://localhost:8080/user/get/" + email, {
      method: "GET",
      headers: {
        Authentication: "Bearer " + token
      }
    });
    const res = await resp.json();
    if (!resp.ok) throw new Error(res.message);
    return res.user;
  }

  async function updateUser() {
    const token = localStorage["token"];
    if (token) {
      try {
        const tokenData: any = jwtDecode(token);
        const email = tokenData.email;
        // fetch user
        const user = await getUser(email, token);
        setUser(user);
      } catch {
        delete localStorage["token"];
      }
    }
  }

  useEffect(() => {
    updateUser();
  }, [_token]);

  return null;
}
