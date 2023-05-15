import type { AppProps } from 'next/app';
import { FC } from 'react';
import { Provider } from 'react-redux';

import { store } from '@/store';
import '@/styles/global.scss';

const App: FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  )
}

export default App;