require("@nomicfoundation/hardhat-toolbox");

// import { ProxyAgent, setGlobalDispatcher } from "undici"
const ProxyAgent = require("undici").ProxyAgent
const setGlobalDispatcher = require("undici").setGlobalDispatcher
const proxyAgent = new ProxyAgent('http://127.0.0.1:7890'); // change to yours
setGlobalDispatcher(proxyAgent);

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.18",
  networks: {
    goerli: {
      url: `http endpoint`,
      accounts: ["privatekey"],
    },
  },
  etherscan: {
    apiKey: "",
  },
};
