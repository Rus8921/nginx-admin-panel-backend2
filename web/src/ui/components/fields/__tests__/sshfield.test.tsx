import renderer from "react-test-renderer";
import { SSHInputField } from "../SSHInputField";
import {
  RegisterOptions,
  SubmitHandler,
  useForm,
  UseFormRegisterReturn,
} from "react-hook-form";
import { SshKeyInputProps } from "../../../../routes/servers/ServerConfigPage";
import { Terminal } from "react-feather";

const SSHInputForm = () => {
  const { register, handleSubmit } = useForm<SshKeyInputProps>();
  const onSubmit: SubmitHandler<SshKeyInputProps> = (data) => console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <SSHInputField
        placeholder={"SSH key"}
        type={"sshKey"}
        register={register}
        required={false}
        IconComponent={Terminal}
        onChange={(e: any) => {
          console.log(e);
        }}
      />
    </form>
  );
};

describe("<SSHInputField />", () => {
  it("renders ssh key input filed correctly", () => {
    const tree = renderer.create(<SSHInputForm />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
