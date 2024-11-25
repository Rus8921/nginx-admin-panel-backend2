import axios, { AxiosResponse } from "axios";
import {
  WebsiteInterface,
  WebsitesDataInterface,
  WebsiteConfigInterface,
} from "../types";
import {
  IAllServersItem,
  IAllServersResponse,
  IServerByIdResponse,
} from "../types/servers";
import { IPermission } from "../types/permissions";

interface ApiErrorResponse {
  status: number;
  message: string;
}

class NginxPanelApiService {
  API = "/stubs";

  async getWebsites(datasetId: string = "index") {
    let resp = await new Promise(
      (
        res: (
          data: AxiosResponse<WebsitesDataInterface, ApiErrorResponse>
        ) => void
      ) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsitesDataInterface, ApiErrorResponse> =
            await axios.get(`${this.API}/websites/${datasetId}.json`);
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }

  async getWebsite(websiteId: string) {
    let resp = await new Promise(
      (
        res: (data: AxiosResponse<WebsiteInterface, ApiErrorResponse>) => void
      ) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsiteInterface, ApiErrorResponse> =
            await axios.get(`${this.API}/websites/1/index.json`);
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }

  async getWebsiteConfig(websiteId: string) {
    let resp = await new Promise(
      (
        res: (
          data: AxiosResponse<WebsiteConfigInterface, ApiErrorResponse>
        ) => void
      ) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsiteConfigInterface, ApiErrorResponse> =
            await axios.get(`${this.API}/websites/1/config.json`);
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }

  async getServers(datasetId: number = 0) {
    let resp = await new Promise(
      (
        res: (
          data: AxiosResponse<IAllServersResponse, ApiErrorResponse>
        ) => void
      ) => {
        setTimeout(async () => {
          let data: AxiosResponse<IAllServersResponse, ApiErrorResponse> =
            await axios.get(`${this.API}/servers/${datasetId}.json`);
          console.log(data);
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }

  async getServerById(serverId: number) {
    let resp = await new Promise(
      (res: (data: AxiosResponse<IServerByIdResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<IServerByIdResponse> = await axios.get(
            `${this.API}/servers/_id/0.json`
          );
          console.log(`data: ${data}`);
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }

  async getAllPermissions() {
    const json = `[{"websiteName":"website_1","users":[{"userName":"user_1","userEmail":"user1@gmail.com","permissions":[1,2],"access":1},{"userName":"user_2","userEmail":"user2@gmail.com","permissions":[1],"access":2}]},{"websiteName":"website_2","users":[{"userName":"user_2","userEmail":"user2@gmail.com","permissions":[1],"access":1},{"userName":"user_4","userEmail":"user4@gmail.com","permissions":[1,2],"access":2}]}]`;
    const data: IPermission[] = JSON.parse(json);
    console.log("data", data);
    await new Promise((resolve) => setTimeout(resolve, 500));
    const dataToReturn = new Promise<IPermission[]>((resolve) => resolve(data));
    return dataToReturn;
  }
}

const nginxPanelApiService = new NginxPanelApiService();

export default nginxPanelApiService;
