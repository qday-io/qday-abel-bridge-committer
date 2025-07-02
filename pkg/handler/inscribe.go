package handler

import (
	"time"

	"github.com/qday-io/qday-abel-bridge-committer/pkg/inscribe"

	"github.com/qday-io/qday-abel-bridge-committer/pkg/log"

	"github.com/pkg/errors"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/schema"
	"github.com/qday-io/qday-abel-bridge-committer/pkg/svc"
)

// Inscribe check proposal statues. process pending proposal.
func Inscribe(ctx *svc.ServiceContext) {
	time.Sleep(30 * time.Second)
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status=?", schema.PendingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.Inscribe] Pending and timeout proposal err: %s\n", errors.WithStack(err).Error())
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[CheckProposalPending] QueryProposalByID err: %s\n", errors.WithStack(err).Error())
			continue
		}
		if proposal.Status == schema.SucceedStatus {
			dbProposal.Status = proposal.Status
			dbProposal.Winner = proposal.Winner
			dbProposal.BtcRevealTxHash = proposal.BitcoinTxHash
			ctx.DB.Save(dbProposal)
		}
		if proposal.Status == schema.PendingStatus &&
			proposal.Winner == ctx.B2NodeConfig.Address && proposal.BitcoinTxHash == "" {
			memo, err := inscribe.GenerateMemoData(ctx.AbecConfig.From, proposal.StateRootHash, proposal.ProofHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] GenerateMemoData err: %s\n", errors.WithStack(err).Error())
				continue
			}

			time.Sleep(5 * time.Second)
			log.Infof("[Handler.Inscribe] Inscribe to abec, memo: %v\n", proposal.Id)
			abecTxHash, err := ctx.AbecClient.UserTransferToSingleRecipient(ctx.AbecConfig, memo, "10000")
			if err != nil {
				log.Errorf("[Handler.Inscribe] UserTransferToSingleRecipient err: %s\n", errors.WithStack(err).Error())
				continue
			}

			log.Infof("[Handler.Inscribe] Inscribe to abec [start], abecTxHash: %s\n", abecTxHash)
			_, err = ctx.NodeClient.BitcoinTx(proposal.Id, proposal.Winner, abecTxHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err).Error())
				continue
			}
			log.Infof("[Handler.Inscribe] Inscribe to abec [end], abecTxHash: %s\n", abecTxHash)

			dbProposal.BtcRevealTxHash = abecTxHash

			ctx.DB.Save(dbProposal)
		}
		if proposal.Status == schema.PendingStatus && proposal.BitcoinTxHash != "" && proposal.Winner != ctx.B2NodeConfig.Address {
			// 确认状态 对比大于6个高度后 确认后就提交提案
			confirmed, blockHeight, err := ctx.AbecClient.GetTxConfirmedStatus(proposal.BitcoinTxHash, ctx.AbecConfig.APPID,
				ctx.AbecConfig.UserID, ctx.AbecConfig.RequestSignature)

			if err != nil {
				log.Errorf("[Handler.Inscribe] GetTxConfirmedStatus err: %s\n", errors.WithStack(err).Error())
				continue
			}

			if confirmed && (ctx.LatestBTCBlockNumber-blockHeight) >= 6 {
				_, err = ctx.NodeClient.BitcoinTx(proposal.Id, proposal.Proposer, proposal.BitcoinTxHash)
				if err != nil {
					log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err).Error())
					continue
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}
