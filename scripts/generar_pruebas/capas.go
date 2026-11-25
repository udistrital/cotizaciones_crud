package main

import (
	"path/filepath"
	"text/template"
)

// Carpetas y archivos de soporte de las pruebas por capa. Todo se emite desde aquí.
const (
	dirUnitControladores = "tests/unitarias/controladores"
	dirUnitModelos       = "tests/unitarias/modelos"
)

// pruebasMain indica qué configuración necesita el TestMain de cada carpeta.
type pruebasMain struct {
	Integracion   bool // conecta la base de datos PostgreSQL de las variables COTIZACIONES_CRUD_PG*
	Controladores bool // configura beego y registra las rutas (requiere el router)
}

// generarCapas escribe el controlador falso, los TestMain de cada carpeta y, por
// entidad con campos escalares opcionales, las pruebas unitarias del controlador y
// del modelo.
func generarCapas(lista []*Entidad) error {
	mainIntegracion := pruebasMain{Integracion: true, Controladores: true}
	if err := escribir(filepath.Join(dirPruebas, "main_gen_test.go"), plantillaMainPruebas, mainIntegracion, importesMain(true)...); err != nil {
		return err
	}
	mainControladores := pruebasMain{Controladores: true}
	if err := escribir(filepath.Join(dirUnitControladores, "main_gen_test.go"), plantillaMainPruebas, mainControladores, importesMain(true)...); err != nil {
		return err
	}
	mainModelos := pruebasMain{}
	if err := escribir(filepath.Join(dirUnitModelos, "main_gen_test.go"), plantillaMainPruebas, mainModelos, importesMain(false)...); err != nil {
		return err
	}
	for _, e := range lista {
		if len(e.EscalaresOpc) == 0 {
			continue
		}
		if err := escribir(filepath.Join(dirUnitControladores, e.Archivo+"_test.go"), plantillaUnitControlador, e); err != nil {
			return err
		}
		if err := escribir(filepath.Join(dirUnitModelos, e.Archivo+"_test.go"), plantillaUnitModelo, e); err != nil {
			return err
		}
	}
	return nil
}

// importesMain son los imports del TestMain. Con controladores también se
// registran las rutas y la configuración de beego.
func importesMain(controladores bool) []string {
	imp := []string{`"log"`, `"os"`}
	if controladores {
		imp = append(imp,
			`beego "github.com/beego/beego/v2/server/web"`,
			`"github.com/udistrital/utils_oas/v2/customerror"`,
			`_ "github.com/udistrital/cotizaciones_crud/routers"`,
		)
	}
	return imp
}

// plantillaMainPruebas emite el TestMain de una carpeta de pruebas.
var plantillaMainPruebas = template.Must(template.New("main").Parse(`
// TestMain prepara el entorno de las pruebas de esta carpeta.
func TestMain(m *testing.M) {
{{- if .Integracion}}
	if err := helpers.IniciarBaseDatos(); err != nil {
		log.Fatal("No se pudo conectar a la base de datos de pruebas: ", err)
	}
	if !helpers.BaseDisponible() {
		log.Println("AVISO: variables COTIZACIONES_CRUD_PG* no definidas; las pruebas de integración se omiten")
	}
{{- else}}
	if err := helpers.ConectarFalso(); err != nil {
		log.Fatal("No se pudo registrar la base de datos falsa: ", err)
	}
{{- end}}
{{- if .Controladores}}
	// Equivalen a copyrequestbody = true y autorender = false de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	// Mismas respuestas de error (Success/Status/Message/Data) que en main.go
	beego.ErrorController(&customerror.CustomErrorController{})
{{- end}}
	os.Exit(m.Run())
}
`))

// plantillaUnitControlador emite las pruebas unitarias del POST de un controlador.
// Cada escenario aplica a todos los campos opcionales de la entidad.
var plantillaUnitControlador = template.Must(template.New("unitControlador").Parse(`
// Pruebas unitarias del POST de {{.Nombre}}Controller ({{.Ruta}}), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const ruta{{.Nombre}} = "{{.Ruta}}"

// ejemplo{{.Nombre}} retorna la entidad con todos sus campos de ejemplo.
func ejemplo{{.Nombre}}() *models.{{.Nombre}} {
	return &models.{{.Nombre}}{
	{{- range .Campos}}{{if .Externo}}
		{{.Nombre}}: {{.Externo}},
	{{- end}}{{end}}
	}
}

func cuerpo{{.Nombre}}(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemplo{{.Nombre}}())
}

// TestPost{{.Nombre}}OpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPost{{.Nombre}}OpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	delete(cuerpo, "{{.Nombre}}")
	{{- end}}
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPost{{.Nombre}}OpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPost{{.Nombre}}OpcionalesNull(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	cuerpo["{{.Nombre}}"] = nil
	{{- end}}
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPost{{.Nombre}}OpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPost{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo{{.Nombre}}(t), []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}

// TestGet{{.Nombre}}OpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGet{{.Nombre}}OpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, ruta{{.Nombre}}, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestGet{{.Nombre}}OpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGet{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarGetOpcionales(t, ruta{{.Nombre}}, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPut{{.Nombre}}OpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	delete(cuerpo, "{{.Nombre}}")
	{{- end}}
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPut{{.Nombre}}OpcionalesNull(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	cuerpo["{{.Nombre}}"] = nil
	{{- end}}
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPut{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo{{.Nombre}}(t), []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}
`))

// plantillaUnitModelo emite, por campo opcional, las pruebas unitarias de
// models.Add<Modelo> con tres subpruebas: omitido (campo no asignado), null (campo
// asignado a nil) y con valor.
var plantillaUnitModelo = template.Must(template.New("unitModelo").Parse(`
// Pruebas unitarias de models.Add{{.Nombre}} para los campos escalares opcionales, sin base de datos.
{{range $f := .EscalaresOpc}}
// Test{{$.Nombre}}{{$f.Nombre}} verifica el valor de {{$f.Columna}} en el INSERT para
// los tres casos del campo {{$f.Nombre}}.
func Test{{$.Nombre}}{{$f.Nombre}}(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
			{{$f.Nombre}}: nil,
		}
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if $c.Externo}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		esperado := helpers.Desreferenciar(m.{{$f.Nombre}})
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", esperado)
	})
}

// Test{{$.Nombre}}{{$f.Nombre}}Get verifica la lectura de {{$f.Columna}}: NULL en la base
// de datos llega como nil y un valor llega igual.
func Test{{$.Nombre}}{{$f.Nombre}}Get(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"{{$f.Columna}}": nil})
		v, err := models.Get{{$.Nombre}}ById(1)
		if err != nil {
			t.Fatal("No se pudo consultar {{$.Nombre}}:", err)
		}
		helpers.VerificarValorLeido(t, "{{$f.Nombre}}", nil, helpers.Desreferenciar(v.{{$f.Nombre}}))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar({{$f.Externo}})
		helpers.DefinirFila(map[string]interface{}{"{{$f.Columna}}": esperado})
		v, err := models.Get{{$.Nombre}}ById(1)
		if err != nil {
			t.Fatal("No se pudo consultar {{$.Nombre}}:", err)
		}
		helpers.VerificarValorLeido(t, "{{$f.Nombre}}", esperado, helpers.Desreferenciar(v.{{$f.Nombre}}))
	})
}

// Test{{$.Nombre}}{{$f.Nombre}}Update verifica el valor de {{$f.Columna}} en el UPDATE
// para los tres casos del campo {{$f.Nombre}}.
func Test{{$.Nombre}}{{$f.Nombre}}Update(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
			{{$f.Nombre}}: nil,
		}
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if $c.Externo}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		esperado := helpers.Desreferenciar(m.{{$f.Nombre}})
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", esperado)
	})
}
{{end}}`))
