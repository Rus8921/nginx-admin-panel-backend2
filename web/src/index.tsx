import React from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import { MainLayout } from "./ui/layouts/MainLayout/MainLayout";
import { LoginPage, loader as loginLoader } from "./routes/login/LoginPage";
import { PermissionsPage } from "./routes/permissions/PermissionsPage";
import { UserPage } from "./routes/permissions/users/UserPage";
import { AddUserPage } from "./routes/permissions/users/AddUserPage";
import { EditUserPage } from "./routes/permissions/users/EditUserPage";
import { WebsitesPage } from "./routes/websites/WebsitesPage";
import { WebsiteIdPage } from "./routes/websites/WebsiteIdPage";
import { AddWebsitePage } from "./routes/websites/AddWebsitePage";
import { EditWebsitePage } from "./routes/websites/EditWebsitePage";
import { ServersPage } from "./routes/servers/ServersPage";
import { ServerIdPage } from "./routes/servers/ServerIdPage";
import { AddServerPage } from "./routes/servers/AddServerPage";
import { EditServerPage } from "./routes/servers/EditServerPage";
import { NotFoundErrorPage } from "./routes/errors/404";
import "./index.css";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import { LogoutPage } from "./routes/logout/LogoutPage";
import {
  RegistrationPage,
  loader as registrationPageLoader,
} from "./routes/registration/RegistrationPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout />,
    errorElement: <NotFoundErrorPage />,
    children: [
      {
        index: true,
        loader: () => {
          return redirect("websites");
        },
      },
      {
        path: "websites",
        children: [
          {
            index: true,
            element: (
              <ProtectedRoute>
                <WebsitesPage />
              </ProtectedRoute>
            ),
          },
          {
            path: "add",
            element: (
              <ProtectedRoute>
                <AddWebsitePage />
              </ProtectedRoute>
            ),
          },
          {
            path: ":websiteId",
            element: (
              <ProtectedRoute>
                <WebsiteIdPage />
              </ProtectedRoute>
            ),
          },
        ],
      },
      {
        path: "servers",
        children: [
          {
            index: true,
            element: (
              <ProtectedRoute>
                <ServersPage />
              </ProtectedRoute>
            ),
          },
          {
            path: "add",
            element: (
              <ProtectedRoute>
                <AddServerPage />
              </ProtectedRoute>
            ),
          },
          {
            path: ":serverId",
            element: (
              <ProtectedRoute>
                <ServerIdPage />
              </ProtectedRoute>
            ),
          },
        ],
      },
      {
        path: "permissions",
        children: [
          {
            index: true,
            element: (
              <ProtectedRoute>
                <PermissionsPage />
              </ProtectedRoute>
            ),
          },
          {
            path: "user/add",
            element: (
              <ProtectedRoute>
                <AddUserPage />
              </ProtectedRoute>
            ),
          },
          {
            path: "user/:userId",
            element: (
              <ProtectedRoute>
                <UserPage />
              </ProtectedRoute>
            ),
          },
          {
            path: "user/:userId/edit",
            element: (
              <ProtectedRoute>
                <EditUserPage />
              </ProtectedRoute>
            ),
          },
        ],
      },
    ],
  },
  {
    path: "/login",
    loader: loginLoader,
    element: <LoginPage />,
  },
  {
    path: "/logout",
    element: <LogoutPage />,
  },
  {
    path: "/registration",
    loader: registrationPageLoader,
    element: <RegistrationPage />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
