export interface WebsiteInterface {
  id: number,
  name: string,
  url: string,
  ipCount: number,
  upstreamsCount: number,
  status: "active" | "inactive",
}