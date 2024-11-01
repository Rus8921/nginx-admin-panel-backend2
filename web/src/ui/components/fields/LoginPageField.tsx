import React from "react";
import { Path, SubmitHandler, useForm, UseFormRegister } from "react-hook-form";
import { LoginInputProps } from "../../../types/loginInputProps";
import { Icon } from "react-feather";

export const LoginPageField = ({
  placeholder,
  type,
  register,
  required,
  IconComponent,
  PasswordIconComponent,
  showPassword,
  setShowPassword,
}: {
  placeholder: string;
  type: Path<LoginInputProps>;
  register: UseFormRegister<LoginInputProps>;
  required: boolean;
  IconComponent: Icon;
  PasswordIconComponent?: Icon;
  showPassword?: boolean;
  setShowPassword?: () => void;
}) => {
  return (
    <div className="flex flex-row items-center w-[440px] relative">
      <IconComponent className=" text-main-clr stroke-1 absolute ml-5" />
      <input
        className="pl-16 w-full border border-solid rounded-md border-scndry-txt-clr p-2 focus:border-main-clr focus-visible:border-main-clr outline-none"
        placeholder={placeholder}
        {...register(type)}
        required={required}
        type={
          type === "password" ? (showPassword ? "text" : "password") : "text"
        }
      ></input>
      {PasswordIconComponent ? (
        <PasswordIconComponent
          className=" text-main-clr stroke-1 absolute right-0 mr-5 hover:cursor-pointer"
          onClick={setShowPassword}
        />
      ) : (
        <></>
      )}
    </div>
  );
};
