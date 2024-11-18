import React, { useState, useEffect } from "react";
import { WebsiteConfigInterface, WebsiteInterface } from "../../../types/index";
import WebsiteCard from "../../../ui/components/cards/WebsiteCard";
import nginxPanelApiService from "../../../api/NginxPanelApiService";
import { Loader, Trash2 } from "react-feather";
import { useParams } from "react-router";
import WebsiteConfigContext, { WebsiteConfigContextInterface } from "./WebsiteConfigContext";
import Section from "../../../ui/components/sections/Section";
import { CommonButton } from "../../../ui/components/buttons/CommonButton";
import IPAdresses from "./ConfigSections/IPAddresses";
import Upstreams from "./ConfigSections/Upstreams";
import Locations from "./ConfigSections/Locations";
import SSLCertificates from "./ConfigSections/SSLCertificates";
import Card from "../../../ui/components/cards/Card";

export const WebsiteIdPage = () => {
  const websiteId = useParams().websiteId;
  const [website, setWebsite] = useState<WebsiteInterface | undefined>(undefined);
  const [configs, setConfigs] = useState<WebsiteConfigInterface | undefined>(undefined);
  const [isLoading, setIsLoading] = useState(true);

  const contextValue: WebsiteConfigContextInterface = {
    configs: configs,
    setConfigs: setConfigs,
  }

  useEffect(() => {
    if (!!websiteId) {
      let data = nginxPanelApiService.getWebsite(websiteId);
      data.then((resp) => {
        if (resp.status === 200) {
          setWebsite(resp.data);
          return nginxPanelApiService.getWebsiteConfig(websiteId);
        } else {
          return Promise.reject(resp)
        }
      }).then((resp) => {
        if (resp.status === 200) {
          setConfigs(resp.data);
        }
      }).finally(() => {
        setIsLoading(false)
      });
    }
  }, [websiteId]);

  return isLoading ? (
    <div className="relative w-full flex flex-col items-center justify-center">
      <Loader className="absolute animate-spin text-main-clr" />
    </div>
  ) : (
    <main className="w-full py-12 px-[4.5rem] flex flex-col gap-9">
      {!!website && (
        <WebsiteConfigContext.Provider value={contextValue}>
          <WebsiteCard data={website} isClickable={false} />
          {!!configs && (
            <>
              <IPAdresses />
              <Upstreams />
              <Locations />
              <SSLCertificates />
            </>
          )}
          <Section title="Danger Zone" className="text-red">
            <CommonButton type="redBgWhiteText" >
              <Trash2 size="1.5rem" />
              Delete
            </CommonButton>
          </Section>
        </WebsiteConfigContext.Provider>
      )}
    </main>
  );
};
