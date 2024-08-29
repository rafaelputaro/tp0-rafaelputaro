package common

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"
	"os"
	"time"

	"github.com/op/go-logging"
)

const SIGNAL_ACTION = "received_a_sigterm"
const SEND_BET_ACTION = "apuesta_enviada"

var log = logging.MustGetLogger("log")

// ClientConfig Configuration used by the client
type ClientConfig struct {
	ID            string
	ServerAddress string
	LoopAmount    int
	LoopPeriod    time.Duration
}

// Client Entity that encapsulates how
type Client struct {
	config ClientConfig
	conn   net.Conn
}

// NewClient Initializes a new client receiving the configuration
// as a parameter
func NewClient(config ClientConfig) *Client {
	client := &Client{
		config: config,
	}
	return client
}

// CreateClientSocket Initializes client socket. In case of
// failure, error is printed in stdout/stderr and exit 1
// is returned
func (c *Client) createClientSocket() error {
	conn, err := net.Dial("tcp", c.config.ServerAddress)
	if err != nil {
		log.Criticalf(
			"action: connect | result: fail | client_id: %v | error: %v",
			c.config.ID,
			err,
		)
	}
	c.conn = conn
	return nil
}

// Send the bets to the server
func (c *Client) SendBets(bets *[]Gambler, singalChannel chan os.Signal) {
loop:
	for _, bet := range *bets {

		// Create the connection the server in every loop iteration
		c.createClientSocket()

		// parse bet
		parsed, errorOnParse := ParseBet(bet)

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

		select {
		case <-singalChannel:
			log.Infof("action: %v | result: success | client_id: %v",
				SIGNAL_ACTION,
				c.config.ID,
			)
			break loop
		case <-time.After(c.config.LoopPeriod):
		}
	}
	log.Infof("action: loop_finished | result: success | client_id: %v", c.config.ID)
}
