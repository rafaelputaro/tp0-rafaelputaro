
## Parte 2: Repaso de Comunicaciones

Las secciones de repaso del trabajo práctico plantean un caso de uso denominado **Lotería Nacional**. Para la resolución de las mismas deberá utilizarse como base al código fuente provisto en la primera parte, con las modificaciones agregadas en el ejercicio 4.

## Ejercicio N°6:
Modificar los clientes para que envíen varias apuestas a la vez (modalidad conocida como procesamiento por chunks o batchs). La información de cada agencia será simulada por la ingesta de su archivo numerado correspondiente, provisto por la cátedra dentro de .data/datasets.zip. Los batchs permiten que el cliente registre varias apuestas en una misma consulta, acortando tiempos de transmisión y procesamiento.

En el servidor, si todas las apuestas del batch fueron procesadas correctamente, imprimir por log: action: apuesta_recibida | result: success | cantidad: ${CANTIDAD_DE_APUESTAS}. En caso de detectar un error con alguna de las apuestas, debe responder con un código de error a elección e imprimir: action: apuesta_recibida | result: fail | cantidad: ${CANTIDAD_DE_APUESTAS}.
 
La cantidad máxima de apuestas dentro de cada batch debe ser configurable desde config.yaml. Respetar la clave batch: maxAmount, pero modificar el valor por defecto de modo tal que los paquetes no excedan los 8kB.

El servidor, por otro lado, deberá responder con éxito solamente si todas las apuestas del batch fueron procesadas correctamente.

### Resolución:

* Cliente: En el docker-compose-dev.yml agrego los clientes con sus respectivas variables de entorno (con valores inventados por el alumno) como servicios independientes. Luego en la carpeta common creo una serie de entidades que modelan la agencia de lotería, la apuesta con los datos del apostador, un módulo destinado al parseo y otro a la aplicación del protocolo. Luego modifico el módulo cliente para que utilizando estas entidades envíe las apuestas al servidor y reciba la respuesta de este último.

* Servidor: En la carpeta common creo una entidad que representa la lotería nacional la cual se encarga de respaldar las apuestas con los métodos provistos en utils. Además creo una entidad encargada de aplicar el protocolo de comunicación utilizado por el cliente la cual es utiliza por el server.

* Protocolo: <longitud en bytes mensaje><id agencia + apuesta como un string separado por comas>

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
