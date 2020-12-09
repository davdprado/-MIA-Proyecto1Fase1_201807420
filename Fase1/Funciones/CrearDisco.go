package Funciones

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateDisk(comando string) {
	comando = strings.ReplaceAll(comando, "mkdisk ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 4 {
		fmt.Println("Un parametro no pertenece")
	} else {
		DiskParamVerification(parametros)
	}
}
func DiskParamVerification(param []string) {
	var Directorio string
	var tam int
	unidad := "m"
	fit := "ff"
	for _, parametro := range param {
		if strings.HasPrefix(parametro, "-path->") {
			ruta := strings.ReplaceAll(parametro, "-path->", "")
			fmt.Println(ruta)
			Directorio = ruta
		} else if strings.HasPrefix(parametro, "-size->") {
			tamo := strings.ReplaceAll(parametro, "-size->", "")
			sizee, _ := strconv.Atoi(tamo)
			if sizee > 0 {
				tam = sizee
			} else {
				fmt.Println("Tamaño invalido")
				return
			}
		} else if strings.HasPrefix(parametro, "-unit->") {
			unit := strings.ReplaceAll(parametro, "-unit->", "")
			if unit == "k" || unit == "m" {
				unidad = unit
			} else {
				fmt.Println("Medida invalida")
				return
			}

		} else if strings.HasPrefix(parametro, "-fit->") {
			aju := strings.ReplaceAll(parametro, "-unit->", "")
			if aju == "bf" || aju == "wf" || aju == "ff" {
				fit = aju
			} else {
				fmt.Println("Ajuste invalido")
				return
			}
		} else {
			fmt.Println("Parametro Desconocido")
			return
		}
	}
	fmt.Printf("Disco creado en %s de tamaño: %d %s con el ajuste %s\n", Directorio, tam, unidad, fit)
	CreateBin(Directorio, tam, fit, unidad)
}

func CreateBin(ruta string, size int, fit string, unida string) {
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
	// llenar el archivo con datos
	var temporal int8 = 0
	s := &temporal
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)

	if unida == "k" {

		file.Seek(int64(size)*1024, 0)
	} else if unida == "m" {
		file.Seek(int64(size)*1024*1024, 0)
	}

	file.Close()

}

func LlenardeBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}
