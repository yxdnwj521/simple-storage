require('dotenv').config();
require('@nomiclabs/hardhat-waffle');

module.exports = {
  solidity: "0.8.0",
  networks: {
    hardhat: {
      // 配置 Hardhat 网络
    },
    rinkeby: {
      url: `https://rinkeby.infura.io/v3/YOUR_INFURA_PROJECT_ID`,
      accounts: [process.env.PRIVATE_KEY]
    }
  }
};
/*
* require('dotenv').config();
require('@nomiclabs/hardhat-waffle');
require('./scripts/compile');

module.exports = {
  solidity: "0.8.0",
  networks: {
    hardhat: {
      // 配置 Hardhat 网络
    },
    local: {
      url: "http://127.0.0.1:8545", // 本地节点地址
      accounts: [process.env.PRIVATE_KEY]
    }
  }
};
* //一般情况下使用不到这串被注释的代码的，如果你打算用scripts/compile.js，可以将上面代码注释并启用这个代码，来进行自定义编译任务
* */