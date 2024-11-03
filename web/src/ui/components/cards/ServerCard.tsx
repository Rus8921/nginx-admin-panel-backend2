import { IAllServersItem } from "../../../types/servers";
import Status from "../text/Status";
import Card from "./Card";

export const ServerCard = ({ server }: { server: IAllServersItem }) => {
  return (
    <Card>
      <div className="flex flex-col gap-1">
        <div className="flex flex-row items-end gap-3">
          <span className="text-scndry-clr text-lg leading-9">
            {server.serverName}
          </span>
          <span className="text-scndry-txt-clr">{server.domainName}</span>
        </div>
        <hr className="text-scndry-txt-clr border"></hr>
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
