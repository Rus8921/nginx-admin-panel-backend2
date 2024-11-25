import React from "react";
import { Icon } from "react-feather";

function Input({
  placeholder,
  Icon,
  onChange,
  ...rest
}: React.ComponentProps<"input"> & {
  Icon: Icon;
}) {
  return (
    <div className="flex flex-row items-center w-[440px] relative">
      <Icon className=" text-main-clr stroke-1 absolute ml-5" />
      <input
        className="pl-16 w-full border border-solid rounded-md border-scndry-txt-clr px-2 py-2.5 focus:border-main-clr focus-visible:border-main-clr outline-none appearance-none"
        placeholder={placeholder}
        // onChange={(e: React.FormEvent<HTMLInputElement>) => {
        //   const newValue = e.currentTarget.value;
        //   onChange(newValue);
        // }} 
        {...rest}
      />
    </div>
  );
};

export default Input;