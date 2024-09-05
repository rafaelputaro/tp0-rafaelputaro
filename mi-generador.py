import sys

ERROR_MESSAGE = "No se han ingresado los argumentos correctos: <archivo de salida> <cantidad clientes>"
HEADER = "name: tp0\nservices:\n  server:\n    container_name: server\n    image: server:latest\n    entrypoint: python3 /main.py\n    environment:\n      - PYTHONUNBUFFERED=1\n      - LOGGING_LEVEL=DEBUG\n    networks:\n      - testing_net\n    volumes:\n      - ./server/config.ini:/config.ini\n"
FOOTER = "networks:\n  testing_net:\n    ipam:\n      driver: default\n      config:\n        - subnet: 172.25.125.0/24\n"

def do_generate_compose_fyle(args):
    file = open(args["file"], "w+")
    file.write(HEADER)
    for client_id in range(1, int(args["clients"])+1):
        file.write(generate_client_code(client_id))
    file.write(FOOTER)
    file.close()

def generate_client_code(client_id):
    return f'  client{client_id}:\n    container_name: client{client_id}\n    image: client:latest\n    entrypoint: /client\n    environment:\n      - CLI_ID={client_id}\n      - CLI_LOG_LEVEL=DEBUG\n    networks:\n      - testing_net\n    depends_on:\n      - server\n    volumes:\n      - ./client/config.yaml:/config.yaml\n'

def generate_compose_fyle():
    args = parse_args()
    if args:
        do_generate_compose_fyle(args)

def parse_args():
    if len(sys.argv) >=3:
        return {
            "file": sys.argv[1],
            "clients": sys.argv[2]
        }
    else:
        print(ERROR_MESSAGE)
        return
    
generate_compose_fyle()