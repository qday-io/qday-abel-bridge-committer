package abec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type AbecClient struct {
	endpoint string
	username string
	password string
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
		"txid":             txid,
		"appID":            appID,
		"userID":           userID,
		"requestSignature": requestSignature,
	}

	resp, err := b.getResponseFromChan("getinfo", params)
	if err != nil {
		return false, -1, err
	}

	var txViewRes TransactionViewResult
	err = json.Unmarshal(resp, &txViewRes)
	if err != nil {
		return false, -1, err
	}

	if txViewRes.BlockHeight > 0 {
		// confirmed height != -1
		return true, txViewRes.BlockHeight, nil
	}

	return false, -1, nil
}

func NewClient(endpoint string, username string, password string) *AbecClient {
	return &AbecClient{
		endpoint: endpoint,
		username: username,
		password: password,
	}
}

type AbecJSONRPCRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      string                 `json:"id"`
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

	//url := "https://testnet-rpc-00.abelian.info"
	url := b.endpoint

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(b.username, b.password)

	return httpReq, nil
}

func (b *AbecClient) getResponseFromChan(method string, params map[string]interface{}) ([]byte, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	req, err := b.newRequest(id, method, params)
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
