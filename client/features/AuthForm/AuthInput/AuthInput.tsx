import { FC } from "react";

import { AuthInputProps } from "./models";

const AuthInput: FC<AuthInputProps> = ({ name, errors, register, options, labelText, type, placeholder, errorMessage }) => {
  return (
    <label
      htmlFor="phone"
      className="inputLabel"
    >
      {labelText}
      <input
        type={type}
        accept={type === 'file' ? 'image/*' : '*'}
        className="input"
        placeholder={placeholder}
        {...register(name, options)}
      />
      <span className="text-red-600 min-h-fit text-left text-[0.8rem]">
        {errors[name]
          ? (
            errors[name]?.type === 'pattern'
              ? errorMessage
              : errors[name]?.message?.toString())
          : ''}
      </span>
    </label>
  )
}

export default AuthInput