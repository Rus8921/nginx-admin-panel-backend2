import axios, { AxiosResponse } from "axios";
import { WebsitesDataInterface } from "../types";

class NginxPanelApiService {
  API = "stubs/websites.json";

  async getWebsites() {
    let resp = await new Promise(
      (res: (data: AxiosResponse<WebsitesDataInterface>) => void) => {
        setTimeout(async () => {
          let data: AxiosResponse<WebsitesDataInterface> = await axios.get(this.API)
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