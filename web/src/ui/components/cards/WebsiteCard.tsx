import Card from "./Card";
import Status from "../text/Status";
import { WebsiteInterface } from "../../../types";
import ContentBreak from "../sections/ContentBreak";
import { ComponentProps } from "react";

export default function WebsiteCard({ data, ...rest }: ComponentProps<"div"> & { data: WebsiteInterface, isClickable: boolean }) {
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
  );
};

export const WebsiteCardInsideServerPage = ({
  data,
}: {
  data: WebsiteInterface;
}) => {
  let { name, url, status } = data;

  return (
    <div className="w-full p-8 bg-white border rounded-2xl border-scndry-clr active:shadow-none">
      <div className="w-full flex flex-row gap-20 justify-between items-center">
        <div className="flex flex-col gap-1">
          <p className="text-md text-scndry-clr font-semibold">{name}</p>
          <p className="text-scndry-txt-clr">{url}</p>
        </div>
        <Status status={status} />
      </div>
    </div>
  );
};