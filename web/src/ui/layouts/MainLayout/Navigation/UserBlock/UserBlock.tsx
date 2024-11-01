import { useUserStore } from "../../../../../stores/userStore";
import UserAvatar from "./UserAvatar";
import UserInfo from "./UserInfo";


function UserBlock() {
  const { user, logout } = useUserStore.getState();

  return (
    <div className="px-3 flex justify-between gap-4 items-center">
      <UserAvatar user={user} />
      <UserInfo user={user} logout={logout} />
    </div>
  );
}

export default UserBlock;