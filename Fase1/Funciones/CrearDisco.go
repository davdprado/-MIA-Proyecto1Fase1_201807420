package Funciones

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"../Estructuras"
)

var iddisco = 0

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
			aju := strings.ReplaceAll(parametro, "-fit->", "")
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
	CreateBin(Directorio, tam, fit, unidad)
	//fmt.Printf("Disco creado en %s de tamaño: %d %s con el ajuste %s\n", Directorio, tam, unidad, fit)

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
	var tamao int64
	s := &temporal
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	if unida == "k" {
		file.Seek(int64(size)*1024, 0)
		tamao = int64(size * 1024)
		LlenardeBytes(file, binario.Bytes())

	} else if unida == "m" {
		file.Seek(int64(size)*1024*1024, 0)
		tamao = int64(size * 1024 * 1024)
		LlenardeBytes(file, binario.Bytes())
	}
	t := time.Now()
	fecha := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	file.Seek(0, 0)

	discoTemp := Estructuras.Mbr{}
	copy(discoTemp.Mfecha[:], fecha)
	copy(discoTemp.Mfit[:], fit)
	discoTemp.MdiscoA = rand.Int63n(1000)
	//tamao = int64(unsafe.Sizeof(discoTemp))
	discoTemp.Mtamano = tamao

	var bufferDisco bytes.Buffer
	binary.Write(&bufferDisco, binary.BigEndian, &discoTemp)
	escribirBytes(file, bufferDisco.Bytes())

	file.Close()
	//fmt.Printf("Disco: %d Fecha: %s Fit: %s Tamaño: %d\n", discoTemp.MdiscoA, fecha, fit, tamao)
}

func LlenardeBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}
