package abec

import (
	core "github.com/pqabelian/abelian-sdk-go"
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
	_, chainInfo, err := c.innerClient.GetChainInfo()
	if err != nil {
		return 0, err
	}

	return chainInfo.NumBlocks, nil
}
