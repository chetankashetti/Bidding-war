package service

import (
	game_abi "github.com/blockchain-challenge/cmd/game-abi"
	"github.com/blockchain-challenge/cmd/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type BidwarService struct {
	ethClient EthClient
	Contract  common.Address
	Instance  *game_abi.BidwarBind
}

func (s *BidwarService) RaiseNewBid(bidRequest *types.BidRequest) (gasSpent uint64, err error) {
	hexToAddress := common.HexToAddress(bidRequest.WalletAddress)
	auth := GetAccountAuth(s.ethClient.Client, hexToAddress.String())
	newBid, err := s.Instance.NewBid(auth, &bidRequest.Amount, s.Contract)
	gasSpent = newBid.Gas()
	if err != nil {
		log.Fatalf("Failed to call new bid method: %v", err)
		return
	}
	return
}

func (s *BidwarService) DecideWinnerAndTransfer() (gasSpent uint64, err error) {
	finalise, err := s.Instance.Finalise(&bind.TransactOpts{}, s.Contract)
	gasSpent = finalise.Gas()
	if err != nil {
		log.Fatalf("Failed to call new bid method: %v", err)
		return
	}
	return
}

func NewBidwarService(ethClient *EthClient, contract string) *BidwarService {
	hexToAddress := common.HexToAddress(contract)
	instance, err := game_abi.NewBidwarBind(hexToAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to connect to ethereum contract: %v", err)
	}

	return &BidwarService{ethClient: *ethClient, Contract: hexToAddress, Instance: instance}
}
