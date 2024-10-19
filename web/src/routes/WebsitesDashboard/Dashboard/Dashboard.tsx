import { useEffect, useState } from "react";
import WebsiteCard from "../../../components/WebsiteCard/WebsiteCard";
import nginxPanelApiService from "../../../services/NginxPanelApiService";
import "./Dashboard.css";
import { WebsiteInterface } from "../../../types";

function Dashboard() {
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
}

export default Dashboard;