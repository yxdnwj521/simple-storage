const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("SimpleStorage", function () {
    let simpleStorage;
    let accounts;

    beforeEach(async function () {
        // 获取合约工厂和部署者账户
        const SimpleStorage = await ethers.getContractFactory("SimpleStorage");
        accounts = await ethers.getSigners();
        simpleStorage = await SimpleStorage.deploy();
        await simpleStorage.deployed();
    });

    it("Should set and get a value", async function () {
        // 设置值
        const key = 1;
        const value = "Hello, Blockchain!";
        await simpleStorage.set(key, value);

        // 获取值
        const storedValue = await simpleStorage.get(key);
        expect(storedValue).to.equal(value);
    });
});