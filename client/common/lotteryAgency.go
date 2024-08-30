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
	// load a bet
	var bets []Bet
	bets = append(bets, *agency.bet)
	// send bet
	agency.client.SendBets(&bets, signalChannel)
	return &agency
}
