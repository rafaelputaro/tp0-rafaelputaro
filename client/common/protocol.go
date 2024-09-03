package common

import (
	"bufio"
	"encoding/binary"
	"io"
	"strconv"
	"strings"
)

const BETS_TAG = "bets"
const ASK_TAG = "asks"
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
		// send tag
		io.WriteString(c.conn, BETS_TAG)
		// how many bets in the batch
		binary.Write(c.conn, binary.BigEndian, uint16(batch_amount_used))
		// how many bytes batch
		binary.Write(c.conn, binary.BigEndian, uint16(len(parsed)))
		io.WriteString(c.conn, parsed)
		batch_amount_rcv_str, err := bufio.NewReader(c.conn).ReadString('\n')
		c.conn.Close()
		if err != nil {
			log.Errorf("action: %v | result: fail | client_id: %v | error: %v",
				SEND_BET_ACTION,
				c.config.ID,
				err,
			)
			return index, err
		} else {
			var batch_amount_rcv, err_parse = strconv.Atoi(strings.Split(batch_amount_rcv_str, "\n")[0])
			if err_parse != nil {
				log.Errorf("action: %v | result: fail | client_id: %v | error: %v",
					SEND_BET_ACTION,
					c.config.ID,
					err_parse,
				)
			} else {
				if batch_amount_rcv == batch_amount_used {
					log.Infof("action: %v | result: success | batch_amount: %v", SEND_BET_ACTION, batch_amount_used)
				} else {
					log.Errorf("action: %v | result: fail | client_id: %v | cantidad_enviada: %v | cantidad_recibida: %v",
						SEND_BET_ACTION,
						c.config.ID,
						batch_amount_used,
						batch_amount_rcv,
					)
				}
			}
		}
	}
	return index + batch_amount_used, nil
}
