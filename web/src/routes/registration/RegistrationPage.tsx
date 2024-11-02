/* eslint-disable react-hooks/rules-of-hooks */
import { useState } from "react";
import { Link, redirect, useNavigate } from "react-router-dom";
import { useUserStore } from "../../stores/userStore";
import { RegistrationPageField } from "../../ui/components/fields/RegistrationPageField";
import { SubmitHandler, useForm } from "react-hook-form";
import { RegistrationInputProps } from "../../types/loginInputProps";
import { Lock, User, Eye, EyeOff, Mail, UserPlus } from "react-feather";
import { CommonButton } from "../../ui/components/buttons/CommonButton";

export async function loader() {
  const isLoggedIn = useUserStore.getState().isLoggedIn;
  if (isLoggedIn()) {
    return redirect("/");
  }
  return null;
}

export const RegistrationPage = () => {
  const {
    register,
    watch,
    handleSubmit,
    formState: { errors },
  } = useForm<RegistrationInputProps>();
  const onSubmit: SubmitHandler<RegistrationInputProps> = (data) => {
    if (passwordsIdentical) {
      useUserStore.getState().login({
        user: {
          email: "example@example.com",
          login: "login",
          token: "token",
          refreshToken: "refreshToken",
          tokenExpiresMilliseconds: 1000,
          role: 1,
        },
      });
      navigate("/");
    }
  };
  const [showFirstPassword, setShowFirstPassword] = useState(false);
  const [showSecondPassword, setShowSecondPassword] = useState(false);
  const [firstPassword, setFirstPassword] = useState("");
  const [secondPassword, setSecondPassword] = useState("");
  const validateTwoPasswords = () => {
    setPasswordsIdentical(false);
    if (firstPassword === "" && secondPassword === "") {
      setPasswordsIdentical(true);
    } else if (firstPassword !== "" && secondPassword !== "") {
      if (firstPassword === secondPassword) {
        setPasswordsIdentical(true);
      } else {
        setPasswordsIdentical(false);
      }
    }
  };
  const [passwordsIdentical, setPasswordsIdentical] = useState(true);

  const navigate = useNavigate();

  return (
    <div className="h-screen w-screen bg-bg-clr flex flex-col justify-center items-center">
      <div className="w-[630px] h-[620px] rounded-2xl outline outline-1 outline-main-clr bg-white flex flex-col justify-center items-center">
        <UserPlus className="h-32 w-32 stroke-1 text-scndry-txt-clr mb-3" />
        <form
          className="flex flex-col gap-4 p-0"
          onSubmit={handleSubmit(onSubmit)}
        >
          <RegistrationPageField
            placeholder="Введите почту"
            type="email"
            register={register}
            required={true}
            IconComponent={Mail}
            onChange={() => {}}
          />
          <RegistrationPageField
            placeholder="Введите логин"
            type="login"
            register={register}
            required={true}
            IconComponent={User}
            onChange={() => {}}
          />
          <RegistrationPageField
            placeholder="Введите пароль"
            type="password"
            register={register}
            required={true}
            IconComponent={Lock}
            PasswordIconComponent={showFirstPassword ? Eye : EyeOff}
            showPassword={showFirstPassword}
            setShowPassword={() => {
              setShowFirstPassword((prev) => !prev);
            }}
            onChange={setFirstPassword}
          />
          <RegistrationPageField
            placeholder="Повторите пароль"
            type="repeatPassword"
            register={register}
            required={true}
            IconComponent={Lock}
            PasswordIconComponent={showSecondPassword ? Eye : EyeOff}
            showPassword={showSecondPassword}
            setShowPassword={() => {
              setShowSecondPassword((prev) => !prev);
            }}
            onChange={setSecondPassword}
          />
          {(errors.email?.type === "required" ||
            errors.login?.type === "required" ||
            errors.password?.type === "required" ||
            errors.repeatPassword?.type === "required" ||
            !passwordsIdentical) && (
            <div className={`flex flex-col`}>
              {errors.email?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Заполните поле для почты
                </p>
              )}
              {errors.login?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Заполните поле для логина
                </p>
              )}
              {errors.password?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Заполните поле для пароля
                </p>
              )}
              {errors.repeatPassword?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Повторите пароль
                </p>
              )}
              {!passwordsIdentical && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Пароли должны совпадать
                </p>
              )}
            </div>
          )}
          <CommonButton
            buttonText={"ЗАРЕГИСТРИРОВАТЬСЯ"}
            isSubmit={true}
            onClick={validateTwoPasswords}
            type="blueBgWhiteText"
          />
        </form>
        <Link
          className="mt-2 text-xs underline-offset-4 text-scndry-clr hover:underline"
          to={"/login"}
        >
          Уже есть аккаунт?
        </Link>
      </div>
    </div>
  );
};
