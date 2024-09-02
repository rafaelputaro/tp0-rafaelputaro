package common

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

// Lottery Agency Entity that encapsulates how
type LotteryAgency struct {
	client *Client
	bet    *Bet
}

func CreateNewLotteryAgency(v *viper.Viper) *LotteryAgency {
	clientConfig := ClientConfig{
		ServerAddress: v.GetString("server.address"),
		ID:            v.GetString("id"),
		LoopAmount:    v.GetInt("loop.amount"),
		LoopPeriod:    v.GetDuration("loop.period"),
	}
	agency := LotteryAgency{
		client: NewClient(clientConfig),
		bet:    LoadBet(v),
	}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	// send bet
	agency.client.SendBet(*agency.bet)
	// waiting signal
loop:
	for {
		select {
		case <-signalChannel:
			log.Infof("action: %v | result: success | client_id: %v",
				SIGNAL_ACTION,
				clientConfig.ID,
			)
			break loop
		case <-time.After(clientConfig.LoopPeriod):
		}
	}
	return &agency
}
