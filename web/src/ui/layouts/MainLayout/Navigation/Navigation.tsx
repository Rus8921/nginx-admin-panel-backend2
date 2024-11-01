import { Layout, Server, Users } from "react-feather";
import { ISidebarItem } from "../../../../types/sidebarItem";
import { SidebarItem } from "./SidebarItem";

function Navigation() {
  const routes: Array<ISidebarItem> = [
    { ItemIcon: Layout, title: "Websites", route: "websites" },
    { ItemIcon: Server, title: "Servers", route: "servers" },
    { ItemIcon: Users, title: "Permissions", route: "permissions" },
  ];
  return (
    <nav className="w-full flex flex-col gap-1">
      {routes.map((item) => (
        <SidebarItem
          ItemIcon={item.ItemIcon}
          title={item.title}
          route={item.route}
          key={item.title}
        />
      ))}
    </nav>
  );
}

export default Navigation;
