
## Parte 2: Repaso de Comunicaciones

Las secciones de repaso del trabajo práctico plantean un caso de uso denominado **Lotería Nacional**. Para la resolución de las mismas deberá utilizarse como base al código fuente provisto en la primera parte, con las modificaciones agregadas en el ejercicio 4.

## Ejercicio N°7:

Modificar los clientes para que notifiquen al servidor al finalizar con el envío de todas las apuestas y así proceder con el sorteo. Inmediatamente después de la notificacion, los clientes consultarán la lista de ganadores del sorteo correspondientes a su agencia. Una vez el cliente obtenga los resultados, deberá imprimir por log: action: consulta_ganadores | result: success | cant_ganadores: ${CANT}.

El servidor deberá esperar la notificación de las 5 agencias para considerar que se realizó el sorteo e imprimir por log: action: sorteo | result: success. Luego de este evento, podrá verificar cada apuesta con las funciones load_bets(...) y has_won(...) y retornar los DNI de los ganadores de la agencia en cuestión. Antes del sorteo, no podrá responder consultas por la lista de ganadores. Las funciones load_bets(...) y has_won(...) son provistas por la cátedra y no podrán ser modificadas por el alumno.

### Resolución: 

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
