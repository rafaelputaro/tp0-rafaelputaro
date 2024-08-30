package common

import (
	"bufio"
	"encoding/binary"
	"io"
)

func apply_protocol(bet Bet, c *Client) {
	// parse bet
	parsed, errorOnParse := ParseBet(c.config.ID, bet)

	if errorOnParse != nil {
		log.Errorf("action: parse_bet | result: fail | client_id: %v | error: %v",
			c.config.ID,
			errorOnParse,
		)
	} else {
		// send message to server: <len><bet parsed>
		binary.Write(c.conn, binary.BigEndian, uint16(len(parsed)))
		io.WriteString(c.conn, parsed)
		_, err := bufio.NewReader(c.conn).ReadString('\n')
		c.conn.Close()

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
