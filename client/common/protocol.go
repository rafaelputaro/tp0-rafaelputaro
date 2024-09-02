package common

import (
	"bufio"
	"encoding/binary"
	"io"
)

const SEND_BET_ACTION = "apuesta_enviada"

func apply_bets_protocol(bets *[]Bet, index int, c *Client) (int, error) {
	// parse bets
	parsed, batch_amount_used, errorOnParse := do_adaptive_parsing(c.config.ID, bets, index, c.config.BatchMaxAmount)
	if errorOnParse != nil {
		log.Errorf("action: parse_bets | result: fail | client_id: %v | error: %v",
			c.config.ID,
			errorOnParse,
		)
		return index, errorOnParse
	} else {
		// send message to server
		// how many bets in the batch
		binary.Write(c.conn, binary.BigEndian, uint16(batch_amount_used))
		// how many bytes batch
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
			return index, err
		}
		log.Infof("action: %v | result: success | batch_amount: %v", SEND_BET_ACTION, batch_amount_used)
	}
	return index + batch_amount_used, nil
}
