/* eslint-disable react-hooks/rules-of-hooks */
import React from "react";
import { redirect } from "react-router-dom";
import { useUserStore } from "../../stores/userStore";

export async function loader() {
  const isLoggedIn = useUserStore.getState().isLoggedIn;
  if (isLoggedIn()) {
    return redirect("/");
  }
  return null;
}

export const LoginPage = () => {
  return <h1>LoginPage</h1>;
};
