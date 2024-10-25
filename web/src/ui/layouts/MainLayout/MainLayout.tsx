import React from "react";
import { Outlet } from "react-router-dom";
import Header from "../Header/Header";
import Navigation from "./Navigation";

export const MainLayout = () => {
  return (
    <div className="min-h-screen flex flex-col bg-bg-clr items-center">
      <Header />
      <div className="w-full flex">
        <Navigation />
        <Outlet />
      </div>
    </div>
  );
};
