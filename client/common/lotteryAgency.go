package common

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// Lottery Agency Entity that encapsulates how
type LotteryAgency struct {
	client *Client
	bets   *[]Bet
}

func CreateNewLotteryAgency(v *viper.Viper) *LotteryAgency {
	clientConfig := ClientConfig{
		ServerAddress:  v.GetString("server.address"),
		ID:             v.GetString("id"),
		LoopAmount:     v.GetInt("loop.amount"),
		LoopPeriod:     v.GetDuration("loop.period"),
		BatchMaxAmount: v.GetInt("batch.maxAmount"),
	}
	agency := LotteryAgency{
		client: NewClient(clientConfig),
		bets:   LoadBets(v),
	}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	agency.client.SendBetsChunks(agency.bets, signalChannel)
	// TODO  agency.client.AskForWinners(signalChannel)
	return &agency
}
