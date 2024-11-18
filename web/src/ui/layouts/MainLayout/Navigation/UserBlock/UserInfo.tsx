import { IUser } from "../../../../../types/user";
import TextButton from "../../../../components/buttons/TextButton";


function UserInfo({ user, logout }: { user: IUser, logout?: () => void }) {
  return (
    <div className="font-normal tracking-tighter flex flex-col">
      <p>{user.login}</p>
      <p className="text-xs text-scndry-txt-clr">{user.email}</p>
      {!!logout && (<TextButton className="self-end w-fit" isDanger={true} onClick={logout}>Log out</TextButton>)}
    </div>
  );
}

export default UserInfo;