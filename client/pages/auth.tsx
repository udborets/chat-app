import AuthForm from '@/features/AuthForm/AuthForm'
import { FC } from 'react'

const AuthPage: FC = () => {
  return (
    <main className='w-full h-full'>
      <AuthForm />
    </main>
  )
}

export default AuthPage;