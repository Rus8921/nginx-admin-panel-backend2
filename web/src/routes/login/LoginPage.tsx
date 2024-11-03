/* eslint-disable react-hooks/rules-of-hooks */
import { useState } from "react";
import { Link, redirect, useNavigate } from "react-router-dom";
import { useUserStore } from "../../stores/userStore";
import { LoginPageField } from "../../ui/components/fields/LoginPageField";
import { SubmitHandler, useForm } from "react-hook-form";
import { LoginInputProps } from "../../types/loginInputProps";
import { Lock, User, Eye, EyeOff } from "react-feather";
import { CommonButton } from "../../ui/components/buttons/CommonButton";

export async function loader() {
  const isLoggedIn = useUserStore.getState().isLoggedIn;
  if (isLoggedIn()) {
    return redirect("/");
  }
  return null;
}

export const LoginPage = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginInputProps>();
  const onSubmit: SubmitHandler<LoginInputProps> = (data) => {
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
  };
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();

  return (
    <div className="h-screen w-screen bg-bg-clr flex flex-col justify-center items-center">
      <div className="w-[630px] h-[420px] rounded-2xl outline outline-1 outline-main-clr bg-white flex flex-col justify-center items-center">
        <User className="h-32 w-32 stroke-1 text-scndry-txt-clr mb-3" />
        <form
          className="flex flex-col gap-4 p-0"
          onSubmit={handleSubmit(onSubmit)}
        >
          <LoginPageField
            placeholder="Логин"
            type="login"
            register={register}
            required={true}
            IconComponent={User}
          />
          <LoginPageField
            placeholder="Пароль"
            type="password"
            register={register}
            required={true}
            IconComponent={Lock}
            PasswordIconComponent={showPassword ? Eye : EyeOff}
            showPassword={showPassword}
            setShowPassword={() => {
              setShowPassword((prev) => !prev);
            }}
          />
          {(errors.login?.type === "required" ||
            errors.password?.type === "required") && (
            <div className={`flex flex-col`}>
              {errors.login?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Введите логин
                </p>
              )}
              {errors.password?.type === "required" && (
                <p role="alert" className="text-scndry-clr text-xs ml-2">
                  Введите пароль
                </p>
              )}
            </div>
          )}
          <CommonButton
            isSubmit={true}
            onClick={() => {}}
            type="blueBgWhiteText"
          >
            <span className="font-roboto">ВОЙТИ</span>
          </CommonButton>
        </form>
        <Link
          className="mt-2 text-xs underline-offset-4 text-scndry-clr hover:underline"
          to={"/registration"}
        >
          Создать аккаунт
        </Link>
      </div>
    </div>
  );
};
