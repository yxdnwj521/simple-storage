# 使用 Node.js 作为基础镜像
FROM node:16

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json 到容器中
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制项目文件到容器中
COPY . .

# 运行测试
CMD ["npx", "hardhat", "test"]