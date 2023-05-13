import { FC } from 'react';

import AuthForm from '@/features/AuthForm/AuthForm';

const AuthPage: FC = () => {
  return (
    <main className='w-full h-full'>
      <AuthForm />
    </main>
  )
}

export default AuthPage;