import React, { useEffect, useState } from "react";
import WebsiteCard from "../../ui/components/cards/WebsiteCard";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { NavLink } from "react-router-dom";
import { WebsiteInterface } from "../../types";
import { Loader } from "react-feather";

export const WebsitesPage = ({ datasetId }: { datasetId?: string }) => {
  const [websites, setWebsites] = useState<WebsiteInterface[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let data = nginxPanelApiService.getWebsites(datasetId);
    data.then((resp) => {
      if (resp.status === 200) {
        setWebsites(resp.data.websites);
        setIsLoading(false);
      }
    });
  }, [datasetId]);

  return isLoading ? (
    <div className="relative w-full flex flex-col items-center justify-center">
      <Loader className="absolute animate-spin text-main-clr" />
    </div>
  ) : (
    <main className="w-full py-12 px-[4.5rem] grid grid-cols-adaptive-cards auto-rows-min justify-items-center gap-6 *:max-w-xl">
      {!!websites.length &&
        websites.map((website) => (
          <NavLink to={`/websites/${website.id}`} className="w-full">
            <WebsiteCard key={website.id} data={website} isClickable />
          </NavLink>
        ))}
    </main>
  );
};
