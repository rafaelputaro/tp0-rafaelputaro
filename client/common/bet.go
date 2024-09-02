package common

import (
	"encoding/csv"
	"os"

	"github.com/spf13/viper"
)

const MSG_CANT_OPEN_FILE = "Can't open file"
const ACTION_OPEN_FILE = "open_file"
const ACTION_CLOSE_FILE = "close_file"
const ACTION_LOAD_BET = "load_bet_from_file"

// Bet data
type Bet struct {
	Name     string
	LastName string
	DNI      string
	Birthday string
	Number   string
}

func PrintBet(bet *Bet) {
	log.Debugf("action: %s | result: success | name: %s | last_name: %s | dni: %s | birth_day: %s | number: %s",
		ACTION_LOAD_BET,
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
	return &bet
}

func CreateBet(fields []string) *Bet {
	bet := Bet{
		Name:     fields[0],
		LastName: fields[1],
		DNI:      fields[2],
		Birthday: fields[3],
		Number:   fields[4],
	}
	return &bet
}

func LoadBets(v *viper.Viper) *[]Bet {
	var filePath = v.GetString("dataset.file")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("%s %s: %s", MSG_CANT_OPEN_FILE, filePath, err)
	}
	log.Debugf("action: %s %s | result: success ", ACTION_OPEN_FILE, filePath)
	reader := csv.NewReader(file)
	var bets []Bet
loop:
	for {
		line, err := reader.Read()
		if err != nil {
			file.Close()
			log.Debugf("action: %s %s | result: success", ACTION_CLOSE_FILE, filePath)
			break loop
		}
		bets = append(bets, *CreateBet(line))
	}
	return &bets
}
