import type { AppProps } from 'next/app';
import Head from 'next/head';
import { FC } from 'react';

import '@/styles/global.scss';

const App: FC<AppProps> = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default App;