package Funciones

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"

	"../Estructuras"
)

func VerifiReportes(comando string) {
	comando = strings.ReplaceAll(comando, "rep ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 3 {
		fmt.Println("Un parametro no pertenece")
	} else {
		ReportParamVerification(parametros)
	}
}
func ReportParamVerification(param []string) {
	var Directorio, name, id string
	for _, parametro := range param {
		if strings.HasPrefix(parametro, "-path->") {
			Directorio = strings.ReplaceAll(parametro, "-path->", "")
		} else if strings.HasPrefix(parametro, "-name->") {
			name = strings.ReplaceAll(parametro, "-name->", "")
		} else if strings.HasPrefix(parametro, "-id->") {
			id = strings.ReplaceAll(parametro, "-id->", "")
		} else {
			fmt.Println("Parametro Desconocido")
			return
		}
	}
	TipodeGrafica(Directorio, name, id)

}
func TipodeGrafica(ruta string, nombre string, id string) {
	if nombre == "disk" {
		GrficaTabla(ruta, id)
	} else if nombre == "mbr" {
		GraficaDisk(ruta, id)
	} else {
		fmt.Println("Tipo de Grapho invalido")
		return
	}
}
func obtenerExtencion(ruta string) string {
	arreglo := strings.Split(ruta, "/")
	ultimo := len(arreglo)
	archivito := arreglo[ultimo-1]
	miniarr := strings.Split(archivito, ".")
	return miniarr[1]
}

func GrficaTabla(ruta string, id string) {
	var direccion = ""
	arreglo := strings.Split(ruta, "/")
	for i := 0; i < (len(arreglo) - 1); i++ {
		direccion = direccion + arreglo[i] + "/"
	}
	err := os.MkdirAll(direccion, 0777)
	file, err := os.Create(ruta)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	Extension := obtenerExtencion(ruta)
	rutdos := ""
	for _, particion := range ListaPartM {
		if particion.Identificador == id {
			rutdos = particion.DiscoR
			break
		} else {
			fmt.Println("La particion " + id + " No existe")
			return
		}
	}
	archi, erro := os.Open(rutdos)
	defer archi.Close()
	if erro != nil {
		panic(erro)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(archi, tazo, discotemp)
	file.Close()
	// escribiremaso el codigo de la tabla
	codigo := "digraph MBR {\n graph [ label = \"Reporte MBR\"];\n node [shape = plain]\n randir =TB\n"
	codigo += "mbr[label=<\n"
	codigo += "<table border=\"1\" cellborder=\"1\" cellspacing=\"0\">\n"
	codigo += "<tr><td colspan='2'> Disco" + obtenerNombre(rutdos) + "</td></tr>͜\n"
	codigo += "<tr><td>mbr_tamaño</td><td>" + strconv.FormatInt(discotemp.Mtamano, 10) + "</td></tr>\n"
	nn := bytes.Index(discotemp.Mfecha[:], []byte{0})
	codigo += "<tr><td>mbr_fecha</td><td>" + string(discotemp.Mfecha[:nn]) + "</td></tr>\n"
	codigo += "<tr><td>mbr_disk_signature</td><td>" + strconv.FormatInt(discotemp.MdiscoA, 10) + "</td></tr>\n"
	codigo += "<tr><td>Particiones</td></tr>\n"
	for i := 0; i < 4; i++ {
		codigo += "<tr><td>Part_status</td><td>" + string(discotemp.MParticiones[i].PartStatus) + "</td></tr>\n"
		nm := bytes.Index(discotemp.MParticiones[i].PartType[:], []byte{0})
		codigo += "<tr><td>Part_Type</td><td>" + string(discotemp.MParticiones[i].PartType[:nm]) + "</td></tr>\n"
		codigo += "<tr><td>Part_start</td><td>" + strconv.FormatInt(discotemp.MParticiones[i].PartStart, 10) + "</td></tr>\n"
		codigo += "<tr><td>Part_size</td><td>" + strconv.FormatInt(discotemp.MParticiones[i].PartSize, 10) + "</td></tr>\n"
		na := bytes.Index(discotemp.MParticiones[i].PartName[:], []byte{0})
		codigo += "<tr><td>Part_name</td><td>" + string(discotemp.MParticiones[i].PartName[:na]) + "</td></tr>\n"
		nf := bytes.Index(discotemp.MParticiones[i].PartFit[:], []byte{0})
		codigo += "<tr><td>Part_fit</td><td>" + string(discotemp.MParticiones[i].PartFit[:nf]) + "</td></tr>\n"
	}
	/* codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n"
	codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n"
	codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n"
	codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n"
	codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n"
	codigo += "<tr><td>"++"</td><td>"++"</td></tr>\n" */
	codigo += "\n</table>\n"
	codigo += ">];\n"
	codigo += "}"
	arrCode := []byte(codigo)
	errr := ioutil.WriteFile("GraficaT.dot", arrCode, 0644)
	if errr != nil {
		panic(errr)
	}

	CorrerDot(ruta, Extension)
	fmt.Println("Se creo la grafica")

}
func GraficaDisk(ruta string, id string) {
	var direccion = ""
	arreglo := strings.Split(ruta, "/")
	for i := 0; i < (len(arreglo) - 1); i++ {
		direccion = direccion + arreglo[i] + "/"
	}
	err := os.MkdirAll(direccion, 0777)
	file, err := os.Create(ruta)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	Extension := obtenerExtencion(ruta)
	rutdos := ""
	for _, particion := range ListaPartM {
		if particion.Identificador == id {
			rutdos = particion.DiscoR
			break
		} else {
			fmt.Println("La particion " + id + " No existe")
			return
		}
	}
	archi, erro := os.Open(rutdos)
	defer archi.Close()
	if erro != nil {
		panic(erro)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(archi, tazo, discotemp)
	file.Close()
	TamaDisco := discotemp.Mtamano
	Part1 := discotemp.MParticiones[0].PartSize
	Part2 := discotemp.MParticiones[1].PartSize
	Part3 := discotemp.MParticiones[2].PartSize
	Part4 := discotemp.MParticiones[3].PartSize
	/* fmt.Println("Tamaño de Disco")
	fmt.Println(discotemp.Mtamano)
	fmt.Println("Tamaño de part")
	fmt.Println(discotemp.MParticiones[0].PartSize)
	fmt.Println(discotemp.MParticiones[1].PartSize)
	fmt.Println(discotemp.MParticiones[2].PartSize)
	fmt.Println(discotemp.MParticiones[3].PartSize) */
	var Por [4]int64
	Por[0] = (Part1 * 100) / TamaDisco
	Por[1] = (Part2 * 100) / TamaDisco
	Por[2] = (Part3 * 100) / TamaDisco
	Por[3] = (Part4 * 100) / TamaDisco
	EspacioLibre := TamaDisco - (Part1 + Part2 + Part3 + Part4)
	LiP := (EspacioLibre * 100) / TamaDisco
	fmt.Println("Libre de disco " + strconv.FormatInt(EspacioLibre, 10))
	fmt.Println("1: " + strconv.FormatInt(Por[0], 10) + "%")
	fmt.Println("2: " + strconv.FormatInt(Por[1], 10) + "%")
	fmt.Println("3: " + strconv.FormatInt(Por[2], 10) + "%")
	fmt.Println("4: " + strconv.FormatInt(Por[3], 10) + "%")
	fmt.Println("N: " + strconv.FormatInt(LiP, 10) + "%")
	// escribiremaso el codigo de la tabla

	p := "p"
	e := "e"
	var arrp [2]byte
	for k, v := range []byte(p) {
		arrp[k] = byte(v)
	}
	var arre [2]byte
	for k, v := range []byte(e) {
		arre[k] = byte(v)
	}

	codigo := "digraph G {\n graph [ label = \"" + obtenerNombre(rutdos) + "\"];\n node [shape = record];\n "
	codigo += "5[label=\"MBR|"
	for i := 0; i < 4; i++ {
		if discotemp.MParticiones[i].PartSize > 0 {
			if discotemp.MParticiones[i].PartType == arrp {
				codigo += "Primaria " + strconv.FormatInt(Por[i], 10) + "% |"
			} else if discotemp.MParticiones[i].PartType == arre {
				codigo += "{Extendida " + strconv.FormatInt(Por[i], 10) + "% |{EBR|Libre}}|"
			} else if discotemp.MParticiones[i].PartStatus == '0' && discotemp.MParticiones[i].PartSize > 0 {
				codigo += "Libre " + strconv.FormatInt(Por[i], 10) + "% |"
			}
		}
	}
	if EspacioLibre > 0 {
		codigo += "Libre " + strconv.FormatInt(LiP, 10) + " %"
	}
	codigo += "|\"];\n"
	codigo += "}"
	arrCode := []byte(codigo)
	errr := ioutil.WriteFile("GraficaT.dot", arrCode, 0644)
	if errr != nil {
		panic(errr)
	}

	CorrerDot(ruta, Extension)
	fmt.Println("Se creo la grafica")

}
func CorrerDot(ruta string, extencion string) {
	out, err := exec.Command("dot", "-T"+extencion, "GraficaT.dot", "-o", ruta).Output()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

func convert(name []byte) string {
	data := ""
	for i := 0; i < len(name); i++ {
		if name[i] != 0 {
			data += string(name[i])
		}
	}
	return data
}

func obtenerNombre(ruta string) string {
	arreglo := strings.Split(ruta, "/")
	numero := len(arreglo)
	nombre := arreglo[numero-1]
	return nombre

}
