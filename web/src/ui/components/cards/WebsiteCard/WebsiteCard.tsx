import "./WebsiteCard.css"
import Card from "../Card/Card";
import Status from "../../text/Status/Status";
import { WebsiteInterface } from "../../../../types";

function WebsiteCard({ data }: { data: WebsiteInterface }) {
  let { name, url, ipCount, upstreamsCount, status } = data;

  return (
    <Card>
      <div className="WebsiteCard">
        <div>
          <h2>{name}</h2>
          <hr className="light-horizontal-line" />
          <p className="info-text">{url}</p>
        </div>
        <div className="stats">
          <div>
            <p className="info-text">{ipCount} IP adresses</p>
            <p className="info-text">{upstreamsCount} upstreams</p>
          </div>
          <Status status={status} />
        </div>
      </div>
    </Card>
  )
}

export default WebsiteCard;