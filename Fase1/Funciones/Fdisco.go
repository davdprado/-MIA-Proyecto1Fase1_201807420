package Funciones

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"../Estructuras"
)

func AdminDisco(comando string) {
	comando = strings.ReplaceAll(comando, "fdisk ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 7 {
		fmt.Println("Un parametro no pertenece")
	} else {
		AdminParamVerification(parametros)
	}
}

func AdminParamVerification(param []string) {
	var Directorio string
	var pdelete, nombreP string
	var tam, vals int
	tam = 0
	vals = 0
	fdelete := false
	fadd := false
	unidad := "k"
	tipo := "p"
	fit := "ff"
	for _, parametro := range param {
		if strings.HasPrefix(parametro, "-path->") {
			ruta := strings.ReplaceAll(parametro, "-path->", "")
			Directorio = ruta
			if !ExisteRuta(Directorio) {
				fmt.Println("La ruta no existe")
				return
			}
		} else if strings.HasPrefix(parametro, "-name->") {
			nombreP = strings.ReplaceAll(parametro, "-name->", "")
		} else if strings.HasPrefix(parametro, "-size->") {
			tamo := strings.ReplaceAll(parametro, "-size->", "")
			sizee, _ := strconv.Atoi(tamo)
			if sizee > 0 {
				tam = sizee
			} else {
				fmt.Println("Tamaño invalido")
				return
			}
		} else if strings.HasPrefix(parametro, "-delete->") {
			pdelet := strings.ReplaceAll(parametro, "-delete->", "")
			fdelete = true
			if pdelete == "fast" || pdelete == "full" {
				pdelete = pdelet
			} else {
				fmt.Println("Tipo de borrado desconocido")
			}
		} else if strings.HasPrefix(parametro, "-add->") {
			valss := strings.ReplaceAll(parametro, "-add->", "")
			vals, _ = strconv.Atoi(valss)
			fadd = true
		} else if strings.HasPrefix(parametro, "-unit->") {
			unit := strings.ReplaceAll(parametro, "-unit->", "")
			if unit == "k" || unit == "m" || unit == "b" {
				unidad = unit
			} else {
				fmt.Println("Medida invalida")
				return
			}

		} else if strings.HasPrefix(parametro, "-type->") {
			tipos := strings.ReplaceAll(parametro, "-type->", "")
			if tipos == "p" || tipos == "e" || tipos == "l" {
				tipo = tipos
			} else {
				fmt.Println("tipo invalido")
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
	secondRevision(Directorio, nombreP, tam, unidad, fdelete, fadd, pdelete, vals, tipo, fit)
	//fmt.Printf("Ruta: %s  NombredeP: %s  tamdeP: %d en %s Se borra?: %b Se agrega?: %b ValorAgrega: %d en %s Ptipo: %s Fit %s tipoBorrado %s\n", Directorio, nombreP, tam, unidad, fdelete, fadd, vals, unidad, tipo, fit, pdelete)
}

func secondRevision(ruta string, nombrep string, size int, unidad string, isdelete bool, isadd bool, tborrado string, valadd int, tipoParticion string, ajuste string) {
	file, err := os.Open(ruta)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(file, tazo, discotemp)
	fmt.Printf("Fecha: %s Tamaño: %d Random: %d\n", discotemp.Mfecha, discotemp.Mtamano, discotemp.MdiscoA)
	defer file.Close()

	if isadd || isdelete {
		if isadd {
			isdelete = false
		} else if isdelete {
			isadd = false
		}
	}
	if isdelete {
		EliminarParticion()

	} else if isadd {
		ModificarParticion()
	} else {
		CrearParticion(nombrep, size, unidad, ajuste, tipoParticion, discotemp)
	}

}
func CrearParticion(nombre string, tama int, tipoF string, tipotam string, tipoP string, disco Estructuras.Mbr) {
	particiontemp := Estructuras.Particion{}
	particiontemp.PartStatus = true
	copy(particiontemp.PartType[:], tipoP)
	copy(particiontemp.PartFit[:], tipoF)
	copy(particiontemp.PartName[:], nombre)
	particiontemp.PartSize = int64(tama)
	partiNo := len(disco.MParticiones)
	if disco.MParticiones {

	}
	fmt.Println("Crear la P")
}
func EliminarParticion() {
	fmt.Println("Eliminar la P")

}
func ModificarParticion() {
	fmt.Println("Modificar la P")

}
func leerDisco(file *os.File, size int, disco Estructuras.Mbr) Estructuras.Mbr {
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &disco)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	return disco
}
func leerBytes(file *os.File, numero int) []byte {
	bytes := make([]byte, numero)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
