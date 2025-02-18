const { ethers } = require("ethers");

async function main() {
    // 连接到本地 Hardhat 网络
    const provider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:8545");
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

    // 获取合约工厂
    const SimpleStorage = await ethers.getContractFactory("SimpleStorage", wallet);

    // 部署合约
    const simpleStorage = await SimpleStorage.deploy();
    await simpleStorage.deployed();

    console.log("SimpleStorage deployed to:", simpleStorage.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });