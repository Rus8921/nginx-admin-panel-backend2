import { useContext } from "react";
import WebsiteConfigContext from "../WebsiteConfigContext";
import Section from "../../../../ui/components/sections/Section";
import FlatTable, { TableData } from "../../../../ui/components/tables/FlatTable";

function Locations() {
  const { configs } = useContext(WebsiteConfigContext);
  let data: TableData[] = !!configs ? [{
    items: configs.locations.map((location) => {
      return {
        id: location.id,
        name: location.name,
        arrowSign: "->",
        upstream: location.upstream.name
      }
    })
  }] : []

  return !!configs ? (
    <Section title="Locations">
      <FlatTable data={data} />
      {/* {configs.locations.map((location) => (
        <p key={location.id}>{location.name} {"->"} {location.upstream.name}</p>
      ))} */}
    </Section>
  ) : (<></>);
};

export default Locations;
