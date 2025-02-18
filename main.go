package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SimpleStorage 是根据合约 ABI 自动生成的绑定
type SimpleStorage struct {
	*bind.BoundContract
}

// NewSimpleStorage 创建一个新的 SimpleStorage 实例
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) *SimpleStorage {
	// 请替换以下 JSON 字符串为实际编译生成的 ABI
	abi, err := abi.JSON(strings.NewReader(`[ABI JSON]`))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}
	return &SimpleStorage{BoundContract: bind.NewBoundContract(address, abi, backend, backend, backend)}
}

// DeploySimpleStorage 部署合约
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	// 请替换以下 JSON 字符串为实际编译生成的 ABI
	abi, err := abi.JSON(strings.NewReader(`[ABI JSON]`))
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed to parse ABI: %w", err)
	}
	// 请替换 [BYTECODE] 为编译后得到的实际字节码
	bytecode := `[BYTECODE]`

	address, tx, simpleStorage, err := bind.DeployContract(auth, abi, common.FromHex(bytecode), backend, nil)
	if err != nil {
		return common.Address{}, nil, nil, fmt.Errorf("failed to deploy contract: %w", err)
	}
	return address, tx, simpleStorage, nil
}

func main() {
	// 连接到本地 Ganache 节点
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// 使用 Ganache 提供的私钥
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY") // 替换 YOUR_PRIVATE_KEY 为你的私钥
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 创建交易选项
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = uint64(3000000) // 设置 Gas 限制
	auth.Value = big.NewInt(0)      // 设置交易金额

	// 部署合约
	contractAddress, _, simpleStorage, err := DeploySimpleStorage(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Println("Contract deployed at:", contractAddress.Hex())

	// 设置数据
	key, value := big.NewInt(1), "Hello, Blockchain!"
	tx, err := simpleStorage.CallTransact(auth, "set", key, value)
	if err != nil {
		log.Fatalf("Failed to set data: %v", err)
	}
	fmt.Printf("Data set with transaction: %s\n", tx.Hash().Hex())

	// 获取数据
	var result string
	err = simpleStorage.Call(&bind.CallOpts{}, &result, "get", key)
	if err != nil {
		log.Fatalf("Failed to get data: %v", err)
	}
	fmt.Printf("Retrieved value: %s\n", result) // 输出应为 "Hello, Blockchain!"
}
