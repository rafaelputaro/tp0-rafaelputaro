import logging

from common.utils import Bet

ACTION_RECEIVE = "receive_message"

def apply_rcv_protocol(client_sock):
    
    length = int.from_bytes(client_sock.recv(2), byteorder='big')
    
    msg = client_sock.recv(length).decode('utf-8').strip()
    
    logging.debug(
        f'action: {ACTION_RECEIVE} | result: success | msg: {msg}')

    return Bet(*msg.split(','))

def apply_res_protocol(client_sock, bet):

    client_sock.send("{}\n".format(bet.number).encode('utf-8'))