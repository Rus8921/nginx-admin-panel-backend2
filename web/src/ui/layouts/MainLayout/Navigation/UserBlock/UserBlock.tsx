import { useNavigate } from "react-router-dom";
import { useUserStore } from "../../../../../stores/userStore";
import UserAvatar from "./UserAvatar";
import UserInfo from "./UserInfo";

function UserBlock() {
  const isLoggedIn = useUserStore((state) => state.isLoggedIn);
  const user = useUserStore((state) => state.user);
  const navigate = useNavigate();

  return isLoggedIn() ?
    (
      <div className="px-3 flex justify-between gap-4 items-center">
        <UserAvatar login={user.login} />
        <UserInfo
          user={user}
          logout={() => {
            navigate("/logout");
          }}
        />
      </div>
    ) : (<></>);
}

export default UserBlock;
