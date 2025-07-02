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
	APPID            string `env:"APPID" envDefault:"8b9ca2d7"`
	RequestSignature string `env:"REQUEST_SIGNATURE" envDefault:"0x338i3jejjd"`
	UserID           string `env:"USERID" envDefault:"abe3238c46312425ffffd1250f3a7024ff31ad8d15fc6eeb5ad38962115640e59e94e8da112a82192d90e66539eea6427c9fb052b27ae534c8f2835b8d9c12adc1ac"`
	From             string `env:"FROM" envDefault:"abe3238c46312425ffffd1250f3a7024ff31ad8d15fc6eeb5ad38962115640e59e94e8da112a82192d90e66539eea6427c9fb052b27ae534c8f2835b8d9c12adc1ac"`
	Recipient        string `env:"RECIPIENT" envDefault:"abe3326bc9dcce62bdaecaa9c7f6b304b698fdf2ebecec442fe8b75b9be12f480aabff25326d2a5af63f16db410a1b2f02f2d1c21a6f79261443c3045444df11032d"`
	PrivateKey       string `env:"PRIVATE_KEY" envDefault:"00000000df5dbe326a891678dce65726de3bc83676835472826a911291520115d7bcf22d2b5ee9836a81defcadd8a656bbeaea33fb53f482dac90c330356d6faee92056fbc8b49d0ac2b3af022821d6f2b60b43ccc0830260291f7506e335e5cd6692c1e0c77d2903015c47b5d3c3bee0b0ef5c7d86e1ebcab0879a8400cffeb4a713bd0"`
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
