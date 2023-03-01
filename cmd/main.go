package main

import (
	"errors"
	"fmt"
	"github.com/blockchain-challenge/cmd/config"
	"github.com/blockchain-challenge/cmd/controller"
	"github.com/blockchain-challenge/cmd/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	log.Info("Application getting started.!")
	err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("unable to load config: %s", err))
	}
	// setup eth client
	ethClient := service.NewEthClient(viper.GetString("app.node_endpoint"))

	// setup service
	bidWarService := service.NewBidwarService(ethClient, "app.contract_address")

	//setup controller
	bidWarController := controller.NewBidwarController(ethClient.Client, bidWarService)

	router := gin.Default()
	router.NoRoute(NoRouteHandler())
	api := router.Group("/api/")

	api.POST("/bid-war/bid", bidWarController.NewBidReceived)
	api.POST("/bid-war/timeout", bidWarController.Timeout)
	startGinServer(router, viper.GetString("app.app_name"), viper.GetInt64("app.port"))
}

func NoRouteHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func startGinServer(router *gin.Engine, name string, port int64) {
	err := router.Run(fmt.Sprintf(":%d", port))
	if err == nil {
		log.Infof("%s started server successfully", name)
	} else {
		panic(errors.New("cannot start http server. error: " + err.Error()))
	}
}
