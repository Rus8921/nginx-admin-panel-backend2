export interface IAllServersItem {
  id: number;
  serverName: string;
  domainName: string;
  ip: string;
  numberOfSites: number;
  active: "active" | "inactive";
}

export interface IAllServersResponse {
  servers: IAllServersItem[];
}
