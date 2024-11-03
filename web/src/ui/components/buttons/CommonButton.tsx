import { Path } from "react-hook-form";
import { CommonButtonTypes } from "../../../types/commonButtonTypes";

export const CommonButton = ({
  isSubmit,
  type,
  onClick,
  children,
}: {
  isSubmit: boolean;
  type: Path<CommonButtonTypes>;
  onClick: () => void;
  children: any;
}) => {
  let additionalClasses = "";
  if (type === "blueBgWhiteText") {
    additionalClasses = "bg-main-clr text-white";
  } else if (type === "redBgWhiteText") {
    additionalClasses = "";
  } else if (type === "transparentBgMainText") {
    additionalClasses = "";
  }
  return (
    <button
      className={`w-[440px] h-11 rounded-md ${additionalClasses}`}
      type={isSubmit ? "submit" : "button"}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
