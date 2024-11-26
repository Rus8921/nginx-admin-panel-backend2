import React from "react";
import { Icon } from "react-feather";
import { FieldPath, FieldValues, useController, UseControllerProps } from "react-hook-form";

function Input<T extends FieldValues>(props: React.ComponentProps<"input"> & UseControllerProps<T, FieldPath<T>> & {
  Icon: Icon;
}) {
  const {name, rules, control, Icon, ...rest} = props
  const { field } = useController({
    name, 
    rules,
    control
  })

  return (
    <div className="flex flex-row items-center w-[440px] relative">
      <Icon className=" text-main-clr stroke-1 absolute ml-5" />
      <input
        className={"pl-16 w-full border border-solid rounded-md border-scndry-txt-clr px-2 py-2.5 focus:border-main-clr focus-visible:border-main-clr outline-none appearance-none "}
        {...field}
        {...rest}
        // onChange={(e: React.FormEvent<HTMLInputElement>) => {
        //   const newValue = e.currentTarget.value;
        //   onChange(newValue);
        // }} 
      />
    </div>
  );
};

export default Input;