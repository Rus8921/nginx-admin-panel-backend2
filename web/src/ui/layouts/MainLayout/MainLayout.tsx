import React from "react";
import { Outlet } from "react-router-dom";
import Header from "../Header/Header";
import NavigationMenu from "./Navigation/NavigationMenu";

export const MainLayout = () => {
  return (
    <div className="w-full min-h-screen flex flex-col bg-bg-clr items-center">
      <Header />
      <div className="w-full flex  overflow-clip">
        <NavigationMenu />
        <Outlet />
      </div>
    </div>
  );
};
