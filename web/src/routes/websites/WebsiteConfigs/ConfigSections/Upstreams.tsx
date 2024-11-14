import { useContext } from "react";
import WebsiteConfigContext from "../WebsiteConfigContext";
import Section from "../../../../ui/components/sections/Section";
import FlatTable, { TableData } from "../../../../ui/components/tables/FlatTable";

function Upstreams() {
  const { configs } = useContext(WebsiteConfigContext);
  let data: TableData[] = configs?.upstreams.map((upstream) => {
    return {
      head: {
        id: upstream.id,
        name: upstream.name
      },
      items: upstream.connectedServers
    }
  }) ?? []

  return !!configs ? (
    <Section title="Upstreams">
      <FlatTable data={data} />
      {/* {configs.upstreams.map((upstream) => (
        <div key={upstream.id}>
          <p>{upstream.name}</p>
          {upstream.connectedServers?.map((server) => (
            <p key={server.id}>{server.name} {server.param}</p>
          ))}
        </div>
      ))} */}
    </Section>
  ) : (<></>);
};

export default Upstreams;
