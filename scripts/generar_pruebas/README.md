# generar_pruebas

Generador de pruebas del API. Lee el código fuente del proyecto como texto y produce tanto los _helpers_ compartidos como las pruebas de cada entidad. No usa plantillas externas: las fuentes de los helpers y las plantillas de las pruebas están embebidas en `helpers.go`, `capas.go` y `main.go`.

```shell
# Desde la raíz del proyecto
go run ./scripts/generar_pruebas
```

## Qué lee

| Fuente | Qué extrae |
| -- | -- |
| `models/*.go` | Structs, tipos de campo y etiquetas `orm`: `pk`, `auto`, `rel(fk)`, `null`, `column`, `type(uuid)`. También el `TableName()` de cada modelo. |
| `routers/router.go` | La ruta de cada controlador (`NSNamespace` + `NSInclude`), prefijada con `/v1`. |
| `controllers/*.go` | El archivo de cada controlador, para nombrar su `_test.go`. |
| `database/agora_cotizaciones.sql` | Columnas, tipo, longitud, `NOT NULL`/`PRIMARY KEY` y llaves foráneas. |

Al cruzar modelo y SQL, el generador **advierte** (por `stderr`, sin detenerse) cuando:

- una columna del modelo no existe en la tabla,
- una llave foránea declara `null` distinto al `NOT NULL` del SQL,
- un tipo de campo no tiene valor de ejemplo,
- una entidad no tiene un campo de texto o entero para la prueba de PUT.

Si un modelo no tiene controlador ni ruta, se omite.

## Conceptos clave

- **Campo opcional (escalar puntero).** Un campo cuyo tipo en el modelo es puntero (`*string`, `*int`, `*float64`...) y **no** es llave foránea. Representa una columna que admite `NULL`. El generador produce pruebas específicas de opcionalidad para cada uno de estos campos. Los produce el migrador `scripts/migrar_opcionales`.
- **Llave foránea (`rel(fk)`).** Campo que referencia a otro modelo. Es obligatoria cuando no tiene `null`, opcional cuando sí.
- **Campos de servicio.** `Id`, `Activo`, `FechaCreacion`, `FechaModificacion`. Los administra el servicio y no se incluyen en los datos de ejemplo.
- **Relación 1:1.** Cuando la columna `id` de la tabla es a su vez llave foránea a otra tabla, el `Id` no es autoincremental sino el `Id` del padre.
- **Valores de ejemplo.** Según el tipo: UUID fijo para `type(uuid)`, `{"prueba": true}` para columnas `json`/`jsonb`, `"Prueba"` (recortado a la longitud de la columna) para texto, `1` para `int`, `int16(2026)`, una fecha fija para `time.Time`, `1.5` para `float64`, `true` para `bool`. Las llaves foráneas usan un registro recién creado (en integración) o `{Id: 1}` (en unitarias).

## Capas de prueba

El generador emite tres capas, cada una en su carpeta, más una capa de helpers compartidos.

| Capa | Carpeta | Base de datos | Qué prueba |
| -- | -- | -- | -- |
| Integración | `tests/integracion` | PostgreSQL real (variables `COTIZACIONES_CRUD_PG*`) | El ciclo CRUD completo de cada controlador contra una base de datos real. |
| Unitaria de controlador | `tests/unitarias/controladores` | Falsa (en memoria) | El POST/GET/PUT de cada controlador con campos opcionales, sin servidor de base de datos. |
| Unitaria de modelo | `tests/unitarias/modelos` | Falsa (en memoria) | El SQL que `models.Add`/`Update`/`Get` envía y lee para cada campo opcional. |
| Helpers | `tests/helpers` | — | Utilidades compartidas por todas las capas. |

Cada carpeta recibe un `main_gen_test.go` con su `TestMain`:

- **Integración:** llama a `IniciarBaseDatos()`; si las variables no están definidas, las pruebas se omiten (`SKIP`). Configura beego y registra las rutas.
- **Unitaria de controlador:** llama a `ConectarFalso()` (base de datos falsa) y configura beego con las rutas.
- **Unitaria de modelo:** solo `ConectarFalso()`; no necesita beego ni rutas.

> Nota: la capa de integración solo se genera para entidades con controlador y ruta. Las capas unitarias (controlador y modelo) solo se generan para entidades que tienen al menos un campo escalar opcional.

## Helpers

Las fuentes de `tests/helpers` se emiten desde el generador. Unas se copian tal cual (escritas a mano) y otras se construyen con plantilla.

### `basedatos.go` (copiado)

Conexión a la base de datos real y utilidades de lectura directa.

| Helper | Qué hace |
| -- | -- |
| `IdInexistente` | Constante de un Id que no existe en ninguna tabla (`999999999`). |
| `IniciarBaseDatos()` | Registra la base PostgreSQL de las variables `COTIZACIONES_CRUD_PG*`. Si faltan, no conecta y las pruebas se omiten. |
| `BaseDisponible()` | Indica si hay base de datos configurada. |
| `RequerirBaseDatos(t)` | Omite (`t.Skip`) la prueba si no hay base de datos. |
| `EliminarAlFinal(t, tabla, id)` | Programa el `DELETE` físico del registro al terminar la prueba (orden inverso a la creación). |
| `Fila(t, tabla, id)` | Devuelve el registro tal como está en la base, como mapa (vía `row_to_json`). |
| `SinColumnas(fila, cols...)` | Copia de la fila sin las columnas indicadas. |
| `Contar(t, tabla)` | Número de registros de la tabla. |

### `http.go` (copiado)

Envío de peticiones al enrutador de beego (sin levantar servidor) y lectura de la respuesta institucional (`Success`/`Status`/`Message`/`Data`/`Total`).

| Helper | Qué hace |
| -- | -- |
| `Peticion(metodo, ruta, cuerpo)` | Envía una petición HTTP al handler de beego y devuelve el `ResponseRecorder`. |
| `JSON(t, v)` | Serializa `v` a JSON para el cuerpo. |
| `AMapa(t, v)` | Convierte un valor (p. ej. un modelo) al mapa que tendría en el JSON de la respuesta. |
| `Respuesta.Dato(t)` | Campo `Data` como objeto. |
| `Respuesta.Lista(t)` | Campo `Data` como lista de objetos. |
| `Respuesta.IdDato(t)` | `Id` del objeto en `Data`. |
| `Respuesta.TotalLista(t)` | Campo `Total` de un GetAll. |
| `VerificarExito(t, w, estado)` | Comprueba el código HTTP, `Success=true` y `Status` igual al código. |
| `VerificarError(t, w, estado)` | Comprueba el código HTTP, `Success=false` y `Status` igual al código. |
| `VerificarFiltroInvalido(t, ruta)` | GetAll con `query` mal formado: 400, `Data` nulo y mensaje con el formato esperado. |
| `VerificarIgual(t, desc, obtenido, esperado)` | Compara dos valores con `reflect.DeepEqual`. |

### `verificaciones.go` (copiado)

Pruebas CRUD de alto nivel usadas por la capa de integración.

| Helper | Qué hace |
| -- | -- |
| `FechaCliente` | Fecha que el cliente envía en campos administrados y que el servicio debe ignorar. |
| `CrearPorPost(t, ruta, tabla, cuerpo)` | Crea por POST, verifica 201 y programa el borrado. Devuelve el Id. |
| `VerificarCamposAdministrados(t, tabla, id)` | Comprueba que el servicio asigna `Activo=true` y las fechas, ignorando lo enviado. |
| `ProbarPostInvalido(t, ruta)` | POST responde 400 ante JSON mal formado o cuerpo vacío. |
| `ProbarGetOne(t, ruta, id)` | GET por Id, verifica 200 y devuelve el objeto. |
| `ProbarGetOneInexistente(t, ruta)` | GET responde 404 ante Id inexistente o no numérico. |
| `ConsultarGetAll(t, ruta, params)` | Ejecuta GetAll y devuelve la lista. |
| `ConsultarPagina(t, ruta, params)` | Ejecuta GetAll y devuelve lista y `Total`. |
| `ContieneId(lista, id)` | Indica si la lista contiene un objeto con ese Id. |
| `ProbarGetAll(t, ruta, tabla, id)` | Prueba el GetAll completo: filtro por Id, lista vacía, `Total`, `limit`/`offset`, `order` asc/desc, `fields` y 400 ante filtros inválidos. |
| `VerificarFKExpandida(t, dato, campo, relacionado)` | Comprueba que la llave foránea venga expandida con los datos del relacionado (anidados solo por Id). |
| `ProbarPut(t, ruta, tabla, id, campo, valor, columna)` | Prueba el PUT: actualiza el campo, conserva `fecha_creacion`/`usuario_creacion_id`, actualiza `fecha_modificacion`, y 400/404 ante Id distinto, JSON inválido o Id inexistente. |
| `ProbarDelete(t, ruta, tabla, id)` | Prueba el borrado lógico: 200, `activo=false`, `fecha_modificacion` actualizada, demás columnas intactas, registro aún consultable y 404 ante Id inexistente. |

### `falso_gen.go` (generado)

Controlador (driver) falso de `database/sql`: registra cada sentencia SQL con sus argumentos y devuelve una fila cuyo `id` es 1. Permite las pruebas unitarias sin servidor de base de datos.

| Helper | Qué hace |
| -- | -- |
| `Sentencia` | Una sentencia SQL recibida, con su texto y argumentos. |
| `CampoOpcional` | Describe un campo opcional: nombre en el modelo/JSON, nombre de columna y valor esperado (`nil` = NULL). |
| `ConectarFalso()` | Registra el alias `default` del ORM con el driver falso. |
| `DefinirFila(valores)` | Fija la fila que devolverá la siguiente consulta (por columna; `nil` es NULL). |
| `UltimaSentencia()` | Última sentencia SQL recibida. |
| `ValorColumna(s, columna)` | Valor que un INSERT o UPDATE envía para esa columna. |
| `Desreferenciar(p)` | Valor apuntado por `p`, o `nil` si es `nil`. |
| `VerificarColumnaEnviada(t, columna, esperado)` | Comprueba el valor que la última sentencia envía (INSERT/UPDATE); `nil` exige NULL. |
| `VerificarValorLeido(t, campo, esperado, obtenido)` | Compara valor leído con el esperado (fechas como instantes). |
| `ProbarPostOpcionales(t, ruta, cuerpo, campos)` | POST por el falso: verifica 201, el valor que llega a la "base" y el que devuelve la respuesta. |
| `ProbarGetOpcionales(t, ruta, campos)` | Fija una fila con esos valores, hace GET del registro 1 y verifica la respuesta. |
| `ProbarPutOpcionales(t, ruta, cuerpo, campos)` | PUT por el falso al registro 1: verifica 200, el valor del UPDATE y el de la respuesta. |

### `punteros_gen.go` (generado)

| Helper | Qué hace |
| -- | -- |
| `Ptr[T](v)` | Devuelve un puntero al valor. Los campos opcionales son punteros y Go no permite tomar la dirección de un literal; los registros de ejemplo usan `Ptr(valor)`. |

### `registros_gen.go` (generado)

Dos funciones por modelo:

| Helper | Qué hace |
| -- | -- |
| `Nuevo<Modelo>(t)` | Registro de ejemplo válido **sin guardar**; crea en la base sus llaves foráneas. En relación 1:1 toma el `Id` del padre. |
| `Crear<Modelo>(t)` | Guarda el registro con `models.Add<Modelo>` (con `Activo` y fechas) y programa su borrado. |

## Pruebas generadas por capa

### Integración — `tests/integracion/<entidad>_test.go`

Un archivo por controlador. Prueban el ciclo CRUD completo contra PostgreSQL. Todas empiezan con `RequerirBaseDatos(t)`, por lo que se omiten si no hay base configurada.

| Prueba | Qué prueba | Entrada | Salida esperada |
| -- | -- | -- | -- |
| `Test<E>Post` | Creación. Con Id autoincremental: dos POST dan Ids crecientes y se ignora el Id enviado. En relación 1:1: el Id es obligatorio y debe existir. Las llaves foráneas obligatorias no pueden ir nulas. El servicio asigna `Activo=true` y las fechas. JSON inválido → 400. | Cuerpo válido con `Activo=false`, fechas de cliente e Id inexistente. | 201 y registro con campos administrados correctos; 400 ante cuerpos inválidos o FK obligatoria nula. |
| `Test<E>GetOne` | Consulta por Id, con llaves foráneas expandidas. Id inexistente/no numérico → 404. | Id de un registro creado. | 200 con el objeto y sus FK expandidas; 404 para Id inexistente. |
| `Test<E>GetAll` | Filtro por Id y por cada FK, lista vacía sin resultados, `Total`, paginación (`limit`/`offset`), orden asc/desc, `fields` y 400 ante filtros inválidos. | Registro creado y parámetros `query`/`sortby`/`order`/`limit`/`offset`/`fields`. | Listas y `Total` coherentes con cada filtro; FK expandidas; 400 ante filtros mal formados. |
| `Test<E>Put` | Actualiza un campo (primer texto editable; si no hay, primer entero; último recurso `UsuarioModificacionId`). Conserva creación, actualiza modificación, 400/404 ante Id distinto o inexistente. | Registro creado + nuevo valor del campo. | 200 y columna actualizada; `fecha_creacion`/`usuario_creacion_id` intactos; 400/404 en los casos de error. |
| `Test<E>Delete` | Borrado lógico. | Id de un registro creado. | 200, `activo=false`, `fecha_modificacion` actualizada, resto intacto, aún consultable; 404 ante Id inexistente. |
| `Test<E><FK>Opcional` | Por cada FK **opcional**: el registro sin esa FK se puede crear, consultar (GetOne/GetAll), actualizar y borrar. | Cuerpo con la FK en `null`. | El registro se crea y la FK sale `null`; CRUD completo funciona. |
| `Test<E><Campo>OpcionalOmitido` | Por cada escalar opcional: si se **omite** del payload, el servicio lo devuelve `null` y la columna queda `NULL`. | Cuerpo sin ese campo. | GetOne devuelve `null`; la fila en base tiene la columna `NULL`. |
| `Test<E><Campo>OpcionalNulo` | Por cada escalar opcional: si se envía `null`, mismo resultado. | Cuerpo con el campo en `null`. | GetOne devuelve `null`; la fila en base tiene la columna `NULL`. |

### Unitaria de controlador — `tests/unitarias/controladores/<entidad>_test.go`

Un archivo por entidad **con campos opcionales**. Ejercen el POST/GET/PUT del controlador contra la base de datos falsa (sin PostgreSQL). Cada prueba cubre todos los campos opcionales de la entidad a la vez. Se apoyan en `ejemplo<E>()` (entidad con todos sus campos de ejemplo, FK como `{Id: 1}`).

| Prueba | Qué prueba | Entrada | Salida esperada |
| -- | -- | -- | -- |
| `TestPost<E>OpcionalesOmitidos` | POST sin los campos opcionales. | Cuerpo sin esos campos. | El INSERT envía `NULL` y la respuesta los trae `null`. |
| `TestPost<E>OpcionalesNull` | POST con los campos opcionales en `null`. | Cuerpo con los campos en `null`. | Igual: `NULL` en el INSERT y `null` en la respuesta. |
| `TestPost<E>OpcionalesConValor` | POST con los campos opcionales con su valor. | Cuerpo con los valores de ejemplo. | El INSERT envía el valor y la respuesta lo devuelve igual. |
| `TestGet<E>OpcionalesNull` | GET cuando la base falsa devuelve `NULL`. | Fila falsa con las columnas en `NULL`. | La respuesta trae los campos `null`. |
| `TestGet<E>OpcionalesConValor` | GET cuando la base falsa devuelve valores. | Fila falsa con los valores. | La respuesta devuelve esos valores. |
| `TestPut<E>OpcionalesOmitidos` | PUT sin los campos opcionales. | Cuerpo sin esos campos. | El UPDATE envía `NULL` y la respuesta los trae `null`. |
| `TestPut<E>OpcionalesNull` | PUT con los campos en `null`. | Cuerpo con los campos en `null`. | Igual que el anterior. |
| `TestPut<E>OpcionalesConValor` | PUT con los campos con su valor. | Cuerpo con los valores. | El UPDATE envía el valor y la respuesta lo devuelve. |

### Unitaria de modelo — `tests/unitarias/modelos/<entidad>_test.go`

Un archivo por entidad **con campos opcionales**. Verifican directamente el SQL que las funciones del modelo generan, usando la base de datos falsa. Por cada campo opcional se emiten tres funciones, cada una con subpruebas (`t.Run`):

| Prueba (por campo) | Subpruebas | Qué prueba | Salida esperada |
| -- | -- | -- | -- |
| `Test<E><Campo>` (sobre `Add<E>`) | `omitido`, `null`, `valor` | Valor que el **INSERT** envía para la columna según cómo venga el campo. | `omitido`/`null` → la columna va `NULL`; `valor` → va el valor. |
| `Test<E><Campo>Get` (sobre `Get<E>ById`) | `null`, `valor` | Lectura de la columna desde la base falsa. | `NULL` llega como `nil`; un valor llega igual. |
| `Test<E><Campo>Update` (sobre `Update<E>ById`) | `omitido`, `null`, `valor` | Valor que el **UPDATE** envía para la columna. | `omitido`/`null` → la columna va `NULL`; `valor` → va el valor. |

En las tres, el campo bajo prueba se omite, se pone en `nil` o se asigna su valor, manteniendo el resto de los campos con sus ejemplos. La verificación usa `VerificarColumnaEnviada` (para INSERT/UPDATE) o `VerificarValorLeido` (para GET).

## Resumen del flujo

1. `leerSQL`, `leerRutas`, `leerControllers`, `leerModelos` cargan las cuatro fuentes.
2. `completar` cruza cada modelo con su tabla SQL, fija los valores de ejemplo, el campo de PUT y clasifica relaciones (obligatorias/opcionales) y escalares opcionales.
3. `generarHelpers` escribe `tests/helpers` (copiados + generados).
4. Por cada entidad se escribe su prueba de integración en `tests/integracion`.
5. `generarCapas` escribe los `TestMain` de cada carpeta y, para las entidades con campos opcionales, las pruebas unitarias de controlador y de modelo.

Todos los archivos generados llevan la cabecera `NO EDITAR` y se regeneran con `go run ./scripts/generar_pruebas`.
