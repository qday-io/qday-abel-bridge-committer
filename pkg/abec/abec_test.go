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
	cfg.PrivateKey = "000000009efa88a8019f647299217d808c652b96c28f7fe8887830fdc8b1803002806a61e61fa689e9f4c99340336d3ea6904053026e3f15eebf09bcd36c85f9fbc186bf74023558625a2672f53602aee7a2d359cad4a451e1e047f3873750e3b59f32ca7a45e864473454f23b675e23f8e24c55323245d12f7516653119a92b0fc4cf387a6948f261dda683c2e56b1c1a0b85509663dfcdac87461fbce44a80530339debd2b6e62d3324909f7643281b4b75b1fdadc398d2801512b0b537c33a99ab1e1e1d7e336d694175dcf9d924cd2402f7b2b168d8dca557ce6b0371d281085fe8a8ebf5a643fda56b47dd057f4710150310f03b26f5d6c48860890790c4ef5978f0378a23192feb2d0527e73813281999ebf25e7aed1d1e8187fff36d9cfd614e2047b94b2e68b8894d4d0ad482a12b93991465d73c0294d9d1cdb7e9fa20df4169aa1a006ce6e39491c0913c6e50654be0bdccb619d07b4b4b0fe323458baafbcc7460676c497cc51d08fb36fb76260615acbab8d27e307c650ae127ac42e7be4c5439af88c8b8197b3726dc4b2a377db6160186c1481f79395fc82931d556678a02f6420534f38cbd7a0a4a1eea83f1e20fd87918b8f54cc2978681c57a6c9f3a445b29483141bfdc4e29c399a907bb8c41cf4e83e9cb6e508de65511a631e51c00f87ce8155439dd1361a5f78836befaa9f345386f75b941f14812894cb8c059bed886d5f7909ca8e60441986cb84262dd1f67bd7ac633c0ce3b267334931f3e8a03acc84e574c6443b6f7b84dcb5c82921be7e9fa2e26f78a39f2e11a6d27d900c37fcac2d23952da734a907bf07bf66c0c4d4231e857f3ba2b7c5c63664052c1597658f919759e557fd6ade6ece085319c1e1ebabe429b51bf725c682f1fc178bea47e215a90aa6495359218d9e5608dad5ad6a56f4b7de6abe3c1c96646fc2380684d546258345104a87c3e2bbb66aa406d5a9f5a364dd332cdea6bbec70c295e73d22dce93ff84ba2b0f94d67b20c29170333cea0561af7d5e84f66195407da9fec24e811b5f3430ebf8de0c44b67f51200481cfcaef157590da38806b463e5582f478e0102a3cb4889a01008b87d17af4f56a0b910098033f843d5d9d7f6aa10f021c1dbe2d678b87983c6e4ba575934752b16428f4a7377f58c81e3c7958f4250c8ae3e303e815854b1475d46c94b07be21369d5c1d088b50303faf4838790814fbf29574c74a860f04158b619b822a9bb6487c6190ceb61cc8d551651cf4cd1730f655903e43779f8f65e6a1483d3058282edfda21c022bc0799185e2c24342c61d78dfb3a34264b601521074488f78872752b187cbce2fc7d4168cf76b096521c890f82dcaf6bdd2e59614507381fd4088ab45acac0d505fb46e3dc7cbcce7e29aabdb6fb0dd51cf9ab3988f60219008b1f4517b00156048fcf467c0a3f51bde284ebe0e5edfa7b33da1274598f11e364e3610672cf47b17b4855a8"
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
