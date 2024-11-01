import { IUser } from "../../../../../types/user";

function UserAvatar({ user }: { user: IUser }) {

  return (
    <div className="w-14 h-14 rounded-full border-2 border-main-clr flex justify-center items-center">
      <p className="text-lg/5 font-normal tracking-tighter text-main-clr">{user.login[0].toUpperCase()}</p>
    </div>
  );
}

export default UserAvatar;