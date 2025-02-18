# 项目名称 simple storage

## 项目简介

这是一个能搭建一个私链，能存储数据和读取数据。通过Solidity 完成数据存储和读取这个功能。通过Solidity进行编写代码。使用Solidity存储进区块链将数据上链，然后就能获取数据了。区块链那边有个接口会来获取数据。

## 原理
该项目实现数据存储和读取的原理基于以太坊区块链和Solidity智能合约技术。
 
以太坊区块链：以太坊是一个开源的有智能合约功能的公共区块链平台。它提供了一个去中心化的分布式账本，由众多节点共同维护。在这个项目中，通过Hardhat启动的本地以太坊网络，模拟了一个私链环境。各节点在这个私链上同步账本数据，确保数据的一致性和不可篡改。当进行数据存储操作时，交易信息会被打包成区块，添加到区块链中；读取数据时，则是从区块链的账本中获取相应信息。
 
Solidity智能合约：Solidity是以太坊平台上用于编写智能合约的编程语言。项目中的智能合约定义了数据存储和读取的逻辑。比如，使用 mapping （类似哈希表）结构来存储数据， mapping(uint => string) private dataStore; 定义了一个以 uint 类型为键、 string 类型为值的存储结构。 storeData 函数用于将数据存储到 dataStore 中， retrieveData 函数用于从 dataStore 中读取数据。这些函数在区块链上以字节码的形式部署和运行，外部通过发送交易来调用这些函数，实现数据的存储和读取操作。
 
交易与交互：在以太坊网络中，数据的存储和读取操作都是通过交易来完成的。当调用智能合约的 storeData 函数存储数据时，会创建一笔交易，交易中包含调用函数的相关信息（如函数名、参数等）以及发送者的签名。这笔交易被广播到以太坊网络中，由矿工节点进行验证和打包进区块。读取数据时，虽然不需要修改区块链状态，不产生实际的交易费用，但同样是通过发送调用 retrieveData 函数的请求到以太坊网络，从区块链账本中获取数据并返回给调用者。


## 目录

- [环境准备](#环境准备)
    - [Windows](#windows-环境准备)
    - [Linux](#linux-环境准备)
- [克隆项目](#克隆项目)
- [安装依赖](#安装依赖)
    - [Windows](#windows-安装依赖)
    - [Linux](#linux-安装依赖)
- [配置环境变量](#配置环境变量)
    - [Windows](#windows-配置环境变量)
    - [Linux](#linux-配置环境变量)
- [启动项目](#启动项目)
    - [Windows](#windows-启动项目)
    - [Linux](#linux-启动项目)
- [使用项目](#使用项目)
- [常见问题及解决方法](#常见问题及解决方法)

## 环境准备

### Windows 环境准备

在开始之前，请确保你的系统满足以下要求：

- **操作系统**：Windows 10 或更高版本
- **终端模拟器**：建议使用 PowerShell 或 Windows Terminal

#### 安装必要的软件包

打开 PowerShell 并运行以下命令来安装必要的软件包：

```powershell
# 安装 Git
winget install --id=Git.Git

# 安装 Node.js
winget install --id=OpenJS.NodeJS

# 安装 curl 和 wget（可选）
winget install --id=GNU.Wget
winget install --id=curl.curl
```

### Linux 环境准备

在开始之前，请确保你的系统满足以下要求：

- **操作系统**：Linux（支持的发行版：Ubuntu, CentOS, Debian等）
- **终端模拟器**：任何标准的终端模拟器（如 `gnome-terminal`, `xterm`, `konsole` 等）

#### 安装必要的软件包

打开终端并运行以下命令来安装必要的软件包：

##### Ubuntu/Debian

```sh
sudo apt update
sudo apt install -y git curl wget
```

##### CentOS

```sh
sudo yum update
sudo yum install -y git curl wget
```

## 克隆项目

首先，你需要克隆这个项目到你的本地机器上。打开终端或 PowerShell 并运行以下命令：

```sh
git clone https://github.com/yxdnwj521/simple-storage.git
cd yourproject
```

## 安装依赖

进入项目目录后，安装项目所需的依赖包。运行以下命令：

### Windows 安装依赖

```powershell
npm install
```

如果你还没有安装 Node.js 和 npm，请参考以下步骤进行安装：

#### 安装 Node.js 和 npm

你可以使用 `nvm`（Node Version Manager）来安装 Node.js 和 npm。运行以下命令来安装 `nvm` 和 Node.js：

```powershell
# 下载 nvm-windows
Invoke-WebRequest -Uri https://github.com/coreybutler/nvm-windows/releases/download/1.1.7/nvm-setup.zip -OutFile nvm-setup.zip
Expand-Archive -Path nvm-setup.zip -DestinationPath $env:TEMP\nvm
$env:Path += ";$env:TEMP\nvm"

# 安装 nvm-windows
& "$env:TEMP\nvm\nvm.exe" /silent

# 重启 PowerShell
exit
```

重新打开 PowerShell 并运行以下命令来安装 Node.js：

```powershell
nvm install node
```

### Linux 安装依赖

```sh
npm install
```

如果你还没有安装 Node.js 和 npm，请参考以下步骤进行安装：

#### 安装 Node.js 和 npm

你可以使用 `nvm`（Node Version Manager）来安装 Node.js 和 npm。运行以下命令来安装 `nvm` 和 Node.js：

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
source ~/.bashrc
nvm install node
```

## 配置环境变量

为了安全起见，我们将私钥和其他敏感信息存储在 `.env` 文件中。请在项目根目录下创建一个 `.env` 文件，并添加以下内容：

```plaintext
PRIVATE_KEY=YOUR_PRIVATE_KEY
```

### 获取私钥

- 如果你还没有以太坊钱包，可以使用 MetaMask 创建一个。
- 在 MetaMask 中导出你的私钥，并将其填入 `.env` 文件中的 `PRIVATE_KEY`。

### Windows 配置环境变量

在 PowerShell 中，你可以使用以下命令来设置环境变量：

```powershell
$env:PRIVATE_KEY = "YOUR_PRIVATE_KEY"
```

### Linux 配置环境变量

在终端中，你可以使用以下命令来设置环境变量：

```sh
export PRIVATE_KEY=YOUR_PRIVATE_KEY
```

## 启动项目

### Windows 启动项目

#### 启动 Hardhat 本地网络

在 PowerShell 中运行以下命令来启动 Hardhat 本地网络：

```powershell
npx hardhat node
```

这将启动一个本地以太坊网络，默认监听端口为 8545。

#### 运行测试

在另一个 PowerShell 窗口中，运行以下命令来执行测试：

```powershell
npx hardhat test
```

这将运行 `test/simpleStorage.test.js` 中的所有测试，并输出详细的日志信息。

#### 部署合约到私链

你可以编写一个部署脚本来将合约部署到本地私链上。运行以下命令：

```powershell
npx hardhat run scripts/deploy.js --network local
```

这将输出合约的地址，例如：`SimpleStorage deployed to: 0x1234567890123456789012345678901234567890`

#### 与合约交互

你可以使用 Ethers.js 与已部署的合约进行交互。运行以下命令来设置和获取数据：

```powershell
npx hardhat run scripts/interact.js --network local
```

这将输出以下内容：`Value set successfully! Stored value: Hello, Blockchain!`

### Linux 启动项目

#### 启动 Hardhat 本地网络

在终端中运行以下命令来启动 Hardhat 本地网络：

```sh
npx hardhat node
```

这将启动一个本地以太坊网络，默认监听端口为 8545。

#### 运行测试

在另一个终端窗口中，运行以下命令来执行测试：

```sh
npx hardhat test
```

这将运行 `test/simpleStorage.test.js` 中的所有测试，并输出详细的日志信息。

#### 部署合约到私链

你可以编写一个部署脚本来将合约部署到本地私链上。运行以下命令：

```sh
npx hardhat run scripts/deploy.js --network local
```

这将输出合约的地址，例如：`SimpleStorage deployed to: 0x1234567890123456789012345678901234567890`

#### 与合约交互

你可以使用 Ethers.js 与已部署的合约进行交互。运行以下命令来设置和获取数据：

```sh
npx hardhat run scripts/interact.js --network local
```

这将输出以下内容：`Value set successfully! Stored value: Hello, Blockchain!`

## 使用项目

### 项目结构

- `contracts/`：存放智能合约文件
- `scripts/`：存放部署和交互脚本
- `test/`：存放测试脚本
- `hardhat.config.js`：Hardhat 配置文件
- `.env`：环境变量文件

### 常用命令

- **编译合约**：
  ```sh
  npx hardhat compile
  ```

- **运行测试**：
  ```sh
  npx hardhat test
  ```

- **部署合约**：
  ```sh
  npx hardhat run scripts/deploy.js --network local
  ```

- **与合约交互**：
  ```sh
  npx hardhat run scripts/interact.js --network local
  ```

## 常见问题及解决方法

### 1. 无法连接到本地节点

**问题描述**：运行 `npx hardhat node` 或其他命令时，提示无法连接到本地节点。

**解决方法**：
- 确保你已经启动了 Hardhat 本地网络。可以在一个新的终端窗口中运行 `npx hardhat node`。
- 检查是否有其他进程占用了 8545 端口。可以使用以下命令查看占用该端口的进程：
    - **Windows**：
      ```powershell
      Get-Process -Id (Get-NetTCPConnection -LocalPort 8545).OwningProcess
      ```
    - **Linux**：
      ```sh
      sudo lsof -i :8545
      ```
- 如果有其他进程占用该端口，可以终止该进程或更改 Hardhat 的端口号。

### 2. 缺少依赖项

**问题描述**：运行 `npm install` 时，提示缺少某些依赖项。

**解决方法**：
- 确保你的网络连接正常。
- 清理 npm 缓存并重新安装依赖项：
    - **Windows**：
      ```powershell
      npm cache clean --force
      Remove-Item -Recurse -Force node_modules
      npm install
      ```
    - **Linux**：
      ```sh
      npm cache clean --force
      rm -rf node_modules
      npm install
      ```

### 3. 私钥错误

**问题描述**：运行部署或交互脚本时，提示私钥错误。

**解决方法**：
- 确保你在 `.env` 文件中正确填写了私钥。
- 检查私钥是否包含空格或其他特殊字符。
- 确保私钥是有效的以太坊私钥。

### 4. 测试失败

**问题描述**：运行测试时，某些测试失败。

**解决方法**：
- 查看测试输出的详细日志，找到具体的错误信息。
- 确保你的合约代码和测试代码没有语法错误。
- 确保你的测试数据和预期结果是正确的。

祝你好运！希望你能顺利使用这个项目！

感谢你的支持！
