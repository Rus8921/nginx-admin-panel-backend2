import React, { useEffect, useState } from "react";
import "./websites.css";
import WebsiteCard from "../../ui/components/cards/WebsiteCard/WebsiteCard";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { WebsiteInterface } from "../../types";

export const WebsitesPage = () => {
  const [websites, setWebsites] = useState<WebsiteInterface[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let data = nginxPanelApiService.getWebsites()
    data.then((resp) => {
      if (resp.status === 200) {
        setWebsites(resp.data.websites)
        setIsLoading(false)
      }
    })
  }, [])

  return (
    <main className="Dashboard">
      {!isLoading ? (
        !!websites.length &&
        websites.map(website => <WebsiteCard key={website.id} data={website} />)
      ) : "loading..."}
    </main>
  )
};