# corregir_sort

Corrige el ordenamiento (`sortby` y `order`) del servicio `GetAll` de los CRUD generados
con beego. Está pensado para cualquier repositorio de la organización que siga la forma de
bee: modelos con una función `GetAll<Modelo>` y controladores con un método `GetAll`.

## Qué corrige

En cada `GetAll<Modelo>` de `models/*.go`:

- Un `sortby` sin `order` se ordena ascendente.
- Cada campo de `sortby` se valida contra el modelo antes de consultar la base de datos.
  Un campo que no existe, un `order` distinto de `asc` o `desc`, o cantidades incompatibles
  devuelven un error que envuelve `models.ErrSortInvalido`.

En cada método `GetAll` de `controllers/*.go`:

- Un error que envuelve `models.ErrSortInvalido` responde **400**. Los demás errores conservan
  su respuesta **404**.

Los campos se aceptan por nombre de campo del struct (`FechaCreacion`) o por nombre de columna
(`fecha_creacion`). Los campos de relación se recorren con `__` (`MstCotizacionId__Id`).

## Uso

Desde la raíz del proyecto beego:

```shell
go run ./scripts/corregir_sort
```

Para otra ruta raíz:

```shell
go run ./scripts/corregir_sort -raiz ruta/al/proyecto
```

El script crea `models/orden.go` con `ErrSortInvalido` y `construirOrden` si el archivo no
existe. Si ya existe, no lo sobrescribe.

Después de ejecutarlo, regenere las pruebas para que incluyan los casos de ordenamiento:

```shell
go run ./scripts/generar_pruebas
```

## Comportamiento

- **Idempotente.** Ejecutarlo sobre código ya corregido no cambia ningún archivo.
- **Sin nombres propios.** No depende del módulo ni de las entidades del repositorio donde
  se ejecuta. Las entidades se descubren desde el código.
- **Conservador.** Si un `GetAll` o un controlador no tiene la forma esperada, el script
  emite una advertencia con el archivo y el método, y no modifica ese código. Un archivo que
  no se puede analizar también se omite con advertencia.
- **Imports.** Agrega `reflect` y `errors` cuando hacen falta, y quita los que la corrección
  deja sin uso. Formatea el resultado con `go/format`.

Las advertencias van a la salida de error con el prefijo `ADVERTENCIA:`. Un error de
lectura o escritura detiene la ejecución con código de salida 1.

## Requisitos del código

El script reconoce el patrón generado por bee:

- En el modelo, un bloque que empieza con `var sortFields []string` y termina con un `if`,
  seguido de una consulta `Count()`. El `GetAll<Modelo>` devuelve tres valores.
- En el controlador, un `if err != nil` con un `c.Abort("404")` dentro del cuerpo.
- El método `GetAll` del controlador llama a `models.GetAll<Modelo>`. El nombre del paquete
  se toma de esa llamada.

## Pruebas

`scripts/corregir_sort/main_test.go` aplica el script a una copia del código original de
`notificacion` (`testdata/antes`) y compara el resultado con el código corregido a mano
(`testdata/despues`). Las pruebas también verifican que una segunda ejecución no cambie
nada, que un `GetAll` no reconocido se omita con advertencia, y que un `models/orden.go`
existente no se sobrescriba.

```shell
go test ./scripts/corregir_sort/
```

## Limitaciones

- Solo cubre el patrón generado por bee descrito arriba. Un CRUD escrito a mano con otra
  forma se omite con advertencia y hay que corregirlo a mano.
- Si el paquete ya declara `ErrSortInvalido` en otro archivo, el script no crea
  `orden.go`, avisa y no modifica modelos ni controladores.
