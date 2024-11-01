export type LoginInputProps = {
  login: string;
  password: string;
  required: boolean;
};

export type RegistrationInputProps = {
  email: string;
  login: string;
  password: string;
  repeatPassword: string;
  passwordsAreIdentical: boolean;
  required: boolean;
};
