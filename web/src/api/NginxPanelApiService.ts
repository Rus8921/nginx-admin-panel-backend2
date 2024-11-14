import axios, { AxiosResponse } from "axios";
import { WebsiteInterface, WebsitesDataInterface, WebsiteConfigInterface } from "../types";
import { IAllServersItem, IAllServersResponse } from "../types/servers";

interface ApiErrorResponse {
  status: number,
  message: string,
}

class NginxPanelApiService {
  API = "/stubs";

  async getWebsites(datasetId: string = "index") {
    let resp = await new Promise(
      (res: (data: AxiosResponse<WebsitesDataInterface, ApiErrorResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsitesDataInterface, ApiErrorResponse> = await axios.get(
            `${this.API}/websites/${datasetId}.json`
          );
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
      (res: (data: AxiosResponse<WebsiteInterface, ApiErrorResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsiteInterface, ApiErrorResponse> = await axios.get(
            `${this.API}/websites/1/index.json`
          );
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
      (res: (data: AxiosResponse<WebsiteConfigInterface, ApiErrorResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsiteConfigInterface, ApiErrorResponse> = await axios.get(
            `${this.API}/websites/1/config.json`
          );
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
      (res: (data: AxiosResponse<IAllServersResponse, ApiErrorResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<IAllServersResponse, ApiErrorResponse> = await axios.get(
            `${this.API}/servers/${datasetId}.json`
          );
          res(data);
        }, 500);
      }
    ).then((data) => {
      return data;
    });

    return resp;
  }
}

const nginxPanelApiService = new NginxPanelApiService();

export default nginxPanelApiService;
