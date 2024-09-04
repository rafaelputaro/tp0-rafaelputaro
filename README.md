
## Parte 3: Repaso de Concurrencia

## Ejercicio N°8:

Modificar el servidor para que permita aceptar conexiones y procesar mensajes en paralelo. En este ejercicio es importante considerar los mecanismos de sincronización a utilizar para el correcto funcionamiento de la persistencia.

En caso de que el alumno implemente el servidor Python utilizando multithreading, deberán tenerse en cuenta las limitaciones propias del lenguaje.


### Resolución: TODO ADAPTAR A LOS NUEVOS CAMBIOS

* Cliente: Se extiende del ejercicio anterior agregando la inclusión de los tags del protocolo y el polling para la consulta de ganadores.

* Servidor: En el servidor se realizan unos pocos cambios que implican el manejo de los tags del nuevo protocolo y la respuesta ante el polling. Internamente guarda un conjunto con las agencias que ya terminaron (evento inferido a través de que ya han hecho una consulta de ganadores) para responder con la lista de ganadores cuando ya no se ralicen más apuestas.

* Protocolo envío de datos desde cliente: 

    * <Tipo consulta "bets"><"cantidad de apuestas del batch"><longitud en bytes mensaje><"id agencia + apuesta como un string separado por comas">;<"id agencia 1 + apuesta 1 como un string separado por comas">;..........<.......>

    * <Tipo consulta "asks"><longitud id><"agency id">

* Protocolo respuesta servidor:

    * Respuesta envío de batchs: <"cantidad apuestas recibidas">
    
    * Respuesta al recibir "asks":

        * Faltan agencias por finalizar: Servidor responde <"bets\n">

        * Terminaron todas las agencias: Servidor responde <"winners\n"><"DNI ganador 1 como string\n">,.... 

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
3) Observar el log donde se podrá verificar lo solicitado:
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

Fuentes:

https://wiki.python.org/moin/GlobalInterpreterLock

https://docs.python.org/es/dev/library/threading.html

https://docs.python.org/3/library/multiprocessing.html