package handler

import (
	"github.com/qday-io/qday-abel-bridge-committer/pkg/svc"
)

func Run(ctx *svc.ServiceContext) {
	// query last block number
	go LatestBlackNumber(ctx)
	// sync blocks
	go SyncBlock(ctx)
	// sync events
	go SyncEvent(ctx)
	// execute committer
	go Committer(ctx)
	// check status
	go CheckStatus(ctx)
	// check and inscribe
	go Inscribe(ctx)
	// check time out
	go CheckStatusTimeOut(ctx)
	// sync proposal
	go SyncProposal(ctx)
}
