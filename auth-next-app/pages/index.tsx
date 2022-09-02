import type { NextPage } from "next";
import { useContext } from "react";
import { WalletContext, WalletProvider } from "../src/contexts/WalletContext";

const Home: NextPage = () => {
  const walletContext = useContext(WalletContext);
  return (
    <div className="flex justify-center items-center h-screen">
      {walletContext.walletAddress ? (
        <button
          className="bg-blue-600 font-bold rounded-lg p-3 text-2xl"
          onClick={walletContext.getWeb3Provider}
        >
          Verify with discord
        </button>
      ) : (
        <button
          className="bg-green-500 font-bold rounded-lg p-3 text-2xl"
          onClick={walletContext.getWeb3Provider}
        >
          Connect Wallet
        </button>
      )}
    </div>
  );
};

export default Home;
