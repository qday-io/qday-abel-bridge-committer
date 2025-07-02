package abec

import (
	"testing"

	"github.com/qday-io/qday-abel-bridge-committer/pkg/types"
	"github.com/sigurn/crc16"
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
	_, r, err := b.GetTxConfirmedStatus("0xd6dfb90c0bb3c533df36c333cbe36376fe32c04d7c5be7b913b391209206cab0", "cce71078", "abe338ce0ce178fb0aca42b4e400cdf395c92cbf9c5c9abd678aa516835f697bd6d280b285815924f862352c5463421c9f8d247f65dc112aa04c25de925bd1d1a334", "0x338i3jejjd")

	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("pass height: %d", r)
}

func TestAbecClient_UserTransferToSingleRecipient(t *testing.T) {

	b := &AbecClient{
		endpoint:    "https://testnet-snode.pqabelian.io/v1/single-account",
		rpcEndpoint: "https://testnet-rpc-00.pqabelian.io",
		username:    "J8y0OnkS2wx9XEgUlW5MqtoRDAQ=",
		password:    "ULlXc/ZMJ375cn6VuSbtU+Y3KGQ=",
		authToken:   "cce71078669ded3517d961a2d57eb440",
	}

	cfg := &types.AbecConfig{}
	cfg.APPID = "8b9ca2d7"
	cfg.RequestSignature = "0x338i3jejjd"
	cfg.UserID = "abe3238c46312425ffffd1250f3a7024ff31ad8d15fc6eeb5ad38962115640e59e94e8da112a82192d90e66539eea6427c9fb052b27ae534c8f2835b8d9c12adc1ac"
	cfg.Recipient = "abe3326bc9dcce62bdaecaa9c7f6b304b698fdf2ebecec442fe8b75b9be12f480aabff25326d2a5af63f16db410a1b2f02f2d1c21a6f79261443c3045444df11032d"
	cfg.PrivateKey = "00000000df5dbe326a891678dce65726de3bc83676835472826a911291520115d7bcf22d2b5ee9836a81defcadd8a656bbeaea33fb53f482dac90c330356d6faee92056fbc8b49d0ac2b3af022821d6f2b60b43ccc0830260291f7506e335e5cd6692c1e0c77d2903015c47b5d3c3bee0b0ef5c7d86e1ebcab0879a8400cffeb4a713bd0"
	memo := []byte("ddddddddddd")

	got, err := b.UserTransferToSingleRecipient(cfg, memo, "100")

	if err != nil {
		t.Errorf("err:%v", err.Error())
		return
	}

	t.Logf("pass txhash: %v", got)

}

func TestAbecClient_SendDepositTx(t *testing.T) {

	// TODO 注意，如果不使用我这个账户，那么authtoken,appid,userid,RequestSignature,privatekey都需要改动
	b := &AbecClient{
		endpoint:  "https://testnet-snode.abelian.info/v1/single-account",
		username:  "",
		password:  "",
		authToken: "8b9ca2d7f0d4d76e17d02f6f5f82e595",
	}

	cfg := &types.AbecConfig{}
	cfg.APPID = "8b9ca2d7"
	cfg.RequestSignature = "randstring"
	cfg.UserID = "abe32f5c9dd67b6f0e11333fc54e4b54d1f05456ea0e2abc6e1459b056271e3de6180f7cca4ca880a8839c72d412987ffd47d7fdca60fce5838bfcbea68dd741146b"

	// TODO  需要换接收者
	cfg.Recipient = "abe338ce0ce178fb0aca42b4e400cdf395c92cbf9c5c9abd678aa516835f697bd6d280b285815924f862352c5463421c9f8d247f65dc112aa04c25de925bd1d1a334"

	cfg.PrivateKey = "0000000064a27b5f97581f0eaeb482d09fb963e0e19f73eb476b6de0d9821967abdc8ea9336bf818d3828d94eb2bfca150fec85dccbbc18c6c6d39a3bd2fbb2a5801c525c42815fe86639ad806246bac5810ea820bdd3ce87d0c1718716019aba621cd3507156e8a72e7a41d81615788392dfd42974ead6a229aeebedf448f091e517d85"

	str := `{
	  "action": "deposit",
	  "protocol": "Mable",
	  "from": "abe32f5c9dd67b6f0e11333fc54e4b54d1f05456ea0e2abc6e1459b056271e3de6180f7cca4ca880a8839c72d412987ffd47d7fdca60fce5838bfcbea68dd741146b",
	  "to": "abe338ce0ce178fb0aca42b4e400cdf395c92cbf9c5c9abd678aa516835f697bd6d280b285815924f862352c5463421c9f8d247f65dc112aa04c25de925bd1d1a334",
	  "receipt": "0xdac17f958d2ee523a2206206994597c13d831ec7",
	  "value": "0x21222200"
	}`
	// TODO 需要构造deposit txmemo
	memo := []byte(str)

	// TODO 根据需求改动amount , 这里100000为0.01ABE
	amount := "100000"

	var RESERVEDFIELD uint16 = 0xFFFF
	var PROTOCOLVERSION byte = 0x10

	table := crc16.MakeTable(crc16.CRC16_XMODEM)
	crcValue := crc16.Checksum(memo, table)

	var memoBytes []byte
	memoBytes = append(memoBytes, 0x00)                                             // 第1字节固定为0x00
	memoBytes = append(memoBytes, PROTOCOLVERSION)                                  // 第2字节为协议版本，暂定0x10
	memoBytes = append(memoBytes, byte(len(memo)>>8), byte(len(memo)&0xFF))         // 第3-4字节为长度
	memoBytes = append(memoBytes, byte(crcValue>>8), byte(crcValue&0xFF))           // 第5-6字节为CRC16校验码
	memoBytes = append(memoBytes, byte(RESERVEDFIELD>>8), byte(RESERVEDFIELD&0xFF)) // 第7-8字节预留字段
	memoBytes = append(memoBytes, memo...)

	got, err := b.UserTransferToSingleRecipient(cfg, memoBytes, amount)

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("pass txhash: %v", got)
}
