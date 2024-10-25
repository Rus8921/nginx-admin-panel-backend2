import LangSwitch from "./LangSwitch";
import Logo from "./Logo";

function Header() {
  return (
    <div className="w-full h-28 pt-0 pb-1 px-[4.5rem] sticky top-0 flex justify-between items-center gap-8 bg-white border-b border-b-scndry-txt-clr shadow-b-xs">
      <Logo />
      <LangSwitch />
    </div>
  );
}

export default Header;