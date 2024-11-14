import { useContext } from "react";
import WebsiteConfigContext from "../WebsiteConfigContext";
import Section from "../../../../ui/components/sections/Section";
import FlatTable, { TableData } from "../../../../ui/components/tables/FlatTable";
import TextButton from "../../../../ui/components/buttons/TextButton";

function SSLCertificates() {
  const { configs } = useContext(WebsiteConfigContext);
  let data: TableData[] = !!configs ? [{
    items: configs.sslCertificates.map((certificate) => {
      return {
        id: certificate.id,
        crtFile: certificate.crtFile,
        keyFile: certificate.keyFile,
        expireAt: getExpirationTimeLeft(certificate.expirationDate),
        isActive: certificate.isActive ? "active" : (<TextButton isDanger={false} className="underline">activate</TextButton>)
      }
    })
  }] : []

  function getDateDifferenceDays(date1: Date, date2: Date) {
    return (date1.getTime() - date2.getTime()) / (1000 * 60 * 60 * 24)
  }

  function getExpirationTimeLeft(expDateString: string) {
    let expDate = new Date(expDateString);
    let dateDiff = getDateDifferenceDays(expDate, new Date())

    let isDanger = dateDiff < 31

    let text = dateDiff < 0 ?
      `expired` : `expires ` + (
        dateDiff < 1 ?
          `today` : dateDiff < 31 ?
            `in ${Math.floor(dateDiff)} days` :
            `on ${expDate.toLocaleDateString()}`)

    return isDanger ? (<span className="text-red">{text}</span>) : text
  }

  return !!configs ? (
    <Section title="SSL Certificates">
      <FlatTable data={data} />
    </Section>
  ) : (<></>);
};

export default SSLCertificates;
