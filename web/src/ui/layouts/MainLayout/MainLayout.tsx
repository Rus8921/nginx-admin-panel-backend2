import React from "react";
import { Outlet } from "react-router-dom";

export const MainLayout = () => {
  return (
    <div>
      <h1 className="flex">MainLayout</h1>
      <Outlet />
    </div>
  );
};
