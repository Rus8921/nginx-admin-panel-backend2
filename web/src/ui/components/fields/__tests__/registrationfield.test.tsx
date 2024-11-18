import renderer from "react-test-renderer";
import { RegistrationPageField } from "../RegistrationPageField";
import { SubmitHandler, useForm } from "react-hook-form";
import { RegistrationInputProps } from "../../../../types/loginInputProps";
import { User, Lock, Eye, EyeOff } from "react-feather";

function RegistrationForm({
  isLogin,
  isPassword,
  isHidden,
  isRepeat,
}: {
  isLogin: boolean;
  isPassword: boolean;
  isHidden?: boolean;
  isRepeat?: boolean;
}) {
  const { register, handleSubmit } = useForm<RegistrationInputProps>();
  const onSubmit: SubmitHandler<RegistrationInputProps> = (data) =>
    console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {isLogin && (
        <RegistrationPageField
          placeholder="Тестовый Логин"
          type="login"
          register={register}
          required={true}
          IconComponent={User}
          onChange={() => {}}
        />
      )}
      {isPassword && (
        <>
          (
          <RegistrationPageField
            placeholder="Пароль"
            type="password"
            register={register}
            required={true}
            IconComponent={Lock}
            PasswordIconComponent={isHidden ? EyeOff : Eye}
            showPassword={!isHidden}
            setShowPassword={() => {}}
            onChange={() => {}}
          />
          )
          {!isRepeat ?? (
            <RegistrationPageField
              placeholder="Пароль"
              type="password"
              register={register}
              required={true}
              IconComponent={Lock}
              PasswordIconComponent={isHidden ? EyeOff : Eye}
              showPassword={!isHidden}
              setShowPassword={() => {}}
              onChange={() => {}}
            />
          )}
        </>
      )}
    </form>
  );
}

describe("<LoginPageField />", () => {
  it("renders login filed correctly", () => {
    const tree = renderer
      .create(<RegistrationForm isLogin={true} isPassword={false} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders password revealed filed correctly", () => {
    const tree = renderer
      .create(
        <RegistrationForm isLogin={false} isPassword={true} isHidden={false} />
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders password repeat filed correctly", () => {
    const tree = renderer
      .create(
        <RegistrationForm
          isLogin={false}
          isPassword={true}
          isHidden={false}
          isRepeat={true}
        />
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders password hidden filed correctly", () => {
    const tree = renderer
      .create(
        <RegistrationForm isLogin={false} isPassword={true} isHidden={true} />
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders registration fileds combination correctly", () => {
    const tree = renderer
      .create(
        <RegistrationForm
          isLogin={true}
          isPassword={true}
          isHidden={true}
          isRepeat={true}
        />
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
