import logging

from common.utils import Bet
from common.national_lottery import NationalLottery

ACTION_RECEIVE = "apuesta recibida"

"""
returns:
    * batch amount in case recieve
"""
def apply_rcv_bet_protocol(client_sock, lottery: NationalLottery) -> int:
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

def apply_res_bet_protocol(client_sock, amount):

    client_sock.send("{}\n".format(amount).encode('utf-8'))

def apply_ask_protocol(client_sock, lottery: NationalLottery) :
    # id agency length
    length = int.from_bytes(client_sock.recv(2), byteorder='big')
    # id agency
    agency_id = client_sock.recv(length).decode('utf-8').strip()
    # notify ends
    lottery.notify_agency_ends(agency_id)
    # ¿lottery time?
    if lottery.all_agencies_ended() :
        client_sock.send("{}\n".format("winners").encode('utf-8'))
        winners: list[str] = lottery.get_winners(agency_id)
        client_sock.send("{}\n".format(','.join(winners)).encode('utf-8'))
    else :
        client_sock.send("{}\n".format("bets").encode('utf-8'))
     