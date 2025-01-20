import { lazy } from "react";

import { RouteObject } from "react-router-dom";

import Loadable from "../components/third-party/Loadable";

import UserLayout from "../layout/UserLayout";


const MainPages = Loadable(lazy(() => import("../pages/authentication/Login")));

const Dashboard = Loadable(lazy(() => import("../pages/customer/palm/dashboard")));

const Customer = Loadable(lazy(() => import("../pages/customer/palm/customer")));

const CreateCustomer = Loadable(lazy(() => import("../pages/customer/palm/customer/create")));

const EditCustomer = Loadable(lazy(() => import("../pages/customer/palm/customer/edit")));

const UserCodes = Loadable(lazy(() => import("../pages/customer/palm/code")));

const ProfileEdit = Loadable(lazy(() => import("../pages/customer/palm/profile")));

const AddAddressPage = Loadable(lazy(() => import("../pages/customer/palm/profile/address")));

const PaymentPage = Loadable(lazy(() => import("../pages/payment/payment")));

const History = Loadable(lazy(() => import("../pages/claim/History")));

const ClaimNotiUser = Loadable(lazy(() => import("../pages/claim/ClaimNotiUser")));


const UserRoutes = (isLoggedIn: boolean): RouteObject => {

  return {
    path: "/",
    element: isLoggedIn ? <UserLayout /> : <MainPages />,
    children: [
      {
        index: true, // Default path
        element: <Dashboard />, // กำหนดเส้นทางเริ่มต้น
      },
      {
        path: "/dashboard",
        element: <Dashboard />,
      },
      {
        path: "/customer",
        children: [
          {
            path: "",
            element: <Customer />,
          },
          {
            path: "create",
            element: <CreateCustomer />,
          },
          {
            path: "edit/:id",
            element: <EditCustomer />,
          },
        ],
      },
      {
        path: "/code",
        element: <UserCodes />,
      },
      {
        path: "/profile",
        children: [
          {
            path: "",
            element: <ProfileEdit />,
          },
          {
            path: "address",
            element: <AddAddressPage />,
          },
        ],
      },
      {
        path: "/payment",
        element: <PaymentPage />,
      },
      {
        path: "/history", 
        element: <History />,
      },
      {
        path: "/claimnotiuser", 
        element: <ClaimNotiUser />,
      },
      {
        path: "/home",
        element: <HomePage />
      },
      {
        path: "/cart",
        element: <CartPage />
      },
      {
        path: "/tracking",
        element: <Tracking />
      },

    ],
  };
};




export default UserRoutes;