import { useEffect, useState } from "react";
import { useUserStore } from "../../../stores/userStore";
import TextButton from "../../components/buttons/TextButton";


function Navigation() {
  const { isLoggedIn, user, logout } = useUserStore.getState();
  console.log(user)

  return (
    <aside className="w-60 h-[calc(100vh-7rem)] py-12 px-2 sticky top-28 flex flex-col justify-between items-center bg-white border-r border-r-scndry-txt-clr shadow-r-xs">
      <nav>navigation to be done</nav>

      {!!isLoggedIn() && !!user &&
        (
          <div className="flex justify-between gap-4 items-center">
            <div className="w-14 h-14 rounded-full border-2 border-main-clr flex justify-center items-center">
              <p className="text-lg/5 font-normal tracking-tighter text-main-clr">{user.login[0]?.toUpperCase()}</p>
            </div>
            <div className="font-normal tracking-tighter">
              <p>{user.login}</p>
              <p className="text-xs text-scndry-txt-clr">{user.email}</p>
              <TextButton onClick={logout}>Log out</TextButton>
            </div>
          </div>
        )}
    </aside>
  );
}

export default Navigation;