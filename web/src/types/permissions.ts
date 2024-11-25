export interface IPermission {
  websiteName: string;
  users: IPermissionUser[];
}

export interface IPermissionUser {
  userName: string;
  userEmail: string;
  permissions: number[];
  access: number;
}
