import logging
from common.utils import Bet
from common.utils import store_bets as do_store_bets
from common.utils import has_won, load_bets

ACTION_STORE = 'apuesta_almacenada'
MAX_AGENCIES = 5

class NationalLottery:
    def __init__(self, lock_store_bet, lock_agencies_ended, shared_agencies_data):
        self._shared_agencies_data = shared_agencies_data
        self._lock_store_bet = lock_store_bet
        self._lock_agencies_ended = lock_agencies_ended
       
    def store_bet(self, bet: Bet):
        logging.info(f'action: {ACTION_STORE} | result: success | dni: {bet.document} | numero: {bet.number}')       
        self.store_bets([bet])

    def store_bets(self, bets):
        with self._lock_store_bet:
            do_store_bets(bets)
        
    def notify_agency_ends(self, agency: str) :
        with self._lock_agencies_ended:
            if agency not in self._shared_agencies_data['agencies_ended']:
                self._shared_agencies_data['agencies_ended'] = self._shared_agencies_data['agencies_ended'] + [agency]
                self._shared_agencies_data['all_agencies_ended'] = len(self._shared_agencies_data['agencies_ended']) >= MAX_AGENCIES
    
    def all_agencies_ended(self) -> bool:
        with self._lock_agencies_ended:
            return self._shared_agencies_data['all_agencies_ended']

    def get_winners(self, agency_id) -> list[str]:
        winners: list[str] = []
        bets: list[Bet] = load_bets()
        for bet in bets:
            if has_won(bet) and (int(bet.agency) == int(agency_id)):
                winners.append(bet.document)
        return winners

