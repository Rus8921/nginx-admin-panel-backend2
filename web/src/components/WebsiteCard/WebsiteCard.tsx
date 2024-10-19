import Card from "../Card/Card";
import Status from "../Status/Status";
import "./WebsiteCard.css"

function WebsiteCard() {
  let { name, url, ipCount, upstreamsCount, status } = { name: "website name", url: "some.url.com", ipCount: 3, upstreamsCount: 3, status: "active" };

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