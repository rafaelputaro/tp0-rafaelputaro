### Ejercicio N°3:

Crear un script de bash `validar-echo-server.sh` que permita verificar el correcto funcionamiento del servidor utilizando el comando `netcat` para interactuar con el mismo. Dado que el servidor es un EchoServer, se debe enviar un mensaje al servidor y esperar recibir el mismo mensaje enviado.

En caso de que la validación sea exitosa imprimir: `action: test_echo_server | result: success`, de lo contrario imprimir:`action: test_echo_server | result: fail`.

El script deberá ubicarse en la raíz del proyecto. Netcat no debe ser instalado en la máquina _host_ y no se puede exponer puertos del servidor para realizar la comunicación (hint: `docker network`). `

### Resolución:

Modifico el docker-compose-dev.yaml para correr el script de validación como un servicio independiente.

### Instrucciones de uso:

1) Crear archivo docker-compose:
```
. generar-compose.sh docker-compose-dev.yaml <nro clientes>
```
2)  Iniciar los clientes y el servidor con docker-compose:
```
make docker-compose-up
```
3) Comprobar que se creo el contenedor:
```
docker ps
```
4) Observar el log donde se podrá verificar el funcionamiento del script de validación:
```
make docker-compose-logs
```
5) Finalizar el contenedor en otra terminal:
```
make docker-compose-down
```
## Tests (para el alumno)

Máquina del trabajo:

Sin logs
```
REPO_PATH=/home/putaro/Workspace/tp0 pytest
```
Con logs:
```
REPO_PATH=/home/putaro/Workspace/tp0 pytest -s
```
NOTA: Paso el test.