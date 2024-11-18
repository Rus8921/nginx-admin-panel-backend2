import { useUserStore } from "../../../../stores/userStore";
import UserBlock from "./UserBlock/UserBlock";
import Navigation from "./Navigation";

function NavigationMenu() {
  const isLoggedIn = useUserStore.getState().isLoggedIn();

  return (
    <aside className="sticky top-28 max-w-56 h-[calc(100vh-7rem)] py-12 px-1 flex flex-col basis-56 flex-shrink justify-between items-center bg-white border-r border-r-scndry-txt-clr shadow-r-xs">
      <Navigation />
      {isLoggedIn && <UserBlock />}
    </aside>
  );
}

export default NavigationMenu;
