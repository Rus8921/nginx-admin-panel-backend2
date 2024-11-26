import { createContext } from "react";
import { WebsiteInterface } from "../../../types";

export interface WebsitesPageContextInterface {
  websites: WebsiteInterface[]
  isNewWebsite: boolean;
  setIsNewWebsite: (value:boolean)=>void,
};

const initialValue:WebsitesPageContextInterface = {
  websites: [],
  isNewWebsite: false,
  setIsNewWebsite: (value)=>{}
};

const WebsitesPageContext = createContext<WebsitesPageContextInterface>(initialValue);

export default WebsitesPageContext;