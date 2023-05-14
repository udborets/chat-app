import { FC } from 'react';

import AuthForm from '@/features/AuthForm/AuthForm';

const AuthPage: FC = () => {
  return (
    <main className='grid place-items-center h-full'>
      <AuthForm />
    </main>
  )
}

export default AuthPage;