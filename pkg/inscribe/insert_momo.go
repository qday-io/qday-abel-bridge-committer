package inscribe

import (
	"encoding/json"

	"github.com/b2network/b2committer/internal/types"
	"github.com/sigurn/crc16"
)

const (
	RESERVEDFIELD   uint16 = 0xFFFF
	PROTOCOLVERSION byte   = 0x10
	ACTION          string = "inscribe"
	PROTOCOL        string = "Mable"
	NETWORK         string = "abe-test"
)

type TxMemo struct {
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
	From          string `json:"from"`
	NetworkName   string `json:"networkname,omitempty"`
	ProofRootHash string `json:"proofRootHash"`
	StateRootHash string `json:"stateRootHash"`
}

type UserTransferToSingleRecipientRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		AppID            string `json:"appID"`
		RequestSignature string `json:"requestSignature"`
		UserID           string `json:"userID"`
		Recipient        string `json:"recipient"`
		Amount           string `json:"amount"`
		PrivateKey       string `json:"privateKey"`
		Memo             []byte `json:"memo"`
	} `json:"params"`
}

type UserTransferToSingleRecipientResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Timestamp             int64       `json:"timestamp"`
		TxHash                string      `json:"txHash"`
		SignedTransactionData interface{} `json:"signedTransactionData"`
	} `json:"result"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func genUserTransferToSingleRecipientReq(abeCfg *types.AbecConfig, memo []byte) *UserTransferToSingleRecipientRequest {
	request := &UserTransferToSingleRecipientRequest{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "abelsn_userTransferToSingleRecipient",
		Params: struct {
			AppID            string `json:"appID"`
			RequestSignature string `json:"requestSignature"`
			UserID           string `json:"userID"`
			Recipient        string `json:"recipient"`
			Amount           string `json:"amount"`
			PrivateKey       string `json:"privateKey"`
			Memo             []byte `json:"memo"`
		}{
			AppID:            abeCfg.APPID,
			RequestSignature: abeCfg.RequestSignature,
			UserID:           abeCfg.UserID,
			Recipient:        abeCfg.Recipient,
			Amount:           "100000", //0.01abe 不能为0
			PrivateKey:       abeCfg.PrivateKey,
			Memo:             memo,
		},
	}

	return request
}
func GenerateMemoData(from, stateRootHash, proofRootHash string) ([]byte, error) {
	memo := TxMemo{
		Action:        ACTION,
		Protocol:      PROTOCOL,
		NetworkName:   NETWORK,
		From:          from,
		ProofRootHash: proofRootHash,
		StateRootHash: stateRootHash,
	}

	jsonBytes, err := json.Marshal(memo)
	if err != nil {
		return nil, err
	}

	table := crc16.MakeTable(crc16.CRC16_XMODEM)
	crcValue := crc16.Checksum(jsonBytes, table)

	var memoBytes []byte
	memoBytes = append(memoBytes, 0x00)                                               // 第1字节固定为0x00
	memoBytes = append(memoBytes, PROTOCOLVERSION)                                    // 第2字节为协议版本，暂定0x10
	memoBytes = append(memoBytes, byte(len(jsonBytes)>>8), byte(len(jsonBytes)&0xFF)) // 第3-4字节为长度
	memoBytes = append(memoBytes, byte(crcValue>>8), byte(crcValue&0xFF))             // 第5-6字节为CRC16校验码
	memoBytes = append(memoBytes, byte(RESERVEDFIELD>>8), byte(RESERVEDFIELD&0xFF))   // 第7-8字节预留字段
	memoBytes = append(memoBytes, jsonBytes...)                                       // base64编码的memo数据

	return memoBytes, nil
}
