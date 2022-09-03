import type { NextPage } from "next";
import { useContext } from "react";
import { WalletContext } from "../src/contexts/WalletContext";
import { Axios } from "axios";
import {
  ApiResponse,
  GetFlowIdPayload,
  PostAuthenticatePayload,
  PostAuthenticateRequest,
} from "./types";
const Home: NextPage = () => {
  const walletContext = useContext(WalletContext);
  const verifyWithDiscord = async () => {
    const axios = new Axios({
      baseURL: process.env.NEXT_PUBLIC_GATEWAY_API,
      transformRequest: (req) => JSON.stringify(req),
      transformResponse: (data) => JSON.parse(data),
    });
    try {
      const flowIdRes = await axios.get<ApiResponse<GetFlowIdPayload>>(
        `flowid`,
        {
          params: {
            walletAddress: walletContext.walletAddress,
          },
        }
      );

      if (flowIdRes.status != 200) {
        throw new Error("failed to call flowid api");
      }

      const { eula, flowId } = flowIdRes.data.payload;
      const message = eula + flowId;
      const signature = await walletContext.web3Provider
        ?.getSigner()
        .signMessage(message);
      if (!signature) {
        alert("Failed to get signature");
        return;
      }
      const authBody: PostAuthenticateRequest = {
        flowId,
        signature,
      };
      const authRes = await axios.post<ApiResponse<PostAuthenticatePayload>>(
        `authenticate`,
        authBody
      );
      if (authRes.status != 200) {
        throw new Error("failed to call authenticate api");
      }
    } catch (error) {
      console.log(error);
      alert("Failed to verify, try again later");
    }
  };
  return (
    <div className="flex justify-center items-center h-screen">
      {walletContext.walletAddress ? (
        <button
          className="bg-blue-600 font-bold rounded-lg p-3 text-2xl"
          onClick={verifyWithDiscord}
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
