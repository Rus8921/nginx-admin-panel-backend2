import axios, { AxiosResponse } from "axios";
import { WebsitesDataInterface } from "../types";
import { IAllServersItem, IAllServersResponse } from "../types/servers";

class NginxPanelApiService {
  API = "stubs";

  async getWebsites(datasetId: number = 0) {
    let resp = await new Promise(
      (res: (data: AxiosResponse<WebsitesDataInterface>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsitesDataInterface> = await axios.get(
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

  async getServers(datasetId: number = 0) {
    let resp = await new Promise(
      (res: (data: AxiosResponse<IAllServersResponse>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<IAllServersResponse> = await axios.get(
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
