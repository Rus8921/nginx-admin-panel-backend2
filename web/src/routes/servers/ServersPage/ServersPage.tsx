import { useEffect, useState } from "react";
import { IAllServersItem } from "../../../types/servers";
import nginxPanelApiService from "../../../api/NginxPanelApiService";
import { Loader } from "react-feather";
import { ServerCard } from "../../../ui/components/cards/ServerCard";
import { Link } from "react-router-dom";
import AddNewButton from "../../../ui/components/buttons/AddNewButton";
import { AddItemTargets } from "../../../types";
import ServersPageContext, { ServersPageContextInterface } from "./ServersPageContext";
import NewServerModal from "./NewServerModal";

export const ServersPage = ({ datasetId }: { datasetId?: number }) => {
  const [servers, setServers] = useState<IAllServersItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isNewServer, setIsNewServer] = useState(false);

  const contextValue: ServersPageContextInterface = {
    servers: servers,
    isNewServer: isNewServer,
    setIsNewServer: setIsNewServer,
  }

  useEffect(() => {
    let data = nginxPanelApiService.getServers(datasetId ?? 0);
    data.then((resp) => {
      if (resp.status === 200) {
        setServers(resp.data.servers);
        setIsLoading(false);
      }
    });
  }, [datasetId]);

  return (<ServersPageContext.Provider value={contextValue}>
    {isLoading ? (
      <main className="relative w-full flex flex-col items-center justify-center">
        <Loader className="absolute animate-spin text-main-clr" />
      </main>
    ) : (
      <main className="w-full flex flex-col items-center py-12 px-[4.5rem] gap-6">
        {servers.length === 0 ? (
          <></>
        ) : (
          servers.map((item) => 
            <Link to={`/servers/${item.id}`} className="w-full">
              <ServerCard key={item.id} server={item} isClickable />
            </Link>
          )
        )}
        <AddNewButton target={AddItemTargets.Server} onClick={()=>{setIsNewServer(true)}}/>
      </main>
    )}
    {isNewServer && (
      <NewServerModal />
    )}
  </ServersPageContext.Provider>);
};
