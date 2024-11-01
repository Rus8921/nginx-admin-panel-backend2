export interface IUser {
  email: string;
  login: string;
  token: string;
  refreshToken: string;
  tokenExpiresMilliseconds: number;
  role: number;
}

export interface IUserContext {
  user: IUser;
  isLoggedIn: () => boolean;
  login: ({ user }: { user: IUser }) => void;
  logout: () => void;
}
