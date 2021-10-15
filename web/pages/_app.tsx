import { AppProps } from 'next/app';

import "modern-css-reset/dist/reset.min.css";

function App({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />;
}

export default App;