import React from "react";
import { Outlet } from "react-router-dom";
import Header from "../Header/Header";
import Navigation from "./Navigation/Navigation";
import "./MainLayout.css"

export const MainLayout = () => {
  return (
    <div className="MainLayout">
      <Header />
      <div className="main">
        <Navigation />
        <Outlet />
      </div>
    </div>
  );
};
