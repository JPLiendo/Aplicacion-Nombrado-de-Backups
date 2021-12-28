package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func nombrado(tit, opc string, nOpc, lOpc int, seg []string) string {
	t := tit
	o := opc
	n := nOpc
	s := seg
	l := lOpc
	for {
		fmt.Println("")
		fmt.Printf(`	Las opciones son: %v`, o)
		fmt.Println("")
		fmt.Printf(`	Ingrese el valor deseado: `)
		fmt.Scan(&n)

		switch {
		case n == 1 && n <= l:
			t = t + s[0]

		case n == 2 && n <= l:
			t = t + s[1]

		case n == 3 && n <= l:
			t = t + s[2]

		case n == 4 && n <= l:
			t = t + s[3]

		case n == 5 && n <= l:
			t = t + s[4]

		case n == 6 && n <= l:
			t = t + s[5]

		case n == 7 && n <= l:
			t = t + s[6]

		case n == 8 && n <= l:
			t = t + s[7]

		case n == 9 && n <= l:
			t = t + s[8]

		case n == 10 && n <= l:
			t = t + s[9]

		default:
			fmt.Println("El valor ingredado no es el correcto, ingrese la opcion nuevamente")
			continue
		}
		break
	}
	return fmt.Sprint(t)
}

func lectura() []fs.FileInfo {
	fmt.Println(`	Los archivos exixtentes son:`)
	archivos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for k, archivo := range archivos {
		fmt.Printf("					Opcion: %v  Nombre: %v\n", k, archivo.Name())

	}
	return archivos
}

func renombrar(archivos []fs.FileInfo, n string) {
	old := archivos
	new := n
	var e error

	fmt.Printf(`	Ingrese el nombre del archivo a renombrar: `)
	i := 0
	fmt.Scan(&i)
	e = os.Rename(old[i].Name(), new)
	if e != nil {
		log.Fatal(e)

	}
	fmt.Println(`	-------------------------------------------`)
	fmt.Println(`	El archivo ha sido renombrado exitosamente.`)
	fmt.Println(`	-------------------------------------------`)
}

func t_formando(t string) {
	titulo := t
	fmt.Println(`	----------------------------------------------------------`)
	fmt.Println(`	Titulo formado:`, titulo)
	fmt.Println(`	----------------------------------------------------------`)
}

func zipSource(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func main() {

	// Cabecera.

	cabecera := `
	*************************************************************************************************	
	                Bienvenido a la aplicacion de nombrado de BACKUPs
	*************************************************************************************************
	`
	fmt.Println(cabecera)

	//Inicio de primer segmento "TIPO DE RESPALDO"
	titulo := ""
	nOpcion := 0
	lOpcion := 2
	segmento_1 := []string{"BCKP-", "IMG-"}
	tipo := `
			1  "BCKP"
			2  "IMG"
	`
	titulo = nombrado(titulo, tipo, nOpcion, lOpcion, segmento_1)
	t_formando(titulo)

	//Inicio del segundo segmento "LINEA"
	lOpcion = 5
	segmento_2 := []string{"L08-", "L09-", "L10-", "L23-", "L30-"}
	linea := `
			1  "LINEA 08"
			2  "LINEA 09"
			3  "LINEA 10"
			4  "LINEA 23"
			5  "LINEA 30"
	`
	titulo = nombrado(titulo, linea, nOpcion, lOpcion, segmento_2)
	t_formando(titulo)

	//Incio del tercer segmento "Dispositivo".
	lOpcion = 5
	segmento_3 := []string{"PLC-", "HMI-", "VFR-", "DRI-", "RFID-"}
	dispositivo := `
			1  "PLC"
			2  "HMI"
			3  "VFR"
			4  "DRI"
			5  "RFID"
	`
	titulo = nombrado(titulo, dispositivo, nOpcion, lOpcion, segmento_3)
	t_formando(titulo)

	//Inicio del cuarto segmento "Maquina".
	lOpcion = 7
	segmento_4 := []string{"SOP-", "LLE-", "TTB-", "TTP-", "ETI-", "HOR-", "PAL-"}
	maquina := `
			1  "SOP"
			2  "LLE"
			3  "TTB"
			4  "TTP"
			5  "ETI"
			6  "HOR"
			7  "PAL"
	`
	titulo = nombrado(titulo, maquina, nOpcion, lOpcion, segmento_4)
	t_formando(titulo)

	//Inicio del quinto segmento "Numero equipo"
	fmt.Println("")
	fmt.Printf(`	Ingrese el numero de equipo: `)
	t := ""
	fmt.Scan(&t)
	titulo = titulo + t + "-"
	fmt.Println("")

	//Inicio del sexto segmento "fecha"
	ahora := time.Now()
	a := ahora.Format("2006_01_02")
	titulo = titulo + a
	t_formando(titulo)
	fmt.Println("")

	//Inicio del septimo segmento "eleccion de archivo"
	lista := lectura()
	renombrar(lista, titulo)
	fmt.Println(`	Realizando la compresion de ` + titulo)
	time.Sleep(1 * time.Second)

	// Comprimiendo el archivo
	if err := zipSource("C:/Users/juanp/OneDrive/Documentos/Juan/GO/"+titulo, "C:/Users/juanp/OneDrive/Documentos/Juan/GO/"+titulo+".zip"); err != nil {
		log.Fatal(err)
	}

	//delay
	time.Sleep(3 * time.Second)
}
