### Node.js
- **版本**: 16.x 或更高
    - **下载命令**:
      ```bash
      sudo apt install nodejs
      sudo apt install npm
      npm install
      ```
- **检查是否安装**:
    - 示例:
      ```bash
        node -v
        npm -v
      ```
``
**打开Ganache**:
- 示例:
- 点击QUICKSTART（这是快速部署，如果能看懂英文，可以试试NEW WORKPLACE）
- 点击第一个账户（或者随便）
- 点击右边的钥匙会出现两行奇怪的数字，双击下面那个然后Ctrl+c复制
- 把它放到根目录的config.json和hardhat.config.js的YOUR_PRIVATE_KEY（记得config.json的密钥前面把0x去掉，另一个不用去掉）
- 然后运行
```bash
    npx hardhat compile
   ```
来进行合约编译与部署 
**部署合约到本地网络**:
   ```bash
   npx hardhat run scripts/deploy.js --network ganache
   ```
   输入上面命令的时候会有如下回复：
   Deploying contracts with the account: YOUR_PRIVATE_KEY
   SimpleStorage deployed to: YOUR_ADDRESS
    下面那个是连接账户时的地址请填到main.go的大约在75行的YOUR_ADDRESS那儿。
4. 
## Go 应用程序
Go 应用程序实现了HTTP服务器功能，允许用户通过API接口设置和获取存储于智能合约中的数据。

### 运行Go服务
首先需要加载配置文件（例如：config.json），其中包含私钥等敏感信息。
然后启动Go HTTP服务器：
```bash
go run main.go
```
现在，您可以访问以下API端点：

- **设置键值对**:
    - 方法: `POST /set?key=<key>&value=<value>`
    - 示例: `curl "http://localhost:8080/set?key=1&value=HelloWorld"`

- **读取键对应的值**:
    - 方法: `GET /get?key=<key>`
    - 示例: `curl "http://localhost:8080/get?key=1"`
- **利用哈希值读取键对应的值**:
    - 方法: `GET /get?BytxHash=<txHash>`
    - 示例: `curl "http://localhost:8080/getByTxHash?txHash=YOUR_HSAH"

请根据实际情况调整上述URL及参数。

