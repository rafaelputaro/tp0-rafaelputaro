import logging
from common.utils import Bet
from common.utils import store_bets as do_store_bets
from common.utils import has_won as do_has_won

ACTION_STORE = 'apuesta_almacenada'
MAX_AGENCIES = '5'

class NationalLottery:
    def __init__(self):
        self._agencies_ended = set()
       
    def store_bet(self, bet: Bet):
        logging.info(f'action: {ACTION_STORE} | result: success | dni: {bet.document} | numero: {bet.number}')       
        self.store_bets([bet])

    def store_bets(self, bets):
        do_store_bets(bets)
        
    def notify_agency_ends(self, agency: int) :
        return self._agencies_ended.add(agency)
    
    def all_agencies_ended(self) -> bool:
        return len(self._agencies_ended) >= MAX_AGENCIES

    def has_won(bet: Bet) -> bool:
        return do_has_won(bet)