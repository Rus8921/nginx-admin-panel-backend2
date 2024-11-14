import { useContext } from "react";
import WebsiteConfigContext from "../WebsiteConfigContext";
import Section from "../../../../ui/components/sections/Section";

function IPAdresses() {
  const { configs } = useContext(WebsiteConfigContext);

  return !!configs ? (
    <Section title="IP Addresses">
      <p>{configs.ipAddresses.map((ipAddress) => (
        <span key={ipAddress}>{ipAddress}; </span>
      ))}</p>
      {/* <input type="text" /> */}
    </Section>
  ) : (<></>);
};

export default IPAdresses;
