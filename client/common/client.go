package common

import (
	"net"
	"os"
	"time"
	"github.com/op/go-logging"
)

const SIGNAL_ACTION = "received_a_sigterm"
const ACTION_WINNER_PROTOCOL = "consultar_ganadores"

var log = logging.MustGetLogger("log")

// ClientConfig Configuration used by the client
type ClientConfig struct {
	ID             string
	ServerAddress  string
	LoopAmount     int
	LoopPeriod     time.Duration
	BatchMaxAmount int
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
func (c *Client) SendBetsChunks(bets *[]Bet, signalChannel chan os.Signal) {
	var index = 0
loop:
	for index < len(*bets) {
		// Create the connection the server in every loop iteration
		c.createClientSocket()
		index, _ = apply_bets_protocol(bets, index, c)
		// handle a signal
		select {
		case <-signalChannel:
			log.Infof("action: %v | result: success | client_id: %v",
				SIGNAL_ACTION,
				c.config.ID,
			)
			break loop
		case <-time.After(c.config.LoopPeriod):
		}
	}
	println(c.config.BatchMaxAmount)
	log.Infof("action: loop_finished | result: success | client_id: %v", c.config.ID)
}

// Send the bets to the server
func (c *Client) AskForWinners(agencyId string, signalChannel chan os.Signal) {
	loop:
	for {
		// Create the connection the server in every loop iteration
		c.createClientSocket()
		err, winners_rcv := apply_winners_protocol(agencyId, c)
		if err != nil {
			log.Errorf("action: %v | result: fail | client_id: %v | error: %v",
				ACTION_WINNER_PROTOCOL,
				c.config.ID,
				err,
			)
		}
		if winners_rcv {
			break loop
		}
		// handle a signal
		select {
		case <-signalChannel:
			log.Infof("action: %v | result: success | client_id: %v",
				SIGNAL_ACTION,
				c.config.ID,
			)
			break loop
		case <-time.After(c.config.LoopPeriod):
		}
	}
	log.Infof("action: loop_winners_finished | result: success | client_id: %v", c.config.ID)
}
