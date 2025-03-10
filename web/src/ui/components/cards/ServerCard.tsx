import { Link } from "react-router-dom";
import { IAllServersItem, IServerByIdResponse } from "../../../types/servers";
import Status from "../text/Status";
import Card from "./Card";
import ContentBreak from "../sections/ContentBreak";

export const ServerCard = ({ server, isClickable }: { server: IAllServersItem, isClickable: boolean }) => {
  return (
    <Card isClickable>
      <div className="flex flex-col gap-1">
        <div className="flex flex-row items-end gap-3">
          <span className="text-scndry-clr text-lg leading-9">
            {server.serverName}
          </span>
          <span className="text-scndry-txt-clr">{server.domainName}</span>
        </div>
        <ContentBreak />
        <div className="flex flex-row justify-between">
          <span className="text-scndry-txt-clr">{server.ip}</span>
          <Status status={server.active} />
        </div>
        <span className="text-scndry-txt-clr">
          {server.numberOfSites} websites connected
        </span>
      </div>
    </Card>
  );
};

export const ServerInfoInsideServerPage = ({
  server,
}: {
  server: IServerByIdResponse;
}) => {
  return (
    <div className="w-full p-8 bg-white border rounded-2xl border-scndry-clr active:shadow-none">
      <div className="flex flex-col gap-1">
        <div className="flex flex-row items-end gap-3">
          <span className="text-scndry-clr text-lg leading-9">
            {server.serverName}
          </span>
        </div>
        <ContentBreak />
        <span className="text-scndry-txt-clr">{server.domainName}</span>
        <span className="text-scndry-txt-clr">{server.ip}</span>
        <div className="flex flex-row justify-between pt-6">
          <span className="text-scndry-txt-clr">
            {server.connectedWebsites.length} websites connected
          </span>
          <Status status={server.active} />
        </div>
      </div>
    </div>
  );
};
