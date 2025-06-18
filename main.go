package main

import (
	"github.com/b2network/b2committer/pkg/log"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/handler"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/svc"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/types"
	"github.com/shopspring/decimal"
)

func main() {
	decimal.DivisionPrecision = 18
	cfg := types.GetConfig()
	btccfg := types.GetBtcConfig()
	b2nodeConfig := types.GetB2nodeConfig()
	abecCfg := types.GetAbecConfig()
	log.Init(cfg.LogLevel, cfg.LogFormat)
	log.Infof("config: %v\n", cfg)
	ctx := svc.NewServiceContext(cfg, btccfg, b2nodeConfig, abecCfg)
	handler.Run(ctx)
	log.Info("listener running...\n")
	select {}
}
