import Card from "./Card";
import Status from "../text/Status";
import { WebsiteInterface } from "../../../types";

function WebsiteCard({ data }: { data: WebsiteInterface }) {
  let { name, url, ipCount, upstreamsCount, status } = data;

  return (
    <Card>
      <div className="w-full flex flex-col gap-20">
        <div className="flex flex-col gap-4">
          <h2>{name}</h2>
          <hr className="light-horizontal-line" />
          <p className="text-scndry-txt-clr">{url}</p>
        </div>
        <div className="flex flex-row justify-between items-end text-scndry-txt-clr">
          <div>
            <p>{ipCount} IP adresses</p>
            <p>{upstreamsCount} upstreams</p>
          </div>
          <Status status={status} />
        </div>
      </div>
    </Card>
  )
}

export default WebsiteCard;