package abec

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/b2network/b2committer/internal/types"
	"github.com/tidwall/gjson"
)

type AbecClient struct {
	endpoint    string
	username    string
	password    string
	authToken   string
	rpcEndpoint string
}

func (b *AbecClient) GetBestBlockHeight() (int64, error) {
	resp, err := b.getResponseFromChan("getinfo", nil)
	if err != nil {
		return 0, err
	}
	var abe AbelianChainInfo
	err = json.Unmarshal(resp, &abe)
	if err != nil {
		return 0, err
	}
	return abe.Blocks, nil
}

func (b *AbecClient) GetTxConfirmedStatus(txid, appID, userID, requestSignature string) (bool, int64, error) {
	params := map[string]interface{}{
		"appID":            appID,
		"userID":           userID,
		"requestSignature": requestSignature,
		"txid":             txid,
	}

	resp, err := b.getResponseFromChan("abelsn_userTransactionView", params)
	if err != nil {
		return false, -1, err
	}
	height := gjson.ParseBytes(resp).Get("tx.blockHeight").Int()
	return true, height, nil
}

func (b *AbecClient) UserTransferToSingleRecipient(abeCfg *types.AbecConfig, memo []byte) (string, error) {
	params := map[string]interface{}{
		"appID":            abeCfg.APPID,
		"requestSignature": abeCfg.RequestSignature,
		"userID":           abeCfg.UserID,
		"recipient":        abeCfg.Recipient,
		"amount":           "100000",
		"privateKey":       abeCfg.PrivateKey,
		"memo":             hex.EncodeToString(memo),
	}

	resp, err := b.getResponseFromChan("abelsn_userTransferToSingleRecipient", params)
	if err != nil {
		return "", err
	}

	var res UserTransferToSingleRecipientResult
	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}
	return res.TxHash, nil
}

func NewClient(endpoint, username, password, authToken, rpcEndpoint string) *AbecClient {
	return &AbecClient{
		endpoint:    endpoint,
		username:    username,
		password:    password,
		authToken:   authToken,
		rpcEndpoint: rpcEndpoint,
	}
}

type AbecJSONRPCRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      any                    `json:"id"`
}

type AbecJSONRPCResponse struct {
	Result []byte `json:"result"`
	Error  []byte `json:"error"`
	ID     string `json:"id"`
}

func (b *AbecClient) newRequest(id string, method string, params map[string]interface{}) (*http.Request, error) {
	jsonReq := &AbecJSONRPCRequest{
		JSONRPC: "1.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
	jsonBody, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, err
	}

	url := b.rpcEndpoint

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(b.username, b.password)

	return httpReq, nil
}

func (b *AbecClient) newRequestAuth2(id int64, method string, params map[string]interface{}) (*http.Request, error) {
	jsonReq := &AbecJSONRPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
	jsonBody, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, err
	}

	url := b.endpoint

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Token "+b.authToken)

	return httpReq, nil
}

func (b *AbecClient) getResponseFromChan(method string, params map[string]interface{}) ([]byte, error) {
	var req *http.Request
	var err error

	if method == "getinfo" {
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		req, err = b.newRequest(id, method, params)
	} else {
		req, err = b.newRequestAuth2(1, method, params)
	}

	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respObj := &AbecJSONRPCResponse{}
	root := gjson.ParseBytes(body)
	respObj.Result = []byte(root.Get("result").String())
	respObj.Error = []byte(root.Get("error").String())
	respObj.ID = root.Get("id").String()

	errorStr := string(respObj.Error)
	if len(errorStr) > 0 && errorStr != "null" {
		return nil, fmt.Errorf("abec.%s: %s", method, respObj.Error)
	}

	return respObj.Result, nil
}

type AbelianChainInfo struct {
	Protocolversion      int     `json:"protocolversion" gorm:"column:protocolversion"`
	Relayfee             float64 `json:"relayfee" gorm:"column:relayfee"`
	Nodetype             string  `json:"nodetype" gorm:"column:nodetype"`
	Timeoffset           int64   `json:"timeoffset" gorm:"column:timeoffset"`
	Blocks               int64   `json:"blocks" gorm:"column:blocks"`
	Witnessserviceheight int64   `json:"witnessserviceheight" gorm:"column:witnessserviceheight"`
	Version              int64   `json:"version" gorm:"column:version"`
	Difficulty           float64 `json:"difficulty" gorm:"column:difficulty"`
	Proxy                string  `json:"proxy" gorm:"column:proxy"`
	Worksum              string  `json:"worksum" gorm:"column:worksum"`
	Bestblockhash        string  `json:"bestblockhash" gorm:"column:bestblockhash"`
	Testnet              bool    `json:"testnet" gorm:"column:testnet"`
	Connections          int64   `json:"connections" gorm:"column:connections"`
	Errors               string  `json:"errors" gorm:"column:errors"`
	NetId                int64   `json:"netid" gorm:"column:netid"`
}

type TransactionViewResult struct {
	BlockHeight int64 `json:"blockHeight"`
}

type UserTransferToSingleRecipientResult struct {
	Timestamp             int64       `json:"timestamp"`
	TxHash                string      `json:"txHash"`
	SignedTransactionData interface{} `json:"signedTransactionData"`
}
