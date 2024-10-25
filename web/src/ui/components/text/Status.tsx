interface StatusProps {
  status: string;
}

function Status({ status }: StatusProps) {
  return (
    <div className="flex items-baseline gap-2 text-scndry-txt-clr">
      <span>{status}</span>
      <div className={"w-3 h-3 rounded-full" + (status === "active" ? " bg-green" : " bg-red")}></div>
    </div>
  )
}

export default Status;