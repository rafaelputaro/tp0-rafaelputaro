### Ejercicio N°2:
Modificar el cliente y el servidor para lograr que realizar cambios en el archivo de configuración no requiera un nuevo build de las imágenes de Docker para que los mismos sean efectivos. La configuración a través del archivo correspondiente (`config.ini` y `config.yaml`, dependiendo de la aplicación) debe ser inyectada en el container y persistida afuera de la imagen (hint: `docker volumes`).

## Resolución:

* Arme una carpeta config_files donde he creado un archivo de configuración para el servidor y otro para el cliente con datos diferentes a los originales del enunciado. Básicamente cambie los puertos y los pase a modo DEBUG, para que se note al lanzar los contenedores que se están usando otros datos diferentes a las imágenes.

* Modifique el docker-compose para que cree volumes para cada app sobre sendos archivos de configuración respectivamente pero como read only.

## Instrucciones:

1) Crear archivo docker-compose:
```
. generar-compose.sh docker-compose-dev.yaml <nro clientes>
```
2) Iniciar los clientes y el servidor con docker-compose. Aquí se puede ver que lo único que hace es copiar los archivos de configuración dejando intactas las capas anteriores:
```
make docker-compose-up
```
3) Comprobar que se crearon los contenedores:
```
docker ps
```
4) Observar el log. En este punto verificar que se están utilizando los puertos 12344 y modo DEBUG en vez 12345 y modo INFO:
```
make docker-compose-logs
```
5) Finalizar los contenedores en otra terminal:
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
NOTA: