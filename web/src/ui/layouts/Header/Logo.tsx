import { Pocket } from "react-feather";

function Logo() {
  return (
    <div className="flex items-center gap-5">
      <Pocket className="mt-3 text-main-clr" size="3rem" strokeWidth={3} />
      <h1><span className="text-main-clr">nginx</span> admin panel</h1>
    </div>
  )
}

export default Logo;