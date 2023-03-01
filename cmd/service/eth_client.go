package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
	"time"
)

type EthClient struct {
	*ethclient.Client
	rpcClient      *rpc.Client
	contextTimeout time.Duration
}

func NewEthClient(endpoint string) *EthClient {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		log.Fatalf("Failed to connect to ethereum client: %v", err)
	}

	return &EthClient{
		Client:         ethclient.NewClient(client),
		rpcClient:      client,
		contextTimeout: 60,
	}
}
