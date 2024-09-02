package common

import (
	"bufio"
	"encoding/binary"
	"io"
	"fmt"
)

const SEND_BET_ACTION = "apuesta_enviada"

func apply_protocol(bet Bet, c *Client) {
	// parse bet
	parsedX, errorOnParse := ParseBet(c.config.ID, bet)
	var parsed = fmt.Sprintf("%s;%s\n", parsedX, parsedX)
	println("Parsed ", parsed)
	if errorOnParse != nil {
		log.Errorf("action: parse_bet | result: fail | client_id: %v | error: %v",
			c.config.ID,
			errorOnParse,
		)
	} else {
		binary.Write(c.conn, binary.BigEndian, uint16(len(parsedX)))
		io.WriteString(c.conn, parsed)
		_, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Errorf("action: %v | result: fail | client_id: %v | error: %v",
				SEND_BET_ACTION,
				c.config.ID,
				err,
			)
			return
		}
		log.Infof("action: %v | result: success | dni: %v | numero: %v",
			SEND_BET_ACTION,
			bet.DNI,
			bet.Number,
		)
	}
}
