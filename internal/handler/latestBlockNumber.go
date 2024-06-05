package handler

import (
	"context"
	"time"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/internal/svc"
)

func LatestBlackNumber(ctx *svc.ServiceContext) {
	for {
		latest, err := ctx.RPC.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("[Handle.LatestBlackNumber]Syncing latest block number error: %s\n", err)
			time.Sleep(3 * time.Second)
			continue
		}
		ctx.LatestBlockNumber = int64(latest)
		log.Infof("[Handle.LatestBlackNumber] Syncing latest block number: %d \n", latest)

		abecLatest, err := ctx.AbecClient.GetBestBlockHeight()
		if err != nil {
			log.Errorf("[Handle.LatestBTCBlackNumber]Syncing btc network latest block number error: %s\n", err)
			time.Sleep(3 * time.Second)
			continue
		}
		ctx.LatestBTCBlockNumber = abecLatest
		log.Infof("[Handle.LatestBTCBlackNumber] Syncing btc network latest block number: %d \n", abecLatest)
		time.Sleep(3 * time.Second)
	}
}
