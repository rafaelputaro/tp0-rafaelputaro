package common

import (
	"errors"
	"fmt"
	"strings"
)

const MAX_LEN = 8192
const MSG_ERROR = "error: message too long"

func doParseBet(idAgency string, bet Bet) string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%s,%v",
		idAgency,
		bet.Name,
		bet.LastName,
		bet.DNI,
		bet.Birthday,
		bet.Number,
	)
}

func ParseBet(idAgency string, bet Bet) (string, error) {
	parsed := doParseBet(idAgency, bet)
	if len(parsed) > MAX_LEN {
		return "", errors.New(MSG_ERROR)
	}
	return parsed, nil
}

// <bet1>;<bet2>,....
func ParseBets(idAgency string, bets *[]Bet) (string, error) {
	var betsParsed []string
	for _, bet := range *bets {
		parsed := doParseBet(idAgency, bet)
		betsParsed = append(betsParsed, parsed)
	}
	var betsParsedStr = strings.Join(betsParsed, ";")
	if len(betsParsedStr) > MAX_LEN {
		return "", errors.New(MSG_ERROR)
	}
	return betsParsedStr, nil
}

// returns: bets parsed, batch_amount_used, error
func do_adaptive_parsing(idAgency string, bets *[]Bet, index int, batch_max_amount int) (string, int, error) {
	var batch_amount_used int
	if len(*bets) > (batch_max_amount + index) {
		batch_amount_used = batch_max_amount
	} else {
		batch_amount_used = len(*bets) - batch_amount_used - index
	}
	for sub := 0; sub < batch_amount_used; sub++ {
		var betsWindow []Bet
		for i := 0; i < batch_amount_used-sub; i++ {
			betsWindow = append(betsWindow, (*bets)[index+i])
		}
		// parse bets
		parsed, errorOnParse := ParseBets(idAgency, &betsWindow)
		if errorOnParse == nil {
			return parsed, batch_amount_used - sub, nil
		}
	}
	return "", 0, errors.New(MSG_ERROR)
}
