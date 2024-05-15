package abec

import (
	core "abelian.info/sdk/core"
	"github.com/btcsuite/btcd/wire"
	"github.com/pqabelian/abec/chainhash"
)

type AbecClient struct {
	innerClient *core.AbecRPCClient
}

func NewClient(endpoint, username, password string) *AbecClient {
	innerClient := core.NewAbecRPCClient(endpoint, username, password)
	return &AbecClient{
		innerClient: innerClient,
	}
}

func (c *AbecClient) GetBestBlockHeight() (int64, error) {
	return c.innerClient.GetBestBlock().height
}

func (c *AbecClient) SendRawTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	return c.innerClient.SendRawTx(tx)
}
