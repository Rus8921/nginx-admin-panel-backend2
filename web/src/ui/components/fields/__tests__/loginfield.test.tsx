import renderer from "react-test-renderer"
import { LoginPageField } from "../LoginPageField";
import { SubmitHandler, useForm } from "react-hook-form";
import { LoginInputProps } from "../../../../types/loginInputProps";
import { User, Lock, Eye, EyeOff } from "react-feather";


function LoginForm({isLogin, isPassword, isHidden}:{isLogin:boolean, isPassword:boolean, isHidden?:boolean}) {
  const { register, handleSubmit } = useForm<LoginInputProps>();
  const onSubmit: SubmitHandler<LoginInputProps> = (data) => console.log(data);

  return (<form onSubmit={handleSubmit(onSubmit)}>
    {isLogin && (<LoginPageField
      placeholder="Тестовый Логин"
      type="login"
      register={register}
      required={true}
      IconComponent={User} />)}

    {isPassword && (<LoginPageField
      placeholder="Пароль"
      type="password"
      register={register}
      required={true}
      IconComponent={Lock}
      PasswordIconComponent={isHidden? EyeOff : Eye}
      showPassword={!isHidden}
      setShowPassword={() => { }} />)}
  </form>);
}

describe("<LoginPageField />", () => {

  it("renders login filed correctly", () => {
    const tree = renderer
      .create(<LoginForm isLogin={true} isPassword={false}/>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders password revealed filed correctly", () => {
    const tree = renderer
      .create(<LoginForm isLogin={false} isPassword={true} isHidden={false}/>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders password hidden filed correctly", () => {
    const tree = renderer
      .create(<LoginForm isLogin={false} isPassword={true} isHidden={true}/>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders login fileds combination correctly", () => {
    const tree = renderer
      .create(<LoginForm isLogin={true} isPassword={true} isHidden={true}/>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
