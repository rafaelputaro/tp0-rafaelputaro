package common

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// Lottery Agency Entity that encapsulates how
type LotteryAgency struct {
	client  *Client
	gambler *Gambler
}

func CreateNewLotteryAgency(v *viper.Viper) *LotteryAgency {
	clientConfig := ClientConfig{
		ServerAddress: v.GetString("server.address"),
		ID:            v.GetString("id"),
		LoopAmount:    v.GetInt("loop.amount"),
		LoopPeriod:    v.GetDuration("loop.period"),
	}
	agency := LotteryAgency{
		client:  NewClient(clientConfig),
		gambler: LoadGambler(v),
	}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	// load a gambler
	var bets []Gambler
	bets = append(bets, *agency.gambler)
	// send bet
	agency.client.SendBets(&bets, signalChannel)
	return &agency
}
