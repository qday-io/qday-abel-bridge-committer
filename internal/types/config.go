package types

import (
	"github.com/b2network/b2committer/pkg/log"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// debug", "info", "warn", "error", "panic", "fatal"
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// "console","json"
	LogFormat            string `env:"LOG_FORMAT" envDefault:"console"`
	MySQLDataSource      string `env:"MYSQL_DATA_SOURCE" envDefault:"root:123456@tcp(127.0.0.1:3306)/b2_committer?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"`
	MySQLMaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS" envDefault:"10"`
	MySQLMaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS" envDefault:"20"`
	MySQLConnMaxLifetime int    `env:"MYSQL_CONN_MAX_LIFETIME" envDefault:"3600"`
	RPCUrl               string `env:"RPC_URL" envDefault:"http://159.138.106.168:8545"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"300"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0x0c4408fab90dae9ee884f255ba647aabd387182809f43aa416fd2d255b0f7c4f"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0xBEfc6B93f8987758231f2381B562996deeE963F9"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"2"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	Address    string `env:"B2NODE_ADDRESS" envDefault:"ethm1e7s8kwnlcnd5vk8fx0jrfpy5cn5wfmwakc5azp"`
	ChainID    string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost   string `env:"B2NODE_GRPC_HOST" envDefault:"159.138.106.168"`
	GRPCPort   uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9090"`
	RPCUrl     string `env:"B2NODE_RPC_URL" envDefault:"http://159.138.106.168:9090"`
	CoinDenom  string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton"`
	PrivateKey string `env:"B2NODE_PRIVATE_KEY" envDefault:"8623eb1173b001788b7dc789513c34d049a3d02c728b50daae5799fca009e111"`
}

type BitcoinRPCConfig struct {
	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"L2Wv2kPRs4iJ8zFwQmTPBvrPQY7DS591LdmTEqu4x6GXX4yVDr7h"`
	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"bc1peuj9pfr4leqnfmem4nrsxl2cqcz2hafw7k6luxzrxcr545fywm3qrsfvtj"`
}

type AbecConfig struct {
	Endpoint         string `env:"ENDPOINT" envDefault:"https://testnet-snode.abelian.info/v1/single-account"`
	Username         string `env:"USERNAME" envDefault:"J8y0OnkS2wx9XEgUlW5MqtoRDAQ="`
	Password         string `env:"PASSWORD" envDefault:"ULlXc/ZMJ375cn6VuSbtU+Y3KGQ="`
	APPID            string `env:"APPID" envDefault:"8b9ca2d7"`
	RequestSignature string `env:"REQUEST_SIGNATURE" envDefault:"randstring"`
	UserID           string `env:"USERID" envDefault:"abe32f5c9dd67b6f0e11333fc54e4b54d1f05456ea0e2abc6e1459b056271e3de6180f7cca4ca880a8839c72d412987ffd47d7fdca60fce5838bfcbea68dd741146b"`
	From             string `env:"FROM" envDefault:"abe32f5c9dd67b6f0e11333fc54e4b54d1f05456ea0e2abc6e1459b056271e3de6180f7cca4ca880a8839c72d412987ffd47d7fdca60fce5838bfcbea68dd741146b"`
	Recipient        string `env:"RECIPIENT" envDefault:"abe338491ef250a530f6b1a771d45ae168f81d6a430f20623849e448b870f0f95e13f12ba51bff83497480db944567750e3cf555cd9811db95b848ca93d45c1448d0"`
	PrivateKey       string `env:"PRIVATE_KEY" envDefault:"0000000064a27b5f97581f0eaeb482d09fb963e0e19f73eb476b6de0d9821967abdc8ea9336bf818d3828d94eb2bfca150fec85dccbbc18c6c6d39a3bd2fbb2a5801c525c42815fe86639ad806246bac5810ea820bdd3ce87d0c1718716019aba621cd3507156e8a72e7a41d81615788392dfd42974ead6a229aeebedf448f091e517d85"`
	AuthToken        string `env:"AUTHTOKEN" envDefault:"8b9ca2d7f0d4d76e17d02f6f5f82e595"`
}

var (
	config       *Config
	btcRPCConfig *BitcoinRPCConfig
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

func GetBtcConfig() *BitcoinRPCConfig {
	if btcRPCConfig == nil {
		cfg := &BitcoinRPCConfig{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("parse config err: %s\n", err)
			return nil
		}
		btcRPCConfig = cfg
	}
	return btcRPCConfig
}

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
