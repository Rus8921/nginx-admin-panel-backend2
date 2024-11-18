import { createContext } from "react";
import { WebsiteConfigInterface } from "../../../types";

// export const configsInitValue: WebsiteConfigInterface = {
//   id: 0,
//   name: "",
//   domain: "",
//   ipAddresses: [],
//   upstreams: [],
//   locations: [],
//   sslCertificates: []
// }

export interface WebsiteConfigContextInterface {
  configs?: WebsiteConfigInterface | undefined,
  setConfigs?: React.Dispatch<React.SetStateAction<WebsiteConfigInterface | undefined>>
};

const WebsiteConfigContext = createContext<WebsiteConfigContextInterface>({});

export default WebsiteConfigContext;