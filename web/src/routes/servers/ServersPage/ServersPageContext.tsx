import { createContext } from "react";
import { IAllServersItem } from "../../../types/servers";

export interface ServersPageContextInterface {
  servers: IAllServersItem[]
  isNewServer: boolean;
  setIsNewServer: (value:boolean)=>void,
};

const initialValue:ServersPageContextInterface = {
  servers: [],
  isNewServer: false,
  setIsNewServer: (value)=>{}
};

const ServersPageContext = createContext<ServersPageContextInterface>(initialValue);

export default ServersPageContext;