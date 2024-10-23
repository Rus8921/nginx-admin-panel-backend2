import "./Status.css"

interface StatusProps {
  status: string;
}

function Status({ status }: StatusProps) {
  return (
    <div className="Status info-text">
      <span>{status}</span>
      <div className={"light " + status}></div>
    </div>
  )
}

export default Status;