import React from "react";
import { Path, SubmitHandler, useForm, UseFormRegister } from "react-hook-form";
import { Icon } from "react-feather";
import { SshKeyInputProps } from "../../../routes/servers/ServerConfigPage";

export const SSHInputField = ({
  placeholder,
  type,
  register,
  required,
  IconComponent,
  onChange,
}: {
  placeholder: string;
  type: Path<SshKeyInputProps>;
  register: UseFormRegister<SshKeyInputProps>;
  required: boolean;
  IconComponent: Icon;
  onChange: any;
}) => {
  return (
    <div className="flex flex-row items-center w-[440px] relative">
      <IconComponent className=" text-main-clr stroke-1 absolute ml-5" />
      <input
        className="pl-16 w-full border border-solid rounded-md border-scndry-txt-clr p-2 focus:border-main-clr focus-visible:border-main-clr outline-none"
        placeholder={placeholder}
        {...register(type, { required: required })}
        onChange={(e: React.FormEvent<HTMLInputElement>) => {
          const newValue = e.currentTarget.value;
          onChange(newValue);
        }}
      ></input>
    </div>
  );
};
