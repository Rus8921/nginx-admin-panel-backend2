import { WebsiteInterface } from ".";

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

export interface IServerByIdResponse {
  id: number;
  serverName: string;
  domainName: string;
  ip: string;
  active: "active" | "inactive";
  sshKey: string;
  connectedWebsites: WebsiteInterface[];
}
