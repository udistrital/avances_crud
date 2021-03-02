# avances_crud

El API avances_crud, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos relacional PostgreSQl (registra avances, tipos de avances, solicitudes de avances, etc).
Esta API representa la capa de datos del subsistema de avances, el cual, a su vez, hace parte de el sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Postgres](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
AVANCES_CRUD_HTTP_PORT=[Puerto asignado para la ejecución del API]
AVANCES_CRUD_RUN_MODE=[Modo de ejecución del api]
AVANCES_CRUD_PGUSER=[Usuario de la base de datos]
AVANCES_CRUD_PGPASS=[Clave del usuario para la conexión a la base de datos]
AVANCES_CRUD_PGHOST=[Host de conexión a la base de datos]
AVANCES_CRUD_PGPORT=[Puerto de base de datos]
AVANCES_CRUD_PGDB=[Nombre de la base de datos]
AVANCES_CRUD_PGSCHEMA=[Esquema de la base de datos]
# Ejemplo
AVANCES_CRUD_HTTP_PORT=8080
AVANCES_CRUD_RUN_MODE=dev
AVANCES_CRUD_PGUSER=postgres
AVANCES_CRUD_PGPASS=***
AVANCES_CRUD_PGHOST=127.0.0.1
AVANCES_CRUD_PGPORT=5432
AVANCES_CRUD_PGDB=avances_db
AVANCES_CRUD_PGSCHEMA=avances
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con AVANCES_CRUD_...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/avances_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/avances_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
AVANCES_CRUD_HTTP_PORT=8080 AVANCES_CRUD_RUN_MODE=dev AVANCES_CRUD_PGUSER=postgres AVANCES_CRUD_PGPASS=**** AVANCES_CRUD_PGHOST=127.0.0.1 AVANCES_CRUD_PGPORT=5432 ...

# 5. Ejecutar bee run
bee run
```

### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
# docker build --tag=avances_crud . --no-cache
# docker run -p 80:80 avances_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/avances_crud

#2. Moverse a la carpeta del repositorio
cd avances_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

## Modelo de datos
[Modelo de datos relacional avances_crud](database/modelobd.png)


## Estado CI
| Develop | Relese  | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/avances_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/avances_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/avances_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/avances_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/avances_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/avances_crud) |


## Licencia
This file is part of avances_crud

avances_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

avances_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with avances_crud. If not, see https://www.gnu.org/licenses/.
