import { HTMLInputTypeAttribute } from "react";
import { FieldErrors, RegisterOptions, UseFormRegister } from "react-hook-form";

export type AuthInputProps = {
  register: UseFormRegister<any>;
  type: HTMLInputTypeAttribute;
  options?: RegisterOptions;
  errors: FieldErrors<any>;
  name: string;
  placeholder: string;
  labelText: string;
  errorMessage: string;
};
