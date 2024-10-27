import React from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
  redirect,
  replace,
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

const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout />,
    errorElement: <NotFoundErrorPage />,
    children: [
      {
        index: true,
        loader: () => {
          const user = true;
          if (!user) {
            return redirect("/login");
          } else {
            return replace("/websites");
          }
        },
      },
      {
        path: "websites",
        children: [
          {
            index: true,
            element: <WebsitesPage />,
          },
          {
            path: "add",
            element: <AddWebsitePage />,
          },
          {
            path: ":websiteId",
            element: <WebsiteIdPage />,
          },
          {
            path: ":websiteId/edit",
            element: <EditWebsitePage />,
          },
        ],
      },
      {
        path: "servers",
        children: [
          {
            index: true,
            element: <ServersPage />,
          },
          {
            path: "add",
            element: <AddServerPage />,
          },
          {
            path: ":serverId",
            element: <ServerIdPage />,
          },
          {
            path: ":serverId/edit",
            element: <EditServerPage />,
          },
        ],
      },
      {
        path: "permissions",
        children: [
          {
            index: true,
            element: <PermissionsPage />,
          },
          {
            path: "user/add",
            element: <AddUserPage />,
          },
          {
            path: "user/:userId",
            element: <UserPage />,
          },
          {
            path: "user/:userId/edit",
            element: <EditUserPage />,
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
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    {/* <UserProvider> */}
    <RouterProvider router={router} />
    {/* </UserProvider> */}
  </React.StrictMode>
);
