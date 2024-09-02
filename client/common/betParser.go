package common

import (
	"errors"
	"fmt"
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
