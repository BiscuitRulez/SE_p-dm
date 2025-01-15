import { useRoutes, RouteObject } from "react-router-dom";
import AdminRoutes from "./AdminRoutes";
import MainRoutes from "./MainRoutes";
import UserRoutes from "./UserRoutes"; // เส้นทางสำหรับ User

function ConfigRoutes() {
  const isLoggedIn = localStorage.getItem("isLogin") === "true";
  const role = localStorage.getItem("role"); // ดึง role จาก localStorage

  let routes: RouteObject[] = [];

  if (isLoggedIn) {
    if (role === "admin") {
      // หาก role เป็น admin
      routes = [AdminRoutes(isLoggedIn), MainRoutes()];
    } else if (role === "user") {
      // หาก role เป็น user
      routes = [UserRoutes(isLoggedIn), MainRoutes()];
    }
  } else {
    // หากยังไม่ล็อกอิน
    routes = [MainRoutes()];
  }

  return useRoutes(routes);
}

export default ConfigRoutes;
