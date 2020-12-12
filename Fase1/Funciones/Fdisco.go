package Funciones

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
			if tipos == "p" || tipos == "e" {
				tipo = tipos
			} else if tipos == "l" {
				fmt.Println("Es Logica")
				return
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
	if isadd || isdelete {
		if isadd {
			isdelete = false
		} else if isdelete {
			isadd = false
		}
	}
	if isdelete {
		EliminarParticion(nombrep, ruta)

	} else if isadd {
		ModificarParticion()
	} else {
		if unidad == "b" {
			CrearParticion(nombrep, size, unidad, ajuste, tipoParticion, ruta)
		} else if unidad == "m" {
			CrearParticion(nombrep, size*1024*1024, unidad, ajuste, tipoParticion, ruta)
		} else if unidad == "k" {
			CrearParticion(nombrep, size*1024, unidad, ajuste, tipoParticion, ruta)
		}

	}

}
func CrearParticion(nombre string, tama int, tipoF string, tipotam string, tipoP string, ruta string) {
	file, err := os.OpenFile(ruta, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(file, tazo, discotemp)
	if VerificarNombre(discotemp, nombre) {
		fmt.Println("YA EXISTE Particion")
		file.Close()
		return
	}
	if TieneExtend(discotemp) && tipoP == "e" {
		fmt.Println("YA EXISTE EXTENDIDA en este disco")
		file.Close()
		return
	}
	particionnnn := Estructuras.ParticionMontada{}
	particionnnn.Estado = '1'
	particionnnn.Tamaño = int64(tama)
	particionnnn.Name = nombre
	particionnnn.Tipo = tipoP
	particionnnn.Fit = tipoF
	particionnnn.EstadoEscrito = false
	//Disk.Particiones = append(Disk.Particiones, particionnnn)
	//fmt.Printf("Fecha: %s Tamaño: %d Random: %d\n", discotemp.Mfecha, discotemp.Mtamano, discotemp.MdiscoA)
	particiontemp := Estructuras.Particion{}
	particiontemp.PartStatus = '1'
	copy(particiontemp.PartType[:], tipoP)
	copy(particiontemp.PartFit[:], tipoF)
	copy(particiontemp.PartName[:], nombre)
	particiontemp.PartSize = int64(tama)
	particiontemp.PartStart = discotemp.Mbit
	discotemp.Mbit = int64(unsafe.Sizeof(particiontemp)) + discotemp.Mbit
	for i := 0; i < 4; i++ {
		if discotemp.MParticiones[i].PartStatus == '0' {
			discotemp.MParticiones[i] = particiontemp
			break
		}
	}

	file.Seek(0, 0)
	var bufferDisco bytes.Buffer
	binary.Write(&bufferDisco, binary.BigEndian, &discotemp)
	escribirBytes(file, bufferDisco.Bytes())
	file.Close()
}
func EliminarParticion(nombre string, ruta string) {
	file, err := os.OpenFile(ruta, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(file, tazo, discotemp)
	if VerificarNombre(discotemp, nombre) == false {
		fmt.Println("No Existe  Particion con ese nombre")
		file.Close()
		return
	}
	//Disk.Particiones = append(Disk.Particiones, particionnnn)
	//fmt.Printf("Fecha: %s Tamaño: %d Random: %d\n", discotemp.Mfecha, discotemp.Mtamano, discotemp.MdiscoA)
	particiontemp := Estructuras.Particion{}
	particiontemp.PartStatus = '0'
	var nomp [16]byte
	for k, v := range []byte(nombre) {
		nomp[k] = byte(v)
	}
	copy(particiontemp.PartName[:], "Libre")
	for i := 0; i < 4; i++ {
		if discotemp.MParticiones[i].PartName == nomp {
			sisou := discotemp.MParticiones[i].PartSize
			particiontemp.PartSize = sisou
			discotemp.MParticiones[i] = particiontemp
			break
		}
	}

	file.Seek(0, 0)
	var bufferDisco bytes.Buffer
	binary.Write(&bufferDisco, binary.BigEndian, &discotemp)
	escribirBytes(file, bufferDisco.Bytes())
	file.Close()
	fmt.Println("Particion Eliminada")
}
func ModificarParticion() {
	fmt.Println("Modificar la P")

}
func leerDisco(file *os.File, size int, disco Estructuras.Mbr) Estructuras.Mbr {
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &disco)
	if err != nil {
		panic(err)
	}
	return disco
}
func leerBytes(file *os.File, numero int) []byte {
	bytes := make([]byte, numero)

	_, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

func TieneExtend(disco Estructuras.Mbr) bool {
	var auxt [2]byte
	for i, j := range []byte("e") {
		auxt[i] = byte(j)
	}
	for i := 0; i < 4; i++ {
		if disco.MParticiones[i].PartType == auxt {
			return true
		}
	}
	return false
}

func VerificarNombre(disco Estructuras.Mbr, nombre string) bool {
	var auxt [16]byte
	for i, j := range []byte(nombre) {
		auxt[i] = byte(j)
	}
	for i := 0; i < 4; i++ {
		if disco.MParticiones[i].PartName == auxt {
			return true
		}
	}
	return false
}
