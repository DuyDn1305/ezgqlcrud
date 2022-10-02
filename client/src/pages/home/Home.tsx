import jwtDecode from "jwt-decode";
import { useEffect } from "react";
import { useUserStore } from "../../stores/userStore";

export default function Home() {
  const user = useUserStore(state => state.user);

  return (
    <div>
      <div>email: {user?.email}</div>
      <div>name: {user?.name}</div>
    </div>
  );
}
