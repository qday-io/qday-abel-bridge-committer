package abec

import (
	"testing"

	"github.com/b2network/b2committer/internal/types"
)

func TestAbecClient_GetBestBlockHeight(t *testing.T) {
	b := &AbecClient{
		endpoint:  "https://testnet-rpc-00.abelian.info",
		username:  "J8y0OnkS2wx9XEgUlW5MqtoRDAQ=",
		password:  "ULlXc/ZMJ375cn6VuSbtU+Y3KGQ=",
		authToken: "",
	}
	h, err := b.GetBestBlockHeight()
	if err != nil {
		t.Errorf("GetBestBlockHeight() error = %v", err)
		return
	}
	t.Logf("pass height: %d", h)

}

func TestAbecClient_GetTxConfirmedStatus(t *testing.T) {
	b := &AbecClient{
		endpoint:  "https://testnet-snode.abelian.info/v1/single-account",
		username:  "",
		password:  "",
		authToken: "cce71078669ded3517d085ae0b986e7d",
	}
	_, r, err := b.GetTxConfirmedStatus("", "cce71078", "abe338ce0ce178fb0aca42b4e400cdf395c92cbf9c5c9abd678aa516835f697bd6d280b285815924f862352c5463421c9f8d247f65dc112aa04c25de925bd1d1a334", "0x338i3jejjd")

	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("pass height: %d", r)
}

func TestAbecClient_UserTransferToSingleRecipient(t *testing.T) {

	b := &AbecClient{
		endpoint:  "https://testnet-snode.abelian.info/v1/single-account",
		username:  "",
		password:  "",
		authToken: "cce71078669ded3517d085ae0b986e7d",
	}

	cfg := &types.AbecConfig{}

	memo := []byte("")

	got, err := b.UserTransferToSingleRecipient(cfg, memo)

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("pass hex: %v", got)

}
