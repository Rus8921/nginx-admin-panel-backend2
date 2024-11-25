import React, { useEffect, useState } from "react";
import WebsiteCard from "../../../ui/components/cards/WebsiteCard";
import nginxPanelApiService from "../../../api/NginxPanelApiService";
import { NavLink } from "react-router-dom";
import { AddItemTargets, WebsiteInterface } from "../../../types";
import { Loader } from "react-feather";
import AddNewButton from "../../../ui/components/buttons/AddNewButton";
import WebsitesPageContext, { WebsitesPageContextInterface } from "./WebsitesPageContext";
import NewWebsiteModal from "./NewWebsiteModal";

export const WebsitesPage = ({ datasetId }: { datasetId?: string }) => {
  const [websites, setWebsites] = useState<WebsiteInterface[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isNewWebsite, setIsNewWebsite] = useState(false);

  const contextValue: WebsitesPageContextInterface = {
    isNewWebsite: isNewWebsite,
    setIsNewWebsite: setIsNewWebsite,
  }

  useEffect(() => {
    let data = nginxPanelApiService.getWebsites(datasetId);
    data.then((resp) => {
      if (resp.status === 200) {
        setWebsites(resp.data.websites);
        setIsLoading(false);
      }
    });
  }, [datasetId]);

  return (<WebsitesPageContext.Provider value={contextValue}>
    {isLoading ? (
      <div className="relative w-full flex flex-col items-center justify-center">
        <Loader className="absolute animate-spin text-main-clr" />
      </div>
    ) : (
      <main className="w-full h-fit py-12 px-[4.5rem] grid grid-cols-adaptive-cards auto-rows-fr items-stretch justify-items-center gap-6 *:max-w-xl">
        {!!websites.length &&
          websites.map((website) => (
            <NavLink to={`/websites/${website.id}`} className="w-full">
              <WebsiteCard key={website.id} data={website} isClickable />
            </NavLink>
        ))}
        <AddNewButton target={AddItemTargets.Website} onClick={()=>{setIsNewWebsite(true)}} />
      </main>
    )}
    {isNewWebsite && (
      <NewWebsiteModal />
    )}
    </WebsitesPageContext.Provider>);
};
