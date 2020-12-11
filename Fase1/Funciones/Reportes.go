package Funciones

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	if nombre == "mbr" {
		GrficaTabla(ruta, id)
	} else if nombre == "disk" {
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
	archi, err := os.Open(ruta)
	defer archi.Close()
	if err != nil {
		panic(err)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(archi, tazo, discotemp)
	file.Close()

	CorrerDot(ruta, Extension)

}
func GraficaDisk(ruta string, id string) {

}
func CorrerDot(ruta string, extencion string) {
	out, err := exec.Command("dot", "-T"+extencion, "GraficaC.dot", "-o", ruta).Output()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
