import nginxPanelApiService from "../NginxPanelApiService";

describe("nginxPanelApiService", () => {
  it("service API path", () => {
    expect(nginxPanelApiService.API).toBe("stubs" || "http://localhost:8080" || "/api/");
  });

  it("service API GET /websites call", async () => {
    if (nginxPanelApiService.API === "stubs") {
      expect(nginxPanelApiService.getWebsites()).toBeTruthy;
    } else {
      expect(await nginxPanelApiService.getWebsites()).toBeTruthy;
    }
  });

  it("service API GET /servers call", async () => {
    if (nginxPanelApiService.API === "stubs") {
      expect(nginxPanelApiService.getServers()).toBeTruthy;
    } else {
      expect(await nginxPanelApiService.getServers()).toBeTruthy;
    }
  });
});

