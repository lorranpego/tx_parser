package memory

import (
	"reflect"
	"sync"
	"testing"

	"tx_parser/pkg/parser"
)

func TestMemoryStorage_Insert(t *testing.T) {
	sd := &memoryStorage{sync.RWMutex{}, make(map[string][]parser.Transaction)}

	// Test case: Insert a transaction into the memory storage
	address := "123"
	transaction := parser.Transaction{
		Type:                 "mockType",
		BlockHash:            "mockBlockHash",
		BlockNumber:          "mockBlockNumber",
		From:                 "mockFrom",
		Gas:                  "mockGas",
		Hash:                 "mockHash",
		Input:                "mockInput",
		Nonce:                "mockNonce",
		To:                   "mockTo",
		TransactionIndex:     "mockTransactionIndex",
		Value:                "mockValue",
		V:                    "mockV",
		R:                    "mockR",
		S:                    "mockS",
		GasPrice:             "mockGasPrice",
		MaxFeePerGas:         "mockMaxFeePerGas",
		MaxPriorityFeePerGas: "mockMaxPriorityFeePerGas",
		ChainId:              "mockChainId",
		AccessList: []struct {
			Address     string   `json:"address"`
			StorageKeys []string `json:"storageKeys"`
		}{
			{
				Address:     "mockAccessListAddress",
				StorageKeys: []string{"mockStorageKey1", "mockStorageKey2"},
			},
		},
	}

	err := sd.Insert(address, transaction)

	if err != nil {
		t.Errorf("Error inserting transaction: %v", err)
	}

	if len(sd.data[address]) != 1 || !reflect.DeepEqual(sd.data[address][0], transaction) {
		t.Errorf("Inserted transaction not found in storage")
	}
}

func TestMemoryStorage_Get(t *testing.T) {
	sd := &memoryStorage{sync.RWMutex{}, make(map[string][]parser.Transaction)}

	address := "123"
	transaction := parser.Transaction{
		Type:                 "mockType",
		BlockHash:            "mockBlockHash",
		BlockNumber:          "mockBlockNumber",
		From:                 "mockFrom",
		Gas:                  "mockGas",
		Hash:                 "mockHash",
		Input:                "mockInput",
		Nonce:                "mockNonce",
		To:                   "mockTo",
		TransactionIndex:     "mockTransactionIndex",
		Value:                "mockValue",
		V:                    "mockV",
		R:                    "mockR",
		S:                    "mockS",
		GasPrice:             "mockGasPrice",
		MaxFeePerGas:         "mockMaxFeePerGas",
		MaxPriorityFeePerGas: "mockMaxPriorityFeePerGas",
		ChainId:              "mockChainId",
		AccessList: []struct {
			Address     string   `json:"address"`
			StorageKeys []string `json:"storageKeys"`
		}{
			{
				Address:     "mockAccessListAddress",
				StorageKeys: []string{"mockStorageKey1", "mockStorageKey2"},
			},
		},
	}
	sd.data[address] = append(sd.data[address], transaction)

	// Test case: Get transactions for a given address
	result, err := sd.Get(address)

	if err != nil {
		t.Errorf("Error getting transactions: %v", err)
	}

	if len(result) != 1 || !reflect.DeepEqual(result[0], transaction) {
		t.Errorf("Retrieved transaction does not match inserted transaction")
	}
}
