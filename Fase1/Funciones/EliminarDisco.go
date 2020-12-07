package Funciones

import (
	"fmt"
	"os"
	"strings"
)

func EliminarDisco(comando string) {
	comando = strings.ReplaceAll(comando, "rmdisk ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 1 {
		fmt.Println("Un parametro desconocido")
	} else {
		verificarparametroE(parametros)
	}
}
func verificarparametroE(parametro []string) {
	if strings.HasPrefix(parametro[0], "-path->") {
		ruta := strings.ReplaceAll(parametro[0], "-path->", "")
		if ExisteRuta(ruta) {
			err := os.Remove(ruta)
			if err != nil {
				fmt.Printf("Error eliminando archivo: %v\n", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		} else {
			fmt.Println("No se encuentra la ruta")
		}

	} else {
		fmt.Println("Parametro Desconocido")
	}
}
