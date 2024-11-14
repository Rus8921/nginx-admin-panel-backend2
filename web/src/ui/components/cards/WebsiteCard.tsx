import Card from "./Card";
import Status from "../text/Status";
import { WebsiteInterface } from "../../../types";
import ContentBreak from "../sections/ContentBreak";
import { ComponentProps } from "react";

function WebsiteCard({ data, ...rest }: ComponentProps<"div"> & { data: WebsiteInterface, isClickable: boolean }) {
  let { id, name, url, ipCount, upstreamsCount, status } = data;

  return (
    <Card className="w-full flex flex-col gap-20" {...rest}>
      <div className="flex flex-col gap-4">
        <h2>{name}</h2>
        <ContentBreak />
        <p className="text-scndry-txt-clr">{url}</p>
      </div>
      <div className="flex flex-row justify-between items-end text-scndry-txt-clr">
        <div>
          <p>{ipCount} IP adresses</p>
          <p>{upstreamsCount} upstreams</p>
        </div>
        <Status status={status} />
      </div>
    </Card>
  )
}

export default WebsiteCard;