import axios from "axios";

class NginxPanelApiService {
  API = "stubs/websites.json";

  async getWebsites() {
    let resp = await new Promise(
      (res) => {
        setTimeout(() => { res() }, 3000)
      }).then(() => {
        return axios.get(this.API);
      })

    return resp;
  }
}

const nginxPanelApiService = new NginxPanelApiService();

export default nginxPanelApiService;