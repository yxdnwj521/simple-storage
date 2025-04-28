const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    // 获取合约工厂
    const SimpleStorage = await ethers.getContractFactory("SimpleStorage");

    // 部署合约
    const simpleStorage = await SimpleStorage.deploy();
    await simpleStorage.deployed();

    // 保存合约地址到文件


    const fs = require('fs');
fs.writeFileSync('contract-address.json', simpleStorage.address);
console.log("SimpleStorage deployed to:", simpleStorage.address);
}

// 运行主函数
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });