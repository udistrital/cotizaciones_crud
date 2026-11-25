// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// Sentencia es una sentencia SQL recibida por el controlador falso, con sus argumentos.
type Sentencia struct {
	SQL  string
	Args []driver.Value
}

// CampoOpcional describe un campo escalar opcional y el valor esperado. Valor nil
// indica que se espera NULL en la base de datos y null en la respuesta.
type CampoOpcional struct {
	Campo   string      // nombre del campo en el modelo y en el JSON
	Columna string      // nombre de la columna en la base de datos
	Valor   interface{} // valor esperado, o nil
}

const nombreFalso = "cotizaciones_falso"

var (
	ultima          Sentencia
	filasSiguientes []map[string]interface{}
	falsoRegistrado bool
)

// DefinirFila fija la única fila que devuelve el controlador falso, por nombre de
// columna. nil es NULL. Las columnas no indicadas devuelven 1.
func DefinirFila(valores map[string]interface{}) {
	filasSiguientes = []map[string]interface{}{valores}
}

// DefinirFilas fija las filas que devuelve el controlador falso, en el orden indicado.
// Sin argumentos, la consulta no devuelve filas. El COUNT(*) devuelve la cantidad de
// filas definidas, como lo haría la base de datos sin LIMIT ni OFFSET.
func DefinirFilas(filas ...map[string]interface{}) {
	filasSiguientes = append([]map[string]interface{}{}, filas...)
}

// LimpiarSentencia borra la última sentencia SQL registrada, para que una prueba pueda
// comprobar que una petición no ejecutó ninguna consulta.
func LimpiarSentencia() {
	ultima = Sentencia{}
}

// filasDefinidas retorna las filas que entrega la consulta. Sin definición, la
// consulta devuelve una fila con los valores por omisión.
func filasDefinidas() []map[string]interface{} {
	if filasSiguientes == nil {
		return []map[string]interface{}{nil}
	}
	return filasSiguientes
}

// ConectarFalso registra el alias default del ORM con una base de datos que no se
// conecta a ningún servidor. Debe llamarse una vez por proceso, desde TestMain.
func ConectarFalso() error {
	if !falsoRegistrado {
		sql.Register(nombreFalso, falsoDriver{})
		falsoRegistrado = true
	}
	db, err := sql.Open(nombreFalso, "")
	if err != nil {
		return err
	}
	return orm.AddAliasWthDB("default", "postgres", db)
}

// UltimaSentencia retorna la última sentencia SQL que recibió el controlador falso.
func UltimaSentencia() Sentencia {
	return ultima
}

// ValorColumna retorna el valor que la sentencia envía para la columna indicada, en un
// INSERT o en un UPDATE. El segundo valor es false si la sentencia no la incluye.
func ValorColumna(s Sentencia, columna string) (driver.Value, bool) {
	if strings.HasPrefix(s.SQL, "UPDATE ") {
		marca := "\"" + columna + "\" = $"
		i := strings.Index(s.SQL, marca)
		if i < 0 {
			return nil, false
		}
		n := 0
		for _, c := range s.SQL[i+len(marca):] {
			if c < '0' || c > '9' {
				break
			}
			n = n*10 + int(c-'0')
		}
		if n < 1 || n > len(s.Args) {
			return nil, false
		}
		return s.Args[n-1], true
	}
	ini := strings.Index(s.SQL, "(")
	fin := strings.Index(s.SQL, ") VALUES")
	if ini < 0 || fin < ini {
		return nil, false
	}
	for i, c := range strings.Split(s.SQL[ini+1:fin], ",") {
		if strings.Trim(strings.TrimSpace(c), "\"") == columna && i < len(s.Args) {
			return s.Args[i], true
		}
	}
	return nil, false
}

// Desreferenciar retorna el valor apuntado por p, o nil si p es nil.
func Desreferenciar[T any](p *T) interface{} {
	if p == nil {
		return nil
	}
	return *p
}

// VerificarColumnaEnviada comprueba el valor que la última sentencia (INSERT o UPDATE)
// envía para la columna. Un esperado nil exige NULL.
func VerificarColumnaEnviada(t *testing.T, columna string, esperado interface{}) {
	t.Helper()
	s := UltimaSentencia()
	valor, ok := ValorColumna(s, columna)
	if !ok {
		t.Fatalf("La sentencia no incluye la columna %s: %s", columna, s.SQL)
	}
	if esperado == nil {
		if valor != nil {
			t.Errorf("Se esperaba NULL en %s y se envió %v", columna, valor)
		}
		return
	}
	if valor == nil {
		t.Errorf("Se esperaba %v en %s y se envió NULL", esperado, columna)
		return
	}
	if fmt.Sprint(valor) != fmt.Sprint(esperado) {
		t.Errorf("Se esperaba %v en %s y se envió %v", esperado, columna, valor)
	}
}

// VerificarValorLeido compara el valor leído de un campo con el esperado. Un esperado
// nil exige nil. Las fechas se comparan como instantes: beego las entrega en la zona
// horaria de la sesión.
func VerificarValorLeido(t *testing.T, campo string, esperado, obtenido interface{}) {
	t.Helper()
	if e, ok := esperado.(time.Time); ok {
		if o, ok := obtenido.(time.Time); ok && e.Equal(o) {
			return
		}
	}
	if !reflect.DeepEqual(esperado, obtenido) {
		t.Errorf("Campo %s: se esperaba %v y se obtuvo %v", campo, esperado, obtenido)
	}
}

// verificarRespuesta compara el campo opcional de la respuesta JSON. Los valores
// esperados se normalizan con JSON para que coincidan con lo que decodifica el test.
// Las fechas se comparan como instantes.
func verificarRespuesta(t *testing.T, c CampoOpcional, obtenido interface{}) {
	t.Helper()
	if c.Valor == nil {
		if obtenido != nil {
			t.Errorf("Campo %s en la respuesta: se esperaba null y se obtuvo %v", c.Campo, obtenido)
		}
		return
	}
	if e, ok := c.Valor.(time.Time); ok {
		cadena, _ := obtenido.(string)
		if o, err := time.Parse(time.RFC3339Nano, cadena); err == nil && e.Equal(o) {
			return
		}
		t.Errorf("Campo %s en la respuesta: se esperaba %v y se obtuvo %v", c.Campo, e, obtenido)
		return
	}
	b, err := json.Marshal(c.Valor)
	if err != nil {
		t.Fatal("No se pudo serializar el valor esperado:", err)
	}
	var esperado interface{}
	if err := json.Unmarshal(b, &esperado); err != nil {
		t.Fatal("No se pudo decodificar el valor esperado:", err)
	}
	VerificarValorLeido(t, c.Campo, esperado, obtenido)
}

// ProbarPostOpcionales envía el cuerpo por POST y verifica la respuesta 201. Comprueba
// el valor que llega a la base de datos y el que devuelve la respuesta.
func ProbarPostOpcionales(t *testing.T, ruta string, cuerpo map[string]interface{}, campos []CampoOpcional) {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo)), http.StatusCreated)
	for _, c := range campos {
		VerificarColumnaEnviada(t, c.Columna, c.Valor)
	}
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// ProbarGetOpcionales define en la base de datos falsa una fila con el valor de cada
// campo opcional, consulta el registro 1 y verifica que la respuesta lo devuelva igual.
func ProbarGetOpcionales(t *testing.T, ruta string, campos []CampoOpcional) {
	t.Helper()
	fila := map[string]interface{}{}
	for _, c := range campos {
		fila[c.Columna] = c.Valor
	}
	DefinirFila(fila)
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/1", ""), http.StatusOK)
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// ProbarPutOpcionales envía el cuerpo por PUT al registro 1 y verifica la respuesta 200.
// Comprueba el valor que llega en el UPDATE y el que devuelve la respuesta.
func ProbarPutOpcionales(t *testing.T, ruta string, cuerpo map[string]interface{}, campos []CampoOpcional) {
	t.Helper()
	DefinirFila(nil)
	cuerpo["Id"] = 1
	r := VerificarExito(t, Peticion(http.MethodPut, ruta+"/1", JSON(t, cuerpo)), http.StatusOK)
	for _, c := range campos {
		VerificarColumnaEnviada(t, c.Columna, c.Valor)
	}
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// falsoDriver es el driver de database/sql del controlador falso.
type falsoDriver struct{}

func (falsoDriver) Open(string) (driver.Conn, error) { return falsoConexion{}, nil }

// falsoConexion registra cada sentencia y no admite preparación ni transacciones.
type falsoConexion struct{}

func (falsoConexion) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("el controlador falso no prepara sentencias")
}
func (falsoConexion) Close() error { return nil }
func (falsoConexion) Begin() (driver.Tx, error) {
	return nil, errors.New("el controlador falso no maneja transacciones")
}

func (falsoConexion) ExecContext(_ context.Context, q string, a []driver.NamedValue) (driver.Result, error) {
	registrar(q, a)
	return falsoResultado{}, nil
}

func (falsoConexion) QueryContext(_ context.Context, q string, a []driver.NamedValue) (driver.Rows, error) {
	registrar(q, a)
	if strings.HasPrefix(q, "SELECT COUNT(*)") {
		return &falsoFilas{nombres: []string{"count"}, filas: []map[string]interface{}{{"count": int64(len(filasDefinidas()))}}}, nil
	}
	nombres := []string{"id"}
	if i := strings.Index(q, " FROM "); strings.HasPrefix(q, "SELECT ") && i > 0 {
		nombres = nil
		for _, p := range strings.Split(q[len("SELECT "):i], ", ") {
			partes := strings.Split(p, ".")
			nombres = append(nombres, strings.Trim(partes[len(partes)-1], "\""))
		}
	}
	return &falsoFilas{nombres: nombres, filas: filasDefinidas()}, nil
}

func registrar(q string, a []driver.NamedValue) {
	args := make([]driver.Value, len(a))
	for i, v := range a {
		args[i] = v.Value
	}
	ultima = Sentencia{SQL: q, Args: args}
}

type falsoResultado struct{}

func (falsoResultado) LastInsertId() (int64, error) { return 1, nil }
func (falsoResultado) RowsAffected() (int64, error) { return 1, nil }

// falsoFilas entrega las filas indicadas en orden. En cada fila, las columnas que no
// están en el mapa devuelven 1; las que sí, el valor indicado (nil es NULL).
type falsoFilas struct {
	nombres []string
	filas   []map[string]interface{}
	i       int
}

func (f *falsoFilas) Columns() []string { return f.nombres }
func (*falsoFilas) Close() error        { return nil }
func (f *falsoFilas) Next(d []driver.Value) error {
	if f.i >= len(f.filas) {
		return io.EOF
	}
	fila := f.filas[f.i]
	f.i++
	for i, n := range f.nombres {
		d[i] = int64(1)
		if v, ok := fila[n]; ok {
			d[i] = v
		}
	}
	return nil
}
