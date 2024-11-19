import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { IServerByIdResponse } from "../../types/servers";
import nginxPanelApiService from "../../api/NginxPanelApiService";
import { Loader, Settings, Terminal, Trash2 } from "react-feather";
import { WebsiteCardInsideServerPage } from "../../ui/components/cards/WebsiteCard";
import { CommonButton } from "../../ui/components/buttons/CommonButton";
import { ServerInfoInsideServerPage } from "../../ui/components/cards/ServerCard";
import { SubmitHandler, useForm } from "react-hook-form";
import { SSHInputField } from "../../ui/components/fields/SSHInputField";
import { useDebouncedCallback } from "use-debounce";

const jsonString = `{"id":1,"serverName":"server 1","domainName":"domain 1","ip":"192.158.1.31","active":"active","sshKey":"key_key_key","connectedWebsites":[{"id":0,"name":"First website","url":"somedomain.com/","ipCount":3,"upstreamsCount":3,"status":"active"},{"id":1,"name":"Second website","url":"anotherdomain.com/","ipCount":5,"upstreamsCount":2,"status":"inactive"},{"id":2,"name":"Another website","url":"somedomain.com/","ipCount":7,"upstreamsCount":1,"status":"active"},{"id":3,"name":"One more website","url":"one.more.very.very.long.domain/","ipCount":15,"upstreamsCount":4,"status":"active"}]}`;

export type SshKeyInputProps = {
  sshKey: string;
};

export const ServerConfigPage = () => {
  const [serverData, setServerData] = useState<IServerByIdResponse | null>(
    null
  );
  const { serverId } = useParams();
  const [sshKey, setSshKey] = useState(``);

  const {
    setValue,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SshKeyInputProps>();
  const onSubmit: SubmitHandler<SshKeyInputProps> = (data) => {};
  console.log(sshKey);

  useEffect(() => {
    if (serverId) {
      new Promise(function (resolve, reject) {
        setTimeout(resolve, 500);
      }).then(function () {
        const data: IServerByIdResponse = JSON.parse(jsonString);
        if (sshKey === "") {
          setSshKey(data.sshKey);
          setValue("sshKey", data.sshKey);
        }
        setServerData(data);
      });
    }
  }, [serverId]);

  const handleSshKeyChange = useDebouncedCallback((newSshKey: string) => {
    setSshKey(newSshKey);
  }, 500);

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
      <form className="flex flex-col gap-3" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-row gap-3 items-center">
          <span className="text-scndry-txt-clr">SSH key</span>
          <SSHInputField
            placeholder={"Enter SSH key"}
            type={"sshKey"}
            register={register}
            required={false}
            IconComponent={Terminal}
            onChange={handleSshKeyChange}
          />
        </div>
        <CommonButton
          isSubmit={false}
          type={"blueBgWhiteText"}
          onClick={() => {}}
        >
          <div className="flex flex-row justify-center items-center gap-3">
            <Settings />
            <span>SAVE</span>
          </div>
        </CommonButton>
      </form>
      <PageSeparator text="Danger zone" isRed={true} />
      <CommonButton isSubmit={false} type={"redBgWhiteText"} onClick={() => {}}>
        <div className="flex flex-row justify-center items-center gap-3">
          <Trash2/>
          <span>DELETE</span>
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
