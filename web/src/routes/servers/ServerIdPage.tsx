import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { IServerByIdResponse } from "../../types/servers";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { Loader, Trash2 } from "react-feather";
import { WebsiteCardInsideServerPage } from "../../ui/components/cards/WebsiteCard";
import { CommonButton } from "../../ui/components/buttons/CommonButton";
import { ServerInfoInsideServerPage } from "../../ui/components/cards/ServerCard";

const jsonString = `{"id":1,"serverName":"server 1","domainName":"domain 1","ip":"192.158.1.31","active":"active","sshKey":"key_key_key","connectedWebsites":[{"id":0,"name":"First website","url":"somedomain.com/","ipCount":3,"upstreamsCount":3,"status":"active"},{"id":1,"name":"Second website","url":"anotherdomain.com/","ipCount":5,"upstreamsCount":2,"status":"inactive"},{"id":2,"name":"Another website","url":"somedomain.com/","ipCount":7,"upstreamsCount":1,"status":"active"},{"id":3,"name":"One more website","url":"one.more.very.very.long.domain/","ipCount":15,"upstreamsCount":4,"status":"active"}]}`;

export const ServerIdPage = () => {
  const [serverData, setServerData] = useState<IServerByIdResponse | null>(
    null
  );
  const { serverId } = useParams();

  useEffect(() => {
    if (serverId) {
      // nginxPanelApiService.getServerById(Number(serverId)).then((response) => {
      //   console.log(`response: ${response.data}`);
      //   if (response.status === 200) {
      //     setServerData(response.data);
      //   }
      // });
      new Promise(function (resolve, reject) {
        setTimeout(resolve, 500);
      }).then(function () {
        const data: IServerByIdResponse = JSON.parse(jsonString);
        setServerData(data);
      });
    }
  }, [serverId]);

  return serverData ? (
    <main className="w-full p-9 flex flex-col">
      <ServerInfoInsideServerPage server={serverData} />
      <PageSeparator text="Websites connected" />
      <div className="flex flex-col gap-4">
        {serverData.connectedWebsites.map((website) => {
          return (
            <WebsiteCardInsideServerPage data={website} key={website.id} />
          );
        })}
      </div>
      <PageSeparator text="Configurations" />
      <form></form>
      <PageSeparator text="Danger zone" isRed={true} />
      <CommonButton isSubmit={false} type={"redBgWhiteText"} onClick={() => {}}>
        <div className="flex flex-row justify-center items-center gap-3">
          <Trash2 className="text-red" />
          <span className="text-red">DELETE</span>
        </div>
      </CommonButton>
    </main>
  ) : (
    <main className="relative w-full flex flex-col items-center justify-center">
      <Loader className="absolute animate-spin text-main-clr" />
    </main>
  );
};

const PageSeparator = ({
  text,
  isRed = false,
}: {
  text: string;
  isRed?: boolean;
}) => {
  let additionalStyle = ``;
  if (isRed) {
    additionalStyle = `text-red`;
  }
  return (
    <div className="py-5">
      <h3 className={`leading-9 ${additionalStyle}`}>{text}</h3>
      <hr className="text-scndry-txt-clr border mt-2"></hr>
    </div>
  );
};
