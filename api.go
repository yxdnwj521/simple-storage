package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/mux"
)

// NewSimpleStorage 创建一个新的 SimpleStorage 实例
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) *SimpleStorage {
	// 请替换以下 JSON 字符串为实际编译生成的 ABI
	abi, _ := abi.JSON(strings.NewReader(`[ABI JSON]`))
	return &SimpleStorage{BoundContract: bind.NewBoundContract(address, abi, backend, backend, backend)}
}

// SetHandler 处理设置数据的请求
func SetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key, value := params["key"], params["value"]

	// 解析 key 为 big.Int
	bigKey, ok := new(big.Int).SetString(key, 10)
	if !ok {
		http.Error(w, "Invalid key", http.StatusBadRequest)
		return
	}

	// 部署合约
	contractAddress, _, simpleStorage, err := DeploySimpleStorage(auth, client)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to deploy contract: %v", err), http.StatusInternalServerError)
		return
	}

	// 设置数据
	tx, err := simpleStorage.CallTransact(auth, "set", bigKey, value)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set data: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"transactionHash": tx.Hash().Hex()}
	json.NewEncoder(w).Encode(response)
}

// GetHandler 处理获取数据的请求
func GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]

	// 解析 key 为 big.Int
	bigKey, ok := new(big.Int).SetString(key, 10)
	if !ok {
		http.Error(w, "Invalid key", http.StatusBadRequest)
		return
	}

	// 部署合约
	contractAddress, _, simpleStorage, err := DeploySimpleStorage(auth, client)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to deploy contract: %v", err), http.StatusInternalServerError)
		return
	}

	// 获取数据
	var result string
	err = simpleStorage.Call(&bind.CallOpts{}, &result, "get", bigKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get data: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"value": result}
	json.NewEncoder(w).Encode(response)
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

	// 设置路由
	router := mux.NewRouter()
	router.HandleFunc("/set/{key}/{value}", SetHandler).Methods("POST")
	router.HandleFunc("/get/{key}", GetHandler).Methods("GET")

	// 启动 HTTP 服务器
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
