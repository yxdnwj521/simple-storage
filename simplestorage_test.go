package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TestSimpleStorage 测试 SimpleStorage 合约的基本功能
func TestSimpleStorage(t *testing.T) {
	// 连接到本地 Ganache 节点
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// 使用 Ganache 提供的私钥
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY") // 替换 YOUR_PRIVATE_KEY 为你的私钥
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	// 创建交易选项
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = uint64(3000000) // 设置 Gas 限制
	auth.Value = big.NewInt(0)      // 设置交易金额

	// 部署合约
	contractAddress, _, simpleStorage, err := DeploySimpleStorage(auth, client)
	if err != nil {
		t.Fatalf("Failed to deploy contract: %v", err)
	}
	t.Logf("Contract deployed at: %s", contractAddress.Hex())

	// 设置数据
	key, value := big.NewInt(1), "Hello, Blockchain!"
	tx, err := simpleStorage.CallTransact(auth, "set", key, value)
	if err != nil {
		t.Fatalf("Failed to set data: %v", err)
	}
	t.Logf("Data set with transaction: %s", tx.Hash().Hex())

	// 获取数据
	var result string
	err = simpleStorage.Call(&bind.CallOpts{}, &result, "get", key)
	if err != nil {
		t.Fatalf("Failed to get data: %v", err)
	}
	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	} else {
		t.Logf("Retrieved value: %s", result)
	}
}
