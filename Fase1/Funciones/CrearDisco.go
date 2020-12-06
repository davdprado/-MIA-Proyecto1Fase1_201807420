package Funciones

import (
	"fmt"
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
	fmt.Println("Crear archivo binario")
}
