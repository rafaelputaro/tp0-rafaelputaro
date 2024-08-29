package common

import (
	"errors"
	"fmt"
)

const MAX_LEN = 8192
const MSG_ERROR = "error: message too long"

func doParseBet(bet Gambler) string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%v\n",
		bet.Name,
		bet.LastName,
		bet.DNI,
		bet.Birthday,
		bet.Number,
	)
}

func ParseBet(bet Gambler) (string, error) {
	parsed := doParseBet(bet)
	if len(parsed) > MAX_LEN {
		return "", errors.New(MSG_ERROR)
	}
	return parsed, nil
}
