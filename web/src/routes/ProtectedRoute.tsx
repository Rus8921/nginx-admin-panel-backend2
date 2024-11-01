import { useUserStore } from "../stores/userStore";
import { Navigate } from "react-router-dom";

export const ProtectedRoute = ({ children }: any) => {
  const isLoggedIn = useUserStore((state) => state.isLoggedIn);
  if (!isLoggedIn()) {
    return <Navigate to="/login" replace />;
  }
  return <>{children}</>;
};
