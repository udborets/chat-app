import { useEffect, useState } from "react";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";

import AuthInput from "./AuthInput/AuthInput";
import { AuthTypes } from "./models";

const AuthForm = () => {
  const [currentAuthType, setCurrentAuthType] = useState<AuthTypes>(AuthTypes.REGISTRATION);
  const {
    register,
    handleSubmit,
    formState,
    reset,

  } = useForm({
    defaultValues: {
      phone: '',
      email: '',
      password: '',
      name: '',
      avatar: null,
    }
  });
  const [isByPhone, setIsByPhone] = useState<boolean>(false);
  const [isByEmail, setIsByEmail] = useState<boolean>(true);
  const submit: SubmitHandler<FieldValues> = (data: FieldValues) => {
    const authParams: Record<string, string> = {};
    if (currentAuthType === AuthTypes.REGISTRATION) {
      authParams.avatar = data.avatar;
      authParams.name = data.name;
    }
    authParams.email = isByEmail ? data.email : null;
    authParams.phone = isByPhone ? data.phone : null;
    authParams.password = data.password;
    console.log(authParams)
    reset();
  }
  useEffect(() => {
    setIsByEmail(true);
    setIsByPhone(false);
  }, [currentAuthType])

  return (
    <>
      <form
        onSubmit={handleSubmit(submit)}
        className="p-2 pc:p-4 flex flex-col gap-3 max-w-[400px] rounded-[10px] "
      >
        {currentAuthType === AuthTypes.REGISTRATION
          ? <>
            <AuthInput
              errors={formState.errors}
              name="name"
              register={register}
              options={{
                required: 'Name is required',
                pattern: /[a-zA-Z]{2,20}/,
              }}
              labelText="Name"
              errorMessage="Name should include from 2 to 20 symbols"
              placeholder="Enter your name..."
              type="text"
            />
            <AuthInput
              errors={formState.errors}
              name="avatar"
              register={register}
              errorMessage="Enter valid avatar"
              options={{}}
              labelText="Avatar"
              placeholder=""
              type="file"
            />
          </>
          : ''}
        {isByEmail
          ? <AuthInput
            errors={formState.errors}
            type="text"
            name="email"
            register={register}
            labelText="Email"
            errorMessage="Enter valid email"
            placeholder="Enter email..."
            options={{
              required: 'Email is required',
              pattern: /\w{4,15}@\w{4,8}\.\w{2,5}/,
            }}
          />
          : ''}
        {isByPhone
          ? <AuthInput
            errors={formState.errors}
            name="phone"
            register={register}
            options={{
              required: 'Phone is required',
            }}
            labelText="Phone"
            placeholder="Enter your phone number..."
            errorMessage="Phone should include from "
            type='number'
          />
          : ''}
        <AuthInput
          errors={formState.errors}
          type="password"
          name="password"
          register={register}
          labelText="Password"
          placeholder="Enter password..."
          errorMessage="Password should include from 8 to 30 symbols"
          options={{
            required: 'Password is required',
            pattern: /[\w*!@#$%^&?]{8,30}/,
          }}
        />
        {currentAuthType === AuthTypes.REGISTRATION
          ? <>
            <button
              type="button"
              onClick={() => setIsByEmail(prev => !prev)}
            >
              {!isByEmail
                ? "Add email"
                : 'Remove email'}
            </button>
            <button
              type="button"
              onClick={() => setIsByPhone(prev => !prev)}>
              {!isByPhone
                ? "Add phone"
                : 'Remove phone'}
            </button>
          </>
          : <>
            <button
              type="button"
              onClick={() => { setIsByEmail(prev => !prev); setIsByPhone(prev => !prev) }}
            >
              {isByEmail
                ? "Login with phone"
                : 'Login with email'}
            </button>
          </>}
        {(!isByEmail && !isByPhone)
          ? <span className="text-red-600">
            You have to enter email or phone
          </span>
          : ''}
        <input
          type="submit"
          disabled={(!isByEmail && !isByPhone) || !formState.isValid}
          className="p-2 rounded-button outline disabled:bg-slate-600 duration-200 transition-all"
          value={currentAuthType === AuthTypes.AUTHORIZATION
            ? 'Log in!'
            : 'Register!'}
        />
        <button
          type="button"
          className="text-[0.8rem]"
          onClick={() => setCurrentAuthType(prev => prev === AuthTypes.AUTHORIZATION ? AuthTypes.REGISTRATION : AuthTypes.AUTHORIZATION)}
        >
          {currentAuthType === AuthTypes.AUTHORIZATION
            ? `Don't have an account? Create one!`
            : `Already have an account? Log in!`}
        </button>
      </form>
    </>
  )
}

export default AuthForm;