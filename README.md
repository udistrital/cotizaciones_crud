# cotizaciones_crud

Servicio CRUD para la gestion de cotizaciones del sistema Agora (etapa de solicitud y publicacion), construido con Beego v2 y PostgreSQL.

## Requisitos

- Go 1.27+
- Bee v2 (`go install github.com/beego/bee/v2@latest`)
- Docker y Docker Compose (opcional para ejecucion en contenedores)
- Red Docker `back_end` creada si usa `docker-compose`

## Variables de entorno

El proyecto usa variables con prefijo `COTIZACIONES_CRUD_` que se consumen desde `conf/app.conf`.

### Aplicacion

- `COTIZACIONES_CRUD_APP_NAME`: nombre de la aplicacion.
- `COTIZACIONES_CRUD_HTTP_PORT`: puerto HTTP del servicio.
- `COTIZACIONES_CRUD_RUN_MODE`: modo de ejecucion (`dev`, `prod`, etc.).
- `COTIZACIONES_CRUD_SERVICE_NAME`: nombre del contenedor del API (Docker).
- `COTIZACIONES_CRUD_API_BASE_DIR`: ruta base dentro del contenedor.
- `COTIZACIONES_CRUD_API_NAME`: nombre del directorio del proyecto.

### Base de datos

- `COTIZACIONES_CRUD_PGUSER`: usuario PostgreSQL.
- `COTIZACIONES_CRUD_PGPASS`: clave PostgreSQL.
- `COTIZACIONES_CRUD_PGHOST`: host PostgreSQL.
- `COTIZACIONES_CRUD_PGPORT`: puerto PostgreSQL.
- `COTIZACIONES_CRUD_PGDB`: nombre de la base de datos.
- `COTIZACIONES_CRUD_PGSCHEMA`: esquema PostgreSQL (`cotizaciones`).

## Ejecucion local

1. Prepare la base de datos ejecutando `database/agora_cotizaciones.sql` (crea el esquema `cotizaciones`).
2. Configure las variables de entorno:

	```bash
	cp .env.example .env # ajustar los valores
	set -a && source .env && set +a
	```

3. Descargue dependencias:

	```bash
	go mod download
	```

4. Ejecute con generacion de documentacion:

	```bash
	bee run -downdoc=true -gendoc=true
	```

5. Pruebe endpoints en:
- API: `http://localhost:${COTIZACIONES_CRUD_HTTP_PORT}/v1/...`
- Swagger UI: `http://localhost:${COTIZACIONES_CRUD_HTTP_PORT}/swagger/`

## Ejecucion con Docker Compose

1. Cree la red (si no existe):

	```bash
	docker network create back_end
	```

2. Levante servicios:

	```bash
	docker-compose up --build
	```

El contenedor de PostgreSQL carga automaticamente el script `database/agora_cotizaciones.sql` al iniciar por primera vez.

## Modelo de datos

El script de esquema se encuentra en `database/agora_cotizaciones.sql`. Corresponde al modelo del Sprint 5 propuesto en [udistrital/agora_documentacion#32](https://github.com/udistrital/agora_documentacion/issues/32) e incluye llaves foraneas, restricciones de unicidad y el diccionario de datos (`COMMENT ON`).

Los catalogos (Parametros), personas y proveedores (Terceros), dependencias (Oikos), unidad ejecutora (Kronos), ordenador del gasto y archivos (gestor documental) se guardan solo como referencia a su `id`; no tienen llave foranea en la base de datos.

Tablas principales:

- `solicitud_cotizacion`
- `cargue_masivo`
- `cargue_masivo_error`
- `documento_solicitud`
- `solicitud_ciiu`
- `solicitud_item`
- `solicitud_necesidad`
- `solicitud_pago_parcial`
- `solicitud_proveedor`
- `solicitud_requisito_documental`

Relaciones principales:

- `solicitud_cotizacion` es entidad raiz.
- `cargue_masivo`, `documento_solicitud`, `solicitud_ciiu`, `solicitud_item`, `solicitud_necesidad`, `solicitud_pago_parcial`, `solicitud_proveedor`, `solicitud_requisito_documental` referencian `solicitud_cotizacion`.
- `cargue_masivo_error` referencia `cargue_masivo`.
- `solicitud_item` tambien referencia `cargue_masivo`.

Todos los recursos manejan auditoria basica (`activo`, `fecha_creacion`, `fecha_modificacion`) y borrado logico (`activo = false`).

## Estado CI

| Develop | Release 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cotizaciones_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/cotizaciones_crud/) |

## Licencia

This file is part of cotizaciones_crud.

cotizaciones_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cotizaciones_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with cotizaciones_crud. If not, see [https://www.gnu.org/licenses/](https://www.gnu.org/licenses/).
