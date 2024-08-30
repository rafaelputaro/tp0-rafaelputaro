package common

import (
	"github.com/spf13/viper"
)

// Bet data
type Bet struct {
	Name     string
	LastName string
	DNI      string
	Birthday string
	Number   string
}

func printBet(bet *Bet) {
	log.Debugf("action: bet_config | result: success | name: %s | last_name: %s | dni: %s | birth_day: %s | number: %s",
		bet.Name,
		bet.LastName,
		bet.DNI,
		bet.Birthday,
		bet.Number,
	)
}

func LoadBet(v *viper.Viper) *Bet {
	bet := Bet{
		Name:     v.GetString("NOMBRE"),
		LastName: v.GetString("APELLIDO"),
		DNI:      v.GetString("DOCUMENTO"),
		Birthday: v.GetString("NACIMIENTO"),
		Number:   v.GetString("NUMERO"),
	}
	printBet(&bet)
	return &bet
}
