### Ejercicio N°4:
Modificar servidor y cliente para que ambos sistemas terminen de forma _graceful_ al recibir la signal SIGTERM. Terminar la aplicación de forma _graceful_ implica que todos los _file descriptors_ (entre los que se encuentran archivos, sockets, threads y procesos) deben cerrarse correctamente antes que el thread de la aplicación principal muera. Loguear mensajes en el cierre de cada recurso (hint: Verificar que hace el flag `-t` utilizado en el comando `docker compose down`).

### Resolución:

* Servidor: En "server/common/server.py" creo una función que inicializa una lista donde guardo los sockets de cada cliente e inicializo el método a utilizar cuando se recibe las señales SIGTERM y SINGINT. El método implicado cierra el socket del servidor y el de cada cliente, logueando estos eventos. Cuando un cliente se conecta se da de alta en la lista de clientes y cuando se cierra el socket de un cliente por un error el mismo es quitado de la lista de clientes.

* Cliente: Se define un canal desde el propio "client/main" con el cual se comunica la señal SIGTERM hacia el código del loop de cliente, finalizando y logueando el evento.

### Instrucciones de uso:

Recomiendo el siguiente uso del programa:

1)  Iniciar los clientes y el servidor con docker-compose:
```
make docker-compose-up
```
2) Comprobar que se creo el contenedor:
```
docker ps
```
3) Observar el log donde se podrá verificar el funcionamiento del script de validación:
```
make docker-compose-logs
```
4) Finalizar el contenedor en otra terminal:
```
make docker-compose-down
```
o bien:
```
docker compose -f docker-compose-dev.yaml down -t <tiempo en segundos para shutdown>
```

### Teoría:

SIGTERM: Señal que se envía el proceso para comunicarle un apagado “amable” (cerrando conexiones, archivos y limpiando sus propios búfer).