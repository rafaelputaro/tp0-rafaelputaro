import logging
from common.utils import Bet
from common.utils import store_bets as do_store_bets

ACTION_STORE = 'apuesta_almacenada'

class NationalLottery:
    def __init__(self):
        pass

    def store_bet(self, bet: Bet):
        logging.info(f'action: {ACTION_STORE} | result: success | dni: {bet.document} | numero: {bet.number}')       
        self.store_bets([bet])

    def store_bets(self, bets):
        do_store_bets(bets)
        
    
    