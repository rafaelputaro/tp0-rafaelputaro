## Parte 1: Introducción a Docker
En esta primera parte del trabajo práctico se plantean una serie de ejercicios que sirven para introducir las herramientas básicas de Docker que se utilizarán a lo largo de la materia. El entendimiento de las mismas será crucial para el desarrollo de los próximos TPs.

### Ejercicio N°1:
Además, definir un script de bash `generar-compose.sh` que permita crear una definición de DockerCompose con una cantidad configurable de clientes.  El nombre de los containers deberá seguir el formato propuesto: client1, client2, client3, etc. 

El script deberá ubicarse en la raíz del proyecto y recibirá por parámetro el nombre del archivo de salida y la cantidad de clientes esperados:

`./generar-compose.sh docker-compose-dev.yaml 5`

Considerar que en el contenido del script pueden invocar un subscript de Go o Python:

```
#!/bin/bash
echo "Nombre del archivo de salida: $1"
echo "Cantidad de clientes: $2"
python3 mi-generador.py $1 $2
```

## Resolución:

Se creo un script bash que llama a un script de python que crea el archivo con los clientes solicitados.

Recomiendo el siguiente uso del programa:

1) Crear archivo docker-compose:
```
. generar-compose.sh docker-compose-dev.yaml <nro clientes>
```
2) Iniciar los clientes y el servidor con docker-compose:
```
make docker-compose-up
```
3) Comprobar que se creo el contenedor:
```
docker ps
```
4) Observar el log:
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
Sin logs
```
REPO_PATH=/home/putaro/Workspace/tp0 pytest -s
```
NOTA: Ha pasado el test.
