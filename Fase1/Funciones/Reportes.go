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
	codigo += "<tr><td colspan='2'>MBR Disco x</td></tr>͜\n"
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
