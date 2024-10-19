import "./Logo.css";
import { Pocket } from "react-feather";

function Logo() {
  return (
    <div className="Logo">
      <Pocket size="3rem" stroke="var(--main-color)" strokeWidth={3} />
      <h1><span className="highlight">nginx</span> admin panel</h1>
    </div>
  )
}

export default Logo;