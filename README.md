# 简单存储项目

## 项目简介

这是一个基于以太坊区块链和Solidity智能合约的简单存储系统，能够通过区块链技术安全地存储和检索数据。项目包含智能合约和Go语言实现的HTTP接口服务。

## 主要功能

- 通过智能合约存储键值对数据
- 通过HTTP接口查询存储的数据
- 支持通过交易哈希查询历史数据

## 环境要求

- Node.js 16.x或更高版本
- Go 1.24或更高版本
- Ganache或Hardhat本地以太坊网络

## 安装步骤

### Windows环境

1. 安装Git:

```powershell
winget install --id=Git.Git

2. 安装Node.js:

```powershell
winget install --id=OpenJS.NodeJS

3. 安装项目依赖:

```powershell
npm install
```

### Linux环境

1. 安装基础工具:

```bash
sudo apt update
sudo apt install -y git curl wget

2. 安装Node.js:

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
source ~/.bashrc
nvm install node

3. 安装项目依赖:

```bash
npm install
```

## 配置环境

1. 创建.env文件并添加私钥:

```bash
PRIVATE_KEY=YOUR_PRIVATE_KEY

2. Windows设置环境变量:

```powershell
$env:PRIVATE_KEY = "YOUR_PRIVATE_KEY"

3. Linux设置环境变量:

```bash
export PRIVATE_KEY=YOUR_PRIVATE_KEY
```

## 使用说明

### 启动本地以太坊网络

```bash
npx hardhat node
```

### 编译合约

```bash
npx hardhat compile
```

### 部署合约

```bash
npx hardhat run scripts/deploy.js --network local
```

### 启动Go服务

```bash
go run main.go
```

### API接口

- 设置数据: `POST /set?key=<key>&value=<value>`
- 获取数据: `GET /get?key=<key>`
- 通过交易哈希获取数据: `GET /getByTxHash?txHash=<txHash>`

## 项目结构

```bash
contracts/    # 智能合约文件
scripts/     # 部署脚本
test/        # 测试脚本
main.go      # Go服务主程序

## 常见问题

### 无法连接到本地节点

- 确保已启动Hardhat本地网络
- 检查8545端口是否被占用

### 缺少依赖项

```bash
npm cache clean --force
rm -rf node_modules
npm install
```

### 私钥错误

- 确保.env文件中私钥格式正确
- 检查私钥是否包含特殊字符

### 测试失败

- 检查合约和测试代码是否有语法错误
- 确认测试数据正确性
