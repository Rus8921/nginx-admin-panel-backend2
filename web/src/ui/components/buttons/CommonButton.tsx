import { Path } from "react-hook-form";
import { CommonButtonTypes } from "../../../types/commonButtonTypes";

export const CommonButton = ({
  isSubmit = false,
  type,
  onClick,
  children,
}: {
  isSubmit?: boolean;
  type?: Path<CommonButtonTypes>;
  onClick?: () => void;
  children?: any;
}) => {
  let additionalClasses = "";
  if (type === "blueBgWhiteText") {
    additionalClasses = "bg-main-clr text-white";
  } else if (type === "redBgWhiteText") {
    additionalClasses = "border-red text-red hover:bg-red hover:text-white active:border-scndry-txt-clr";
  } else if (type === "transparentBgMainText") {
    additionalClasses = "";
  }
  return (
    <button
      className={`w-[440px] h-11 border rounded-md flex justify-center gap-4 items-center text-md/10 tracking-normal uppercase whitespace-nowrap transition-colors ${additionalClasses}`}
      type={isSubmit ? "submit" : "button"}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
