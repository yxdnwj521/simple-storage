package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SimpleStorage 是智能合约的绑定结构体
type SimpleStorage struct {
	contract *bind.BoundContract
	abi      abi.ABI
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, r *Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

// NewSimpleStorage 创建一个新的 SimpleStorage 实例
func NewSimpleStorage(address common.Address, backend bind.ContractBackend, abiJSON string) (*SimpleStorage, error) {
	parsedAbi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsedAbi, backend, backend, backend)
	return &SimpleStorage{contract: contract, abi: parsedAbi}, nil
}

// Call 调用合约的只读方法
func (s *SimpleStorage) Call(opts *bind.CallOpts, method string, args ...interface{}) ([]string, error) {
	var result []interface{}
	err := s.contract.Call(opts, &result, method, args...)
	if err != nil {
		return nil, err
	}

	var resultValues []string
	for _, res := range result {
		if str, ok := res.(string); ok {
			resultValues = append(resultValues, str)
		} else {
			return nil, fmt.Errorf("unexpected result type: %T", res)
		}
	}
	return resultValues, nil
}

// Transact 调用合约的写入方法
func (s *SimpleStorage) Transact(auth *bind.TransactOpts, method string, args ...interface{}) (*types.Transaction, error) {
	tx, err := s.contract.Transact(auth, method, args...)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

var (
	contractAddress = common.HexToAddress("0x5a57F93f7C2510E643e3cB1d3bcfBD028D8DAfd2")
	localNetworkURL = "http://localhost:7545"
	abiJSON         []byte
	privateKey      string
)

type Config struct {
	PrivateKey string `json:"privateKey"`
}

func loadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func setHandler(client *ethclient.Client, storage *SimpleStorage, privateKeyECDSA *ecdsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyStr := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")

		keyUint, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: "无效的键值类型，必须为整数"})
			return
		}

		// 动态获取当前网络的链 ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: "无法获取网络链 ID"})
			return
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: "创建交易授权失败"})
			return
		}

		bigKey := big.NewInt(0).SetUint64(keyUint)

		tx, err := storage.Transact(auth, "set", bigKey, value)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: fmt.Sprintf("设置值失败: %v", err)})
			return
		}

		writeResponse(w, &Response{Msg: "ok", Data: tx.Hash().Hex()})
	}
}

func getHandler(storage *SimpleStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyStr := r.URL.Query().Get("key")
		if keyStr == "" {
			writeResponse(w, &Response{Msg: "error", Data: "缺少键值"})
			return
		}

		keyUint, err := strconv.ParseUint(keyStr, 10, 64)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: "无效的键值类型，必须为整数"})
			return
		}

		bigKey := big.NewInt(0).SetUint64(keyUint)

		opts := &bind.CallOpts{
			Context: context.Background(),
		}

		result, err := storage.Call(opts, "get", bigKey)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: fmt.Sprintf("获取值失败: %v", err)})
			return
		}

		writeResponse(w, &Response{Msg: "ok", Data: result[0]})
	}
}

// getByTxHashHandler 通过交易哈希查询存储在智能合约中的值
func getByTxHashHandler(client *ethclient.Client, storage *SimpleStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txHashStr := r.URL.Query().Get("txHash")

		if txHashStr == "" {
			writeResponse(w, &Response{Msg: "error", Data: "缺少 txHash 参数"})
			return
		}

		// 通过交易哈希获取键值
		txHash := common.HexToHash(txHashStr)

		// 获取交易详情
		tx, _, err := getTransactionDetails(client, txHash)
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: fmt.Sprintf("无法获取交易详情: %v", err)})
			return
		}

		// 解析交易输入
		methodName, args, err := parseTransactionInput(string(abiJSON), tx.Data())
		if err != nil {
			writeResponse(w, &Response{Msg: "error", Data: fmt.Sprintf("无法解析交易输入: %v", err)})
			return
		}

		// 假设方法名为 "set"，并且第一个参数是键
		if methodName == "set" && len(args) > 0 {
			key, ok := args[0].(*big.Int)
			if !ok {
				writeResponse(w, &Response{Msg: "error", Data: "键值类型不匹配"})
				return
			}

			// 通过键值调用智能合约的 get 方法
			callOpts := &bind.CallOpts{
				Pending: false,
			}
			result, err := storage.Call(callOpts, "get", key)
			if err != nil {
				writeResponse(w, &Response{Msg: "error", Data: fmt.Sprintf("获取数据失败: %v", err)})
				return
			}

			// 假设结果只包含一个字符串值
			if len(result) > 0 {
				writeResponse(w, &Response{Msg: "ok", Data: result[0]})
			} else {
				writeResponse(w, &Response{Msg: "error", Data: "未从合约中获得任何结果"})
			}
		} else {
			writeResponse(w, &Response{Msg: "error", Data: "无效的方法或参数"})
		}
	}
}

// getTransactionDetails 获取交易详情
func getTransactionDetails(client *ethclient.Client, txHash common.Hash) (*types.Transaction, *types.Receipt, error) {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, nil, fmt.Errorf("无法获取交易: %v", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, nil, fmt.Errorf("无法获取交易收据: %v", err)
	}

	return tx, receipt, nil
}

// parseTransactionInput 解析交易输入
func parseTransactionInput(abiJSON string, input []byte) (methodName string, args []interface{}, err error) {
	parsedAbi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", nil, fmt.Errorf("解析 ABI 失败: %v", err)
	}

	method, err := parsedAbi.MethodById(input[:4])
	if err != nil {
		return "", nil, fmt.Errorf("无法找到方法: %v", err)
	}

	decodedArgs, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return "", nil, fmt.Errorf("解包输入失败: %v", err)
	}

	return method.Name, decodedArgs, nil
}

func main() {
	// 加载配置文件
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	privateKey = config.PrivateKey

	// 连接到本地以太坊节点
	client, err := ethclient.Dial(localNetworkURL)
	if err != nil {
		log.Fatalf("连接到以太坊节点失败: %v", err)
	}

	// 加载 ABI JSON
	abiJSON, err = ioutil.ReadFile("SimpleStorage.abi")
	if err != nil {
		log.Fatalf("加载 ABI 文件失败: %v", err)
	}

	// 创建 SimpleStorage 实例
	storage, err := NewSimpleStorage(contractAddress, client, string(abiJSON))
	if err != nil {
		log.Fatalf("创建 SimpleStorage 实例失败: %v", err)
	}

	// 解析私钥
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	// 设置 HTTP 处理函数
	http.HandleFunc("/set", setHandler(client, storage, privateKeyECDSA))
	http.HandleFunc("/get", getHandler(storage))
	http.HandleFunc("/getByTxHash", getByTxHashHandler(client, storage))

	// 启动 HTTP 服务器
	log.Println("启动 HTTP 服务器...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
