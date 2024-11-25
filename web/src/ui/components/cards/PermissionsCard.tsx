import { IPermission } from "../../../types/permissions";

export const PermissionsCard = ({ data }: { data: IPermission }) => {
  return (
    <>
      <div className="p-1 w-full flex flex-row justify-between items-center text-md font-bold text-scndry-txt-clr">
        {data.websiteName}
      </div>
      <hr className="text-scndry-txt-clr" />
      {data.users.map((user) => {
        return (
          <>
            <div className="w-full flex flex-row justify-between items-center">
              <UserDataComponent data={user.userName} />
              <UserDataComponent data={user.userEmail} />
              <UserDataComponent
                data={<PermissionComponent permissions={user.permissions} />}
              />
              <UserDataComponent
                data={<AccessComponent access={user.access} />}
              />
            </div>
            <hr className="text-scndry-txt-clr" />
          </>
        );
      })}
      <div className="p-1 w-full flex flex-row justify-center gap-1 items-center content-center text-scndry-txt-clr">
        -+-
      </div>
      <hr className="text-scndry-txt-clr" />
    </>
  );
};

const UserDataComponent = ({ data }: { data: any }) => {
  return (
    <div className="w-1/4 whitespace-nowrap overflow-hidden text-ellipsis h-6 text-scndry-txt-clr">
      {data}
    </div>
  );
};

const PermissionComponent = ({ permissions }: { permissions: number[] }) => {
  const permission1 = permissions.includes(1) ? "read" : "";
  const permission2 = permissions.includes(2) ? "write" : "";
  const divider = permissions.length > 1 ? "/" : "";
  return (
    <span className="text-scndry-txt-clr">
      {permission1}
      {divider}
      {permission2}
    </span>
  );
};

const AccessComponent = ({ access }: { access: number }) => {
  let accessValue = "";
  switch (access) {
    case 1:
      accessValue = "pending";
      break;
    case 2:
      accessValue = "granted";
      break;
    default:
      accessValue = "unknown";
      break;
  }
  return <span className="text-scndry-txt-clr">{accessValue}</span>;
};
