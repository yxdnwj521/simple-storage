require('@nomiclabs/hardhat-waffle');

module.exports = {
  solidity: "0.8.0",
  networks: {
    ganache: {
      url: "http://localhost:7545",
      accounts: [
        // 在这里添加 Ganache 提供的私钥
        "fe1223aa919374b26f4a1d0e2c9caf6e92ac71bb3e7b704f74612018f39214cc",
        // 可以添加更多私钥
      ],
    },
  },
};