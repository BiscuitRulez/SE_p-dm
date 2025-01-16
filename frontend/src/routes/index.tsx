import { useRoutes, RouteObject, Navigate } from "react-router-dom";
import AdminRoutes from "./AdminRoutes";
import MainRoutes from "./MainRoutes";
import UserRoutes from "./UserRoutes"; // เส้นทางสำหรับ User

function ConfigRoutes() {
  const isLoggedIn = localStorage.getItem("isLogin") === "true";
  const role = localStorage.getItem("role") || "";

  let routes: RouteObject[] = [];

  if (isLoggedIn) {
    if (role === "admin") {
      routes = [AdminRoutes(isLoggedIn)];
    } else if (role === "user") {
      routes = [UserRoutes(isLoggedIn)];
    } else {
      routes = [{ path: "*", element: <Navigate to="/login" replace /> }];
    }
  } else {
    routes = [MainRoutes()];
  }

  // เพิ่ม fallback เสมอ
  routes.push({ path: "*", element: <Navigate to="/" replace /> });

  return useRoutes(routes);
}

export default ConfigRoutes;

