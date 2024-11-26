export interface WebsiteInterface {
  id: number;
  name: string;
  url: string;
  ipCount: number;
  upstreamsCount: number;
  status: "active" | "inactive";
}

export interface WebsitesDataInterface {
  websites: WebsiteInterface[];
}

export interface UpstreamServerConfigInterface {
  id: number,
  name: string,
  param: string,
}

export interface UpstreamInterface {
  id: number,
  name: string,
  connectedServers: UpstreamServerConfigInterface[],
}

export interface LocationInterface {
  id: number,
  name: string,
  upstream: UpstreamInterface
}

export interface SSLCertificateInterface {
  id: number,
  crtFile: string,
  keyFile: string,
  expirationDate: string,
  isActive: boolean
}

export interface WebsiteConfigInterface {
  id: number,
  name: string,
  domain: string,
  ipAddresses: string[],
  upstreams: UpstreamInterface[],
  locations: LocationInterface[],
  sslCertificates: SSLCertificateInterface[],
}

export enum AddItemTargets {
  Website = "website",
  Server = "server",
}