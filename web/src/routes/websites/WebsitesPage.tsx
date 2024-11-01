import React, { useEffect, useState } from "react";
import WebsiteCard from "../../ui/components/cards/WebsiteCard";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { WebsiteInterface } from "../../types";

export const WebsitesPage = ({ datasetId }: { datasetId?: number }) => {
  const [websites, setWebsites] = useState<WebsiteInterface[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let data = nginxPanelApiService.getWebsites(datasetId ?? 0)
    data.then((resp) => {
      if (resp.status === 200) {
        setWebsites(resp.data.websites)
        setIsLoading(false)
      }
    })
  }, [])

  return (
    <main className="w-full py-12 px-[4.5rem] grid grid-cols-adaptive-cards auto-rows-min gap-6 *:max-w-xl">
      {!isLoading ? (
        !!websites.length &&
        websites.map(website => <WebsiteCard key={website.id} data={website} />)
      ) : "loading..."}
    </main>
  )
};