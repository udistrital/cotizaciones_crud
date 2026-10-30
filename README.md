# cotizaciones_crud

El API provee la gestión de las solicitudes de cotización del sistema Ágora en su etapa de solicitud y publicación, con sus necesidades, ítems, cargues masivos, forma de pago, segmentación por actividad económica, documentos, requisitos documentales y proveedores convocados.

## Especificaciones Técnicas

### Tecnologías y Versiones

- [Golang](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/golang.md)
- [Beego](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/beego.md)
- [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md)
- [UtilsOAS](https://github.com/udistrital/lineamientos_oas/blob/master/generacion_de_apis/control_error_json_crud.md)

### Variables de Entorno

Configurar antes de ejecutar (ver también `.env.example` y `conf/app.conf`).

```shell
COTIZACIONES_CRUD_API_NAME=[nombre de la API]
COTIZACIONES_CRUD_SERVICE_NAME=[nombre del servicio]
COTIZACIONES_CRUD_API_BASE_DIR=[directorio base de la API con respecto a $GOPATH/src]
COTIZACIONES_CRUD_HTTP_PORT=[puerto de la API]
COTIZACIONES_CRUD_RUN_MODE=[dev|prod]
COTIZACIONES_CRUD_PGUSER=[usuario BD]
COTIZACIONES_CRUD_PGPASS=[password BD]
COTIZACIONES_CRUD_PGHOST=[host BD]
COTIZACIONES_CRUD_PGPORT=[puerto BD]
COTIZACIONES_CRUD_PGDB=[nombre BD]
COTIZACIONES_CRUD_PGSCHEMA=[esquema BD]
```

### Ejecución del Proyecto

#### Clonar el repositorio

```shell
mkdir -p $GOPATH/src/github.com/udistrital
cd $GOPATH/src/github.com/udistrital
git clone https://github.com/udistrital/cotizaciones_crud.git
cd cotizaciones_crud
```

#### Ejecución para Desarrollo

Requiere una base de datos [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md) con el esquema y las tablas requeridas. Ejecute `database/agora_cotizaciones.sql` para prepararla y luego:

```shell
cp .env.example .env # Configurar las variables de entorno
set -a && source .env && set +a

go mod tidy
bee run
```

#### Ejecución con Docker

Compile el proyecto y la imagen:

```shell
go mod tidy
go build -o main
docker build -t cotizaciones_crud .
```

Para ejecutar el servicio en un contenedor. Requiere una base de datos [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md) con el esquema y las tablas requeridas. Ejecute `database/agora_cotizaciones.sql` para prepararla y luego:

```shell
cp .env.example .env # Configurar las variables de entorno
set -a && source .env && set +a
docker run --name cotizaciones_crud -p "$COTIZACIONES_CRUD_HTTP_PORT:$COTIZACIONES_CRUD_HTTP_PORT" --env-file .env cotizaciones_crud
```

### Ejecución con Docker Compose

Para consumo local durante el desarrollo de otros servicios, ejecute Docker Compose después de [compilar el proyecto y la imagen](#ejecución-con-docker). Incluye la base de datos y ejecuta `database/agora_cotizaciones.sql` al inicializarla:

```shell
# Creando red bridge para consumo local
docker network create \
  --driver bridge \
  --subnet 172.28.0.0/16 \
  --gateway 172.28.0.1 \
  back_end

cp .env.example .env # Configurar las variables de entorno
docker compose --env-file .env up
```

El script se ejecuta únicamente cuando se crea el volumen de PostgreSQL. Para repetir la inicialización, elimine el volumen `postgres_data` y levante Compose nuevamente.

### Ejecución Pruebas

Pruebas unitarias

Las pruebas de `tests/models` y `tests/controllers` usan una base de datos SQLite temporal que se crea y elimina en cada ejecución, por lo que no requieren PostgreSQL ni variables de entorno.

```shell
cd tests
go test ./... -v

# Cobertura de models y controllers
go test ./... -coverpkg=github.com/udistrital/cotizaciones_crud/models,github.com/udistrital/cotizaciones_crud/controllers -cover

# Limpiar la caché de las pruebas
go clean -testcache
```

Pruebas de integración

SQLite no aplica los índices únicos ni las llaves foráneas de `database/agora_cotizaciones.sql`. Con `COTIZACIONES_CRUD_PRUEBAS_POSTGRES=true`, las pruebas crean en el servidor PostgreSQL configurado una base de datos temporal con ese script y la eliminan al terminar. Así se ejecutan contra PostgreSQL las pruebas de `tests/models` y `tests/controllers` y además las de `tests/integracion`, que validan los índices únicos, las llaves foráneas, la longitud y el rango de los datos y los mensajes de error de la base de datos. Sin la variable, las pruebas de `tests/integracion` se omiten.

La conexión usa las mismas variables `COTIZACIONES_CRUD_PG*` del API; el usuario debe tener permiso para crear bases de datos. La base de datos del API no se modifica.

```shell
export COTIZACIONES_CRUD_PRUEBAS_POSTGRES=true
cd tests
go test ./... -v
```

## Estado CI

| Develop | Release 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) |

## Modelo de Datos

[Modelo de Datos API CRUD Cotizaciones](https://github.com/udistrital/cotizaciones_crud/blob/develop/database/agora_cotizaciones.sql)

## Licencia

This file is part of cotizaciones_crud.

cotizaciones_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cotizaciones_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with cotizaciones_crud. If not, see [https://www.gnu.org/licenses/](https://www.gnu.org/licenses/).
