import socket
import logging
import sys
import signal
import multiprocessing
from common.national_lottery import NationalLottery
from common.protocol import apply_ask_protocol, apply_rcv_bet_protocol, apply_res_bet_protocol

SIGNAL_HANDLER_ACTION="received_a_signal"
CLOSE_SERVER_SOCKET_ACTION="closing_server_socket"
CLOSE_SOCKET_ACTION="closing_a_client_socket"
JOIN_CHILD_ACTION="join_child_process"
SIZE_TAG = 4

class Server:
    def __init__(self, port, listen_backlog):
        # Initialize server socket
        self._server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self._server_socket.bind(('', port))
        self._server_socket.listen(listen_backlog)
        self.__init_sign_handling()
        self._childs_processses = []
        self._processes_manager = multiprocessing.Manager()
        self._lottery = NationalLottery(self._processes_manager.Lock(), self._processes_manager.Lock())

    """
    Inicialización de manejo de señales
    """
    def __init_sign_handling(self):
        signal.signal(signal.SIGINT, self.__handle_a_signal)
        signal.signal(signal.SIGTERM, self.__handle_a_signal)

    """
    Manejo de la señal
    """
    def __handle_a_signal(self, signal_number, _stack):
        logging.info(f'action: {SIGNAL_HANDLER_ACTION} | signal_number: {signal_number}')
        self._server_socket.close()
        logging.debug(f'action: {CLOSE_SERVER_SOCKET_ACTION} | result: sucess')
        for child in self._childs_processses:
            child.join()
            logging.debug(f'action: {JOIN_CHILD_ACTION} | result: sucess')
        sys.exit(0)

    def run(self):
        """
        Dummy Server loop

        Server that accept a new connections and establishes a
        communication with a client. After client with communucation
        finishes, servers starts to accept new connections again
        """

        # TODO: Modify this program to handle signal to graceful shutdown
        # the server
        while True:
            client_sock = self.__accept_new_connection()
            process = multiprocessing.Process(
                target=self.__handle_client_connection, args=(client_sock,))
            process.start()
            self._childs_processses.append(process)            

    def __handle_client_connection(self, client_sock):
        """
        Read message from a specific client socket and closes the socket

        If a problem arises in the communication with the client, the
        client socket will also be closed
        """
        try:
            # read tag
            tag = client_sock.recv(SIZE_TAG).decode('utf-8').strip()
            if tag == "bets":
                amount = apply_rcv_bet_protocol(client_sock, self._lottery)
                apply_res_bet_protocol(client_sock, amount)
            if tag == "asks":
                apply_ask_protocol(client_sock, self._lottery)
        except OSError as e:
            logging.error("action: receive_message | result: fail | error: {e}")
        finally:
            client_sock.close()

    def __accept_new_connection(self):
        """
        Accept new connections

        Function blocks until a connection to a client is made.
        Then connection created is printed and returned
        """

        # Connection arrived
        logging.info('action: accept_connections | result: in_progress')
        c, addr = self._server_socket.accept()
        logging.info(f'action: accept_connections | result: success | ip: {addr[0]}')
        return c
