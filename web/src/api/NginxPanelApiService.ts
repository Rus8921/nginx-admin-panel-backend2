import axios, { AxiosResponse } from "axios";
import { WebsitesDataInterface } from "../types";

class NginxPanelApiService {
  API = "stubs/websites";

  async getWebsites(datasetId: number = 0) {
    let resp = await new Promise(
      (res: (data: AxiosResponse<WebsitesDataInterface>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsitesDataInterface> = await axios.get(`${this.API}/${datasetId}.json`)
          res(data)
        }, 500)
      }).then((data) => {
        return data;
      })

    return resp;
  }
}

const nginxPanelApiService = new NginxPanelApiService();

export default nginxPanelApiService;