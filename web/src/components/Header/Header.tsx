import "./Header.css";
import LangSwitch from "./LangSwitch/LangSwitch";
import Logo from "./Logo/Logo";

function Header() {
  return (
    <div className="Header">
      <Logo />
      <LangSwitch />
    </div>
  );
}

export default Header;