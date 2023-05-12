import type { AppProps } from 'next/app';
import { FC } from 'react';

import '@/styles/global.scss';

const App: FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <>
      <Component {...pageProps} />
    </>
  )
}

export default App;