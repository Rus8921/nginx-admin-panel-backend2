import { Navigate, useNavigate } from "react-router-dom";
import { useUserStore } from "../../stores/userStore";

export const LogoutPage = () => {
  const logout = useUserStore((state) => state.logout);
  logout();
  return <Navigate to={"/login"} replace />;
};
