package Funciones

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"../Estructuras"
)

func Montaje(comando string) {
	comando = strings.ReplaceAll(comando, "mount ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 2 {
		fmt.Println("Un parametro no pertenece")
	} else {
		ParamVerificationMount(parametros)
	}
}

func ParamVerificationMount(param []string) {
	var Directorio, name string
	for _, parametro := range param {
		//fmt.Println("Evaluo " + parametro)
		if strings.HasPrefix(parametro, "-path->") {
			Directorio = strings.ReplaceAll(parametro, "-path->", "")
			if !ExisteRuta(Directorio) {
				fmt.Println("La ruta no existe")
				return
			}
		} else if strings.HasPrefix(parametro, "-name->") {
			name = strings.ReplaceAll(parametro, "-name->", "")
		} else {
			fmt.Println("Parametro Desconocido")
			return
		}
	}
	AMontar(Directorio, name)

}

func AMontar(ruta string, nombre string) {
	file, err := os.Open(ruta)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	discotemp := Estructuras.Mbr{}
	var tazo int = int(unsafe.Sizeof(discotemp))
	file.Seek(0, 0)
	discotemp = leerDisco(file, tazo, discotemp)
	if !VerificarNombre(discotemp, nombre) {
		fmt.Println("No EXISTE Particion")
		file.Close()
		return
	}
	for _, parti := range ListaPartM {
		if parti.Name == nombre && parti.DiscoR == ruta {
			fmt.Println("La particion ya esta montada")
			return
		}
	}
	file.Close()
	indice := BuscarenLista(ruta)
	letra := ListaDiscos[indice].Identificador
	nume := ListaDiscos[indice].IDs
	ListaDiscos[indice].IDs = ListaDiscos[indice].IDs + 1
	NuevoPartM := Estructuras.ParticionMontada{}
	Ident := "vd" + string(letra) + strconv.Itoa(nume)
	NuevoPartM.Identificador = Ident
	NuevoPartM.DiscoR = ruta
	NuevoPartM.Name = nombre
	ListaPartM = append(ListaPartM, NuevoPartM)
	fmt.Println("Se monto la particion " + Ident + " Con exito")

}

func Desmontaje(comando string) {
	comando = strings.ReplaceAll(comando, "unmount ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 1 {
		fmt.Println("Un parametro no pertenece")
	} else {
		ParamVerificationUnmount(parametros)
	}
}
func ParamVerificationUnmount(parametros []string) {
	var id string
	for _, parametro := range parametros {
		//fmt.Println("Evaluo " + parametro)
		if strings.HasPrefix(parametro, "-id->") {
			id = strings.ReplaceAll(parametro, "-id->", "")
			for _, particion := range ListaPartM {
				if particion.Identificador == id {
					particion.Identificador = ""
					fmt.Println("Se desmonto")
					return
				}
			}
		} else {
			fmt.Println("Parametro Desconocido")
			return
		}
	}
}

func BuscarenLista(ruta string) int {
	for i, disco := range ListaDiscos {
		if disco.Path == ruta {
			return i
		}
	}
	return 0
}
