package svc

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/abec"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/b2node"

	"github.com/qday-io/qday-abel-bridge-committer/pkg/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var svc *ServiceContext

type ServiceContext struct {
	Config *types.Config
	RPC    *ethclient.Client
	DB     *gorm.DB
	// BTCConfig            *types.BitcoinRPCConfig
	B2NodeConfig         *types.B2NODEConfig
	AbecConfig           *types.AbecConfig
	LatestBlockNumber    int64
	SyncedBlockNumber    int64
	SyncedBlockHash      common.Hash
	NodeClient           *b2node.NodeClient
	AbecClient           *abec.AbecClient
	LatestBTCBlockNumber int64
}

func NewServiceContext(cfg *types.Config, b2nodeConfig *types.B2NODEConfig, abecCfg *types.AbecConfig) *ServiceContext {
	storage, err := connectToDB(cfg)
	if err != nil {
		log.Panicf("[svc]gorm get db after retries failed: %s\n", err)
	}

	sqlDB, err := storage.DB()
	if err != nil {
		log.Panicf("[svc]gorm get sqlDB panic: %s\n", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MySQLConnMaxLifetime) * time.Second)

	rpc, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Panicf("[svc] get eth client panic: %s\n", err)
	}

	privateKeHex := b2nodeConfig.PrivateKey
	chainID := b2nodeConfig.ChainID
	address := b2nodeConfig.Address
	grpcConn, err := types.GetClientConnection(b2nodeConfig.GRPCHost, types.WithClientPortOption(b2nodeConfig.GRPCPort))
	if err != nil {
		log.Panicf("[svc] init b2node grpc panic: %s\n", err)
	}

	nodeClient := b2node.NewNodeClient(privateKeHex, chainID, address, grpcConn, b2nodeConfig.CoinDenom)
	abecClient := abec.NewClient(abecCfg.Endpoint, abecCfg.Username, abecCfg.Password, abecCfg.AuthToken, abecCfg.RpcEndpoint)

	svc = &ServiceContext{
		// BTCConfig:         bitcoinCfg,
		DB:                storage,
		Config:            cfg,
		RPC:               rpc,
		LatestBlockNumber: cfg.InitBlockNumber,
		B2NodeConfig:      b2nodeConfig,
		AbecConfig:        abecCfg,
		NodeClient:        nodeClient,
		AbecClient:        abecClient,
	}
	return svc
}
func connectToDB(cfg *types.Config) (*gorm.DB, error) {
	var (
		err     error
		storage *gorm.DB
	)

	operation := func() error {
		storage, err = gorm.Open(mysql.Open(cfg.MySQLDataSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		return err
	}

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 30 * time.Second

	err = backoff.Retry(operation, bo)
	if err != nil {
		return nil, err
	}

	return storage, nil
}
