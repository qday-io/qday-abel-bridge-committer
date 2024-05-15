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
	RPCUrl               string `env:"RPC_URL" envDefault:"http://124.243.137.251:8888"`
	Blockchain           string `env:"BLOCKCHAIN" envDefault:"b2-node"`
	InitBlockNumber      int64  `env:"INIT_BLOCK_NUMBER" envDefault:"32332"`
	InitBlockHash        string `env:"INIT_BLOCK_HASH" envDefault:"0xb20dc0f9df4e923ac60bb210f0417cf2fce6c3e5441a61944cfc6561b444a96b"`
	PolygonZKEVMAddress  string `env:"POLYGON_ZKEVM_ADDRESS" envDefault:"0x8aCd85898458400f7Db866d53FCFF6f0D49741FF"`
	LimitNum             int    `evn:"PROPOSAL_BATCHES_LIMITNUM" envDefault:"10"`
	InitProposalID       uint64 `evn:"INIT_PROPOSAL_ID" envDefault:"1"`
}

type B2NODEConfig struct {
	Address    string `env:"B2NODE_ADDRESS" envDefault:"ethm10pax3a2aqxc33s3ya85e25us05y7pfmcmkd20p"`
	ChainID    string `env:"B2NODE_CHAIN_ID" envDefault:"ethermint_9000-1"`
	GRPCHost   string `env:"B2NODE_GRPC_HOST" envDefault:"124.243.137.251"`
	GRPCPort   uint32 `env:"B2NODE_GRPC_PORT" envDefault:"9199"`
	RPCUrl     string `env:"B2NODE_RPC_URL" envDefault:"http://124.243.137.251:8888"`
	CoinDenom  string `env:"B2NODE_COIN_DENOM" envDefault:"aphoton"`
	PrivateKey string `env:"B2NODE_PRIVATE_KEY" envDefault:"C37E6DB22B966D4091B10F86B1ED63F1C3F1372525946DBAF3F31B1A18A567E5"`
}

type BitcoinRPCConfig struct {
	NetworkName        string `env:"BITCOIN_NETWORK_NAME" envDefault:"signet"`
	PrivateKey         string `env:"BITCOIN_PRIVATE_KEY" envDefault:"L2Wv2kPRs4iJ8zFwQmTPBvrPQY7DS591LdmTEqu4x6GXX4yVDr7h"`
	DestinationAddress string `env:"COMMITTER_DESTINATION_ADDRESS" envDefault:"bc1peuj9pfr4leqnfmem4nrsxl2cqcz2hafw7k6luxzrxcr545fywm3qrsfvtj"`
}

type AbecConfig struct {
	Endpoint string `env:"ABEC_ENDPOINT" envDefault:"https://testnet-rpc-exchange.abelian.info"`
	Username string `env:"ABEC_USERNAME" envDefault:"KFf5krbZiLyfo5KaIsNb3Fr2QZs="`
	Password string `env:"ABEC_PASSWORD" envDefault:"M+DxFwon2FYyiLgaJoTZ9qCr6Jc="`
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
