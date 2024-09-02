import logging

from common.utils import Bet
from common.national_lottery import NationalLottery

ACTION_RECEIVE = "apuesta recibida"

"""
returns:
    * batch amount in case recieve
"""
def apply_rcv_protocol(client_sock, lottery) :
    # message length
    batch_amount = int.from_bytes(client_sock.recv(2), byteorder='big')
    # message length
    length = int.from_bytes(client_sock.recv(2), byteorder='big')
    # recive message
    msg = client_sock.recv(length).decode('utf-8').strip()
    # split in bets
    bets_msg = msg.split(';')
    # store bets
    for bet_msg in bets_msg:
        bet = Bet(*bet_msg.split(','))
        lottery.store_bet(bet)
    # how many bets
    amount = len(bets_msg)
    if (batch_amount == amount) :
        logging.info(
            f'action: {ACTION_RECEIVE} | result: success | cantidad: {batch_amount}')
    else:
        logging.error(
            f'action: {ACTION_RECEIVE} | result: fail | cantidad: {batch_amount}')
    return amount

def apply_res_protocol(client_sock, amount):

    client_sock.send("{}\n".format(amount).encode('utf-8'))