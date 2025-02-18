const { ethers } = require("ethers");

async function main() {
  // 连接到本地 Hardhat 网络
  const provider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:8545");
  const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

  // 合约地址（替换为实际的合约地址）
  const contractAddress = "0x1234567890123456789012345678901234567890";
  const abi = [
    "function set(uint256, string) public",
    "function get(uint256) public view returns (string memory)"
  ];

  // 创建合约实例
  const contract = new ethers.Contract(contractAddress, abi, wallet);

  // 设置值
  const key = 1;
  const value = "Hello, Blockchain!";
  const setTx = await contract.set(key, value);
  await setTx.wait();
  console.log("Value set successfully!");

  // 获取值
  const storedValue = await contract.get(key);
  console.log("Stored value:", storedValue);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });