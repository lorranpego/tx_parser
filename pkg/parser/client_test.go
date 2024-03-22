package parser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMockClient(handler http.Handler) *client {
	server := httptest.NewServer(handler)
	return &client{url: server.URL, seq: 0}
}

func TestClient_GetRecentBlockNumber(t *testing.T) {
	mockClient := newMockClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := blockNumberResponse{Result: "0x1b4"}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}))

	result, err := mockClient.getRecentBlockNumber()

	if err != nil {
		t.Errorf("Error getting recent block number: %v", err)
	}

	expected := "0x1b4"
	if result.Result != expected {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected, result.Result)
	}
}

func TestClient_GetBlockByNumber(t *testing.T) {
	mockClient := newMockClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := blockResponse{
			Jsonrpc: "2.0",
			Id:      1,
			Result: struct {
				Number           string        `json:"number"`
				Hash             string        `json:"hash"`
				ParentHash       string        `json:"parentHash"`
				Sha3Uncles       string        `json:"sha3Uncles"`
				LogsBloom        string        `json:"logsBloom"`
				TransactionsRoot string        `json:"transactionsRoot"`
				StateRoot        string        `json:"stateRoot"`
				ReceiptsRoot     string        `json:"receiptsRoot"`
				Miner            string        `json:"miner"`
				Difficulty       string        `json:"difficulty"`
				TotalDifficulty  string        `json:"totalDifficulty"`
				ExtraData        string        `json:"extraData"`
				Size             string        `json:"size"`
				GasLimit         string        `json:"gasLimit"`
				GasUsed          string        `json:"gasUsed"`
				Timestamp        string        `json:"timestamp"`
				Transactions     []Transaction `json:"transactions"`
				Uncles           []interface{} `json:"uncles"`
				BaseFeePerGas    string        `json:"baseFeePerGas"`
				Nonce            string        `json:"nonce"`
				MixHash          string        `json:"mixHash"`
			}{
				Number:           "0x1b4",
				Hash:             "mockHash",
				ParentHash:       "mockParentHash",
				Sha3Uncles:       "mockSha3Uncles",
				LogsBloom:        "mockLogsBloom",
				TransactionsRoot: "mockTransactionsRoot",
				StateRoot:        "mockStateRoot",
				ReceiptsRoot:     "mockReceiptsRoot",
				Miner:            "mockMiner",
				Difficulty:       "mockDifficulty",
				TotalDifficulty:  "mockTotalDifficulty",
				ExtraData:        "mockExtraData",
				Size:             "mockSize",
				GasLimit:         "mockGasLimit",
				GasUsed:          "mockGasUsed",
				Timestamp:        "mockTimestamp",
				Transactions: []Transaction{
					{From: "mockFrom1", To: "mockTo1", Value: "mockValue1"},
					{From: "mockFrom2", To: "mockTo2", Value: "mockValue2"},
				},
				Uncles:        []interface{}{"mockUncle1", "mockUncle2"},
				BaseFeePerGas: "mockBaseFeePerGas",
				Nonce:         "mockNonce",
				MixHash:       "mockMixHash",
			},
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}))

	// Call the getBlockByNumber method
	result, err := mockClient.getBlockByNumber("0x1b4")

	if err != nil {
		t.Errorf("Error getting block by number: %v", err)
	}

	expectedNumber := "0x1b4"
	if result.Result.Number != expectedNumber {
		t.Errorf("Unexpected block number. Expected: %s, Got: %s", expectedNumber, result.Result.Number)
	}

	expectedHash := "mockHash"
	if result.Result.Hash != expectedHash {
		t.Errorf("Unexpected block hash. Expected: %s, Got: %s", expectedHash, result.Result.Hash)
	}
}
