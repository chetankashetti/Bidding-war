package controller

import (
	"github.com/blockchain-challenge/cmd/service"
	"github.com/blockchain-challenge/cmd/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BidwarController struct {
	ethClient     ethclient.Client
	bidWarService service.BidwarService
}

func (bw BidwarController) NewBidReceived(c *gin.Context) {
	var bidAmount types.BidRequest
	err := c.ShouldBindJSON(&bidAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	gasSpent, err := bw.bidWarService.RaiseNewBid(&bidAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.RespMessage{
			Message:  err.Error(),
			GasSpent: gasSpent,
		})
		return
	}

	c.JSON(http.StatusOK, types.RespMessage{
		Message:  "New Bid is Successful.",
		GasSpent: gasSpent,
	})
}

func (bw BidwarController) Timeout(c *gin.Context) {
	gasSpent, err := bw.bidWarService.DecideWinnerAndTransfer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.RespMessage{
			Message:  err.Error(),
			GasSpent: gasSpent,
		})
		return
	}

	c.JSON(http.StatusOK, types.RespMessage{
		Message:  "Finalising Bid is Successful.",
		GasSpent: gasSpent,
	})
}

func NewBidwarController(client *ethclient.Client, bidwarService *service.BidwarService) *BidwarController {
	return &BidwarController{ethClient: *client, bidWarService: *bidwarService}
}
