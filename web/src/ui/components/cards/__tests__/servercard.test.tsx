import { getByTestId, render } from "@testing-library/react";
import renderer from "react-test-renderer";
import { WebsiteInterface } from "../../../../types";
import { ServerCard, ServerInfoInsideServerPage } from "../ServerCard";
import {
  IAllServersItem,
  IServerByIdResponse,
} from "../../../../types/servers";

const serversTestData: IAllServersItem = {
  id: 1,
  serverName: "server 1",
  domainName: "domain 1",
  ip: "192.158.1.31",
  numberOfSites: 1,
  active: "active",
};

const serverConfigTestData: IServerByIdResponse = {
  id: 1,
  serverName: "server 1",
  domainName: "domain 1",
  ip: "192.158.1.31",
  active: "active",
  sshKey: "key_key_key",
  connectedWebsites: [
    {
      id: 0,
      name: "First website",
      url: "somedomain.com/",
      ipCount: 3,
      upstreamsCount: 3,
      status: "active",
    },
    {
      id: 1,
      name: "Second website",
      url: "anotherdomain.com/",
      ipCount: 5,
      upstreamsCount: 2,
      status: "inactive",
    },
  ],
};

describe("<ServerCard />", () => {
  it("renders server's card correctly", () => {
    const tree = renderer
      .create(<ServerCard server={serversTestData} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});

describe("<ServerInfoInsideServerPage />", () => {
  it("renders server config info card correctly inside server page", () => {
    const tree = renderer
      .create(<ServerInfoInsideServerPage server={serverConfigTestData} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
