package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// FormatoFecha es el formato de las columnas date en el JSON: solo el día, sin hora.
const FormatoFecha = "2006-01-02"

// Fecha representa un día del calendario (columna date), sin hora ni zona horaria.
// En el JSON se escribe y se lee como "2006-01-02"; también acepta una fecha y hora en
// formato RFC 3339, de la que toma el día tal como está escrito.
//
// Se guarda y se lee como día, sin pasar por conversiones de zona horaria: con time.Time,
// la ORM convertía la medianoche a la zona de la base y el día se corría al anterior.
type Fecha struct {
	time.Time
}

// NuevaFecha retorna la fecha del día indicado.
func NuevaFecha(anio int, mes time.Month, dia int) Fecha {
	return Fecha{Time: time.Date(anio, mes, dia, 0, 0, 0, 0, time.UTC)}
}

// String retorna la fecha en formato 2006-01-02.
func (f Fecha) String() string {
	return f.Format(FormatoFecha)
}

// MarshalJSON escribe la fecha como "2006-01-02".
func (f Fecha) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

// UnmarshalJSON lee "2006-01-02" o una fecha y hora RFC 3339 (toma el día escrito, sin
// convertir zonas). null deja la fecha sin asignar.
func (f *Fecha) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var texto string
	if err := json.Unmarshal(b, &texto); err != nil {
		return fmt.Errorf("fecha inválida %s: se espera el formato %s", b, FormatoFecha)
	}
	t, err := time.Parse(FormatoFecha, texto)
	if err != nil {
		t, err = time.Parse(time.RFC3339Nano, texto)
	}
	if err != nil {
		return fmt.Errorf("fecha inválida %q: se espera el formato %s", texto, FormatoFecha)
	}
	*f = NuevaFecha(t.Year(), t.Month(), t.Day())
	return nil
}

// FieldType indica a la ORM que la columna es de tipo date.
func (f *Fecha) FieldType() int {
	return orm.TypeDateField
}

// SetRaw asigna el valor leído de la base de datos. El driver lib/pq entrega la columna
// date como la medianoche de ese día marcada como UTC, y la ORM la pasa a la zona de la
// base sin cambiar el instante; en UTC se recupera el día guardado.
func (f *Fecha) SetRaw(valor interface{}) error {
	switch v := valor.(type) {
	case nil:
		*f = Fecha{}
	case time.Time:
		t := v.UTC()
		*f = NuevaFecha(t.Year(), t.Month(), t.Day())
	case string, []byte:
		texto := strings.TrimSpace(fmt.Sprintf("%s", v))
		if len(texto) > len(FormatoFecha) {
			texto = texto[:len(FormatoFecha)]
		}
		t, err := time.Parse(FormatoFecha, texto)
		if err != nil {
			return err
		}
		*f = NuevaFecha(t.Year(), t.Month(), t.Day())
	default:
		return fmt.Errorf("no se puede convertir %v (%T) en Fecha", valor, valor)
	}
	return nil
}

// RawValue retorna el valor que se guarda en la base de datos: el día como texto, sin
// hora ni zona, para que no se corra al convertir zonas horarias.
func (f *Fecha) RawValue() interface{} {
	return f.String()
}
