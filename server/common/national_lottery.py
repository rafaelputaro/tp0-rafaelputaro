import logging
from common.utils import Bet
from common.utils import store_bets as do_store_bets
from common.utils import has_won, load_bets

ACTION_STORE = 'apuesta_almacenada'
MAX_AGENCIES = 5

class NationalLottery:
    def __init__(self):
        self._agencies_ended = set()
       
    def store_bet(self, bet: Bet):
        logging.info(f'action: {ACTION_STORE} | result: success | dni: {bet.document} | numero: {bet.number}')       
        self.store_bets([bet])

    def store_bets(self, bets):
        do_store_bets(bets)
        
    def notify_agency_ends(self, agency: str) :
        return self._agencies_ended.add(agency)
    
    def all_agencies_ended(self) -> bool:
        return len(self._agencies_ended) >= MAX_AGENCIES

    def get_winners(self, agency_id) -> list[str]:
        winners: list[str] = []
        bets: list[Bet] = load_bets()
        for bet in bets:
            if has_won(bet) and (int(bet.agency) == int(agency_id)):
                winners.append(bet.document)
        return winners

