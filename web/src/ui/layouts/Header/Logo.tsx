import { Pocket } from "react-feather";

function Logo() {
  return (
    <div className="flex items-center gap-5">
      <Pocket className="mt-3" size="3rem" stroke="var(--main-color)" strokeWidth={3} />
      <h1><span className="text-main-clr">nginx</span> admin panel</h1>
    </div>
  )
}

export default Logo;