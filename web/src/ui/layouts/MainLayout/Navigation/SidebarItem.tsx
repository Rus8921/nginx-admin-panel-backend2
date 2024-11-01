import { NavLink, useLocation } from "react-router-dom";
import { ISidebarItem } from "../../../../types/sidebarItem";

export const SidebarItem = ({ ItemIcon, title, route }: ISidebarItem) => {
  const location = useLocation();
  const active = location.pathname.replaceAll("/", "") === route;
  const activeStyleBg = active ? `bg-highlight` : "";
  const activeStyleText = active ? `text-main-clr` : "";
  const hoverStyleBg = active ? `bg-highlight` : `bg-scndry-highlight`;

  return (
    <NavLink
      to={route}
      className={`w-full pl-4 flex flex-row gap-2 items-center content-center p-2 rounded-md ${activeStyleBg} hover:${hoverStyleBg}`}
    >
      <ItemIcon className={`w-5 h-5 ${activeStyleText}`} />
      <span className={`text-sm ${activeStyleText} `}>{title}</span>
    </NavLink>
  );
};
