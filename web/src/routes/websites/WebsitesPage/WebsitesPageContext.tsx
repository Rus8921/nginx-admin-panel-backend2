import { createContext } from "react";

export interface WebsitesPageContextInterface {
  isNewWebsite: boolean;
  setIsNewWebsite: (value:boolean)=>void,
};

const initialValue:WebsitesPageContextInterface = {
  isNewWebsite: false,
  setIsNewWebsite: (value)=>{}
};

const WebsitesPageContext = createContext<WebsitesPageContextInterface>(initialValue);

export default WebsitesPageContext;