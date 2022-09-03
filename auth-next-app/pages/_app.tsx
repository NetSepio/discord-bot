import "../styles/globals.css";
import type { AppProps } from "next/app";
import { WalletProvider } from "../src/contexts/WalletContext";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <WalletProvider>
      <Component {...pageProps} />
    </WalletProvider>
  );
}

export default MyApp;
