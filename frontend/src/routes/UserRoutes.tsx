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

const HisCaim = Loadable(lazy(() => import("../pages/claim/HistoryClaim")));

const Claimrequest = Loadable(lazy(() => import("../pages/claim/Claim_request")));


const UserRoutes = (isLoggedIn: boolean): RouteObject => {

  return {
    path: "/",
    element: isLoggedIn ? <UserLayout /> : <MainPages />,
    children: [
      {
        path: "/dashboard",
        element: <Dashboard />,
      },

      {
        path: "/customer",
        children: [
          {
            path: "/customer",
            element: <Customer />,
          },
          {
            path: "/customer/create",
            element: <CreateCustomer />,
          },
          {
            path: "/customer/edit/:id",
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
            path: "/profile",
            element: <ProfileEdit />,
          },
          {
            path: "/profile/address",
            element: <AddAddressPage />,
          }
        ]
      },
      {
        path: "/payment", 
        element: <PaymentPage />,
      },
      {
        path: "/historyclaim", 
        element: <HisCaim />,
      },
      {
        path: "/claimrequest", 
        element: <Claimrequest />,
      },
    ],
  };
};



export default UserRoutes;