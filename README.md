### Ejercicio N°3:

Crear un script de bash `validar-echo-server.sh` que permita verificar el correcto funcionamiento del servidor utilizando el comando `netcat` para interactuar con el mismo. Dado que el servidor es un EchoServer, se debe enviar un mensaje al servidor y esperar recibir el mismo mensaje enviado.

En caso de que la validación sea exitosa imprimir: `action: test_echo_server | result: success`, de lo contrario imprimir:`action: test_echo_server | result: fail`.

El script deberá ubicarse en la raíz del proyecto. Netcat no debe ser instalado en la máquina _host_ y no se puede exponer puertos del servidor para realizar la comunicación (hint: `docker network`). `

### Resolución:



### Material teórico para el alumno: 

* Instalar netcat (GNU) en Manjaro:
```
sudo pacman -S netcat
```
* Tutorial Netcat: https://www.youtube.com/watch?v=utdXuu4fNQE

* Obtener información de red en Manjaro:
```
ip addr
```



https://lathack.com/senales-y-prioridades-de-procesos-en-linux/