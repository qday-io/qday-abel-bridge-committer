package types

import (
	"github.com/caarlos0/env/v6"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/log"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat            string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource      string `env:"MYSQL_DATA_SOURCE" envDefault:"root:root@tcp(127.0.0.1:3366)/abe_committer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl               string `env:"RPC_URL" envDefault:"http://190.92.213.101:8545"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"da-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"1"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0x1db9fbef276c30cc9fdfba2a21d7cc4dd760fd2e76b8906571c6f56a706624bc"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0x59C123c2901245F3A38E009Eb341d34CC4609a8E"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"2"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	Address    string `env:"B2NODE_ADDRESS" envDefault:"ethm17nhkv58y35ye5jtjafd5zndtsmsenz7nlxh60c"`
	ChainID    string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost   string `env:"B2NODE_GRPC_HOST" envDefault:"190.92.213.101"`
	GRPCPort   uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9090"`
	RPCUrl     string `env:"B2NODE_RPC_URL" envDefault:"http://190.92.213.101:8545"`
	CoinDenom  string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton"`
	PrivateKey string `env:"B2NODE_PRIVATE_KEY" envDefault:"2ee789a68207020b45607f5adb71933de0946baebbaaab74af7cbd69c8a90573"`
}

// type BitcoinRPCConfig struct {
// 	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
// 	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"L2Wv2kPRs4iJ8zFwQmTPBvrPQY7DS591LdmTEqu4x6GXX4yVDr7h"`
// 	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"bc1peuj9pfr4leqnfmem4nrsxl2cqcz2hafw7k6luxzrxcr545fywm3qrsfvtj"`
// }

type AbecConfig struct {
	Endpoint         string `env:"ENDPOINT" envDefault:"https://testnet-snode.pqabelian.io/v1/single-account"`
	RpcEndpoint      string `env:"RPCENDPOINT" envDefault:"https://testnet-rpc-00.pqabelian.io"`
	Username         string `env:"USERNAME" envDefault:"J8y0OnkS2wx9XEgUlW5MqtoRDAQ="`
	Password         string `env:"PASSWORD" envDefault:"ULlXc/ZMJ375cn6VuSbtU+Y3KGQ="`
	AuthToken        string `env:"AUTHTOKEN" envDefault:"cce71078669ded3517d961a2d57eb440"`
	APPID            string `env:"APPID" envDefault:"cce71078"`
	RequestSignature string `env:"REQUEST_SIGNATURE" envDefault:"0x338i3jejjd"`
	UserID           string `env:"USERID" envDefault:"abe3238c46312425ffffd1250f3a7024ff31ad8d15fc6eeb5ad38962115640e59e94e8da112a82192d90e66539eea6427c9fb052b27ae534c8f2835b8d9c12adc1ac"`
	From             string `env:"FROM" envDefault:"abe3238c46312425ffffd1250f3a7024ff31ad8d15fc6eeb5ad38962115640e59e94e8da112a82192d90e66539eea6427c9fb052b27ae534c8f2835b8d9c12adc1ac"`
	Recipient        string `env:"RECIPIENT" envDefault:"abe3326bc9dcce62bdaecaa9c7f6b304b698fdf2ebecec442fe8b75b9be12f480aabff25326d2a5af63f16db410a1b2f02f2d1c21a6f79261443c3045444df11032d"`
	PrivateKey       string `env:"PRIVATE_KEY" envDefault:"000000009efa88a8019f647299217d808c652b96c28f7fe8887830fdc8b1803002806a61e61fa689e9f4c99340336d3ea6904053026e3f15eebf09bcd36c85f9fbc186bf74023558625a2672f53602aee7a2d359cad4a451e1e047f3873750e3b59f32ca7a45e864473454f23b675e23f8e24c55323245d12f7516653119a92b0fc4cf387a6948f261dda683c2e56b1c1a0b85509663dfcdac87461fbce44a80530339debd2b6e62d3324909f7643281b4b75b1fdadc398d2801512b0b537c33a99ab1e1e1d7e336d694175dcf9d924cd2402f7b2b168d8dca557ce6b0371d281085fe8a8ebf5a643fda56b47dd057f4710150310f03b26f5d6c48860890790c4ef5978f0378a23192feb2d0527e73813281999ebf25e7aed1d1e8187fff36d9cfd614e2047b94b2e68b8894d4d0ad482a12b93991465d73c0294d9d1cdb7e9fa20df4169aa1a006ce6e39491c0913c6e50654be0bdccb619d07b4b4b0fe323458baafbcc7460676c497cc51d08fb36fb76260615acbab8d27e307c650ae127ac42e7be4c5439af88c8b8197b3726dc4b2a377db6160186c1481f79395fc82931d556678a02f6420534f38cbd7a0a4a1eea83f1e20fd87918b8f54cc2978681c57a6c9f3a445b29483141bfdc4e29c399a907bb8c41cf4e83e9cb6e508de65511a631e51c00f87ce8155439dd1361a5f78836befaa9f345386f75b941f14812894cb8c059bed886d5f7909ca8e60441986cb84262dd1f67bd7ac633c0ce3b267334931f3e8a03acc84e574c6443b6f7b84dcb5c82921be7e9fa2e26f78a39f2e11a6d27d900c37fcac2d23952da734a907bf07bf66c0c4d4231e857f3ba2b7c5c63664052c1597658f919759e557fd6ade6ece085319c1e1ebabe429b51bf725c682f1fc178bea47e215a90aa6495359218d9e5608dad5ad6a56f4b7de6abe3c1c96646fc2380684d546258345104a87c3e2bbb66aa406d5a9f5a364dd332cdea6bbec70c295e73d22dce93ff84ba2b0f94d67b20c29170333cea0561af7d5e84f66195407da9fec24e811b5f3430ebf8de0c44b67f51200481cfcaef157590da38806b463e5582f478e0102a3cb4889a01008b87d17af4f56a0b910098033f843d5d9d7f6aa10f021c1dbe2d678b87983c6e4ba575934752b16428f4a7377f58c81e3c7958f4250c8ae3e303e815854b1475d46c94b07be21369d5c1d088b50303faf4838790814fbf29574c74a860f04158b619b822a9bb6487c6190ceb61cc8d551651cf4cd1730f655903e43779f8f65e6a1483d3058282edfda21c022bc0799185e2c24342c61d78dfb3a34264b601521074488f78872752b187cbce2fc7d4168cf76b096521c890f82dcaf6bdd2e59614507381fd4088ab45acac0d505fb46e3dc7cbcce7e29aabdb6fb0dd51cf9ab3988f60219008b1f4517b00156048fcf467c0a3f51bde284ebe0e5edfa7b33da1274598f11e364e3610672cf47b17b4855a8"`
}

var (
	config *Config
	// btcRPCConfig *BitcoinRPCConfig
	b2nodeConfig *B2NODEConfig
	abecCfg      *AbecConfig
)

func GetConfig() *Config {
	if config == nil {
		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		config = cfg
	}
	return config
}

// func GetBtcConfig() *BitcoinRPCConfig {
// 	if btcRPCConfig == nil {
// 		cfg := &BitcoinRPCConfig{}
// 		if err := env.Parse(cfg); err != nil {
// 			log.Panicf("parse config err: %s\n", err)
// 			return nil
// 		}
// 		btcRPCConfig = cfg
// 	}
// 	return btcRPCConfig
// }

func GetB2nodeConfig() *B2NODEConfig {
	if b2nodeConfig == nil {
		cfg := &B2NODEConfig{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		b2nodeConfig = cfg
	}
	return b2nodeConfig
}

func GetAbecConfig() *AbecConfig {
	if abecCfg == nil {
		cfg := &AbecConfig{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse abec config err: %s\n", err)
			return nil
		}
		abecCfg = cfg
	}
	return abecCfg
}
