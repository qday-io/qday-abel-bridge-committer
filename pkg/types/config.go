package types

import (
	"github.com/qday-io/qday-abel-bridge-committer/pkg/log"
	"github.com/caarlos0/env/v6"
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
	RPCUrl               string `env:"RPC_URL" envDefault:"http://124.243.132.119:8545"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"15000"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0xa8489891c72afd883114141653e8529a2175cd516ce825c118c05cc76a8be045"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0xdde8f57d008D36C9476c5DF5525F7C5B52B03647"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"2"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	Address    string `env:"B2NODE_ADDRESS" envDefault:"ethm14vggzqz4nlyhzhr36m5fmskdqms3q5p6xp3vrl"`
	ChainID    string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost   string `env:"B2NODE_GRPC_HOST" envDefault:"124.243.132.119"`
	GRPCPort   uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9090"`
	RPCUrl     string `env:"B2NODE_RPC_URL" envDefault:"http://124.243.132.119:9090"`
	CoinDenom  string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton"`
	PrivateKey string `env:"B2NODE_PRIVATE_KEY" envDefault:"F07814758C2CF31ED1D7D1C696896B4EEC57D4A1837E7A10E4D5100E23620715"`
}

type BitcoinRPCConfig struct {
	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"L2Wv2kPRs4iJ8zFwQmTPBvrPQY7DS591LdmTEqu4x6GXX4yVDr7h"`
	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"bc1peuj9pfr4leqnfmem4nrsxl2cqcz2hafw7k6luxzrxcr545fywm3qrsfvtj"`
}

type AbecConfig struct {
	Endpoint         string `env:"ENDPOINT" envDefault:"https://testnet-snode.abelian.info/v1/single-account"`
	RpcEndpoint      string `env:"RPCENDPOINT" envDefault:"https://testnet-rpc-00.abelian.info"`
	Username         string `env:"USERNAME" envDefault:"J8y0OnkS2wx9XEgUlW5MqtoRDAQ="`
	Password         string `env:"PASSWORD" envDefault:"ULlXc/ZMJ375cn6VuSbtU+Y3KGQ="`
	APPID            string `env:"APPID" envDefault:"cce71078"`
	RequestSignature string `env:"REQUEST_SIGNATURE" envDefault:"0x338i3jejjd"`
	UserID           string `env:"USERID" envDefault:"abe3b614871d6db00503f6e8108260f1943a3a70c09557cbb7dbce3df9a411fdb9b161572931a185fb853ccc8d2f2a6f12d8cf295b40659e2e1608650783c7ecdc78"`
	From             string `env:"FROM" envDefault:"abe3b614871d6db00503f6e8108260f1943a3a70c09557cbb7dbce3df9a411fdb9b161572931a185fb853ccc8d2f2a6f12d8cf295b40659e2e1608650783c7ecdc78"`
	Recipient        string `env:"RECIPIENT" envDefault:"abe338491ef250a530f6b1a771d45ae168f81d6a430f20623849e448b870f0f95e13f12ba51bff83497480db944567750e3cf555cd9811db95b848ca93d45c1448d0"`
	PrivateKey       string `env:"PRIVATE_KEY" envDefault:"000000009e7521a87e1030fec16e2c1bf46b4f49976c6a4db43e64373875cce74b69385d016b8fd9c17dc62d570389bf9686952d5528420b046d38d5d2737e1598e4cdbcc4de0165cfb8953a17b79af4b3b9da9de7d39f8806eacfeea9fae965c07bc561074e7b2fa2a76baf2eb41006b6fb075f41249f923bb5c7f98b1db92967b1c2b5"`
	AuthToken        string `env:"AUTHTOKEN" envDefault:"cce71078669ded3517d961a2d57eb440"`
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
