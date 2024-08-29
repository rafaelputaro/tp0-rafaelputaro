package common

import (
	"github.com/spf13/viper"
)

// Gambler data
type Gambler struct {
	Name     string
	LastName string
	DNI      string
	Birthday string
	Number   string
}

func printGambler(gambler *Gambler) {
	log.Debugf("action: gambler_config | result: success | name: %s | last_name: %s | dni: %s | birth_day: %s | number: %s",
		gambler.Name,
		gambler.LastName,
		gambler.DNI,
		gambler.Birthday,
		gambler.Number,
	)
}

func LoadGambler(v *viper.Viper) *Gambler {
	gambler := Gambler{
		Name:     v.GetString("NOMBRE"),
		LastName: v.GetString("APELLIDO"),
		DNI:      v.GetString("DOCUMENTO"),
		Birthday: v.GetString("NACIMIENTO"),
		Number:   v.GetString("NUMERO"),
	}
	printGambler(&gambler)
	return &gambler
}
