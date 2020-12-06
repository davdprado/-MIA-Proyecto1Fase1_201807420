package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./Funciones"
)

func main() {
	fmt.Print("Comando: ")
	var instruccion string
	leer := bufio.NewReader(os.Stdin)
	entrada, _ := leer.ReadString('\n')
	instruccion = strings.TrimRight(entrada, "\r\n")
	instruccion = strings.ToLower(instruccion)
	if instruccion == "exit" {
		println("UwU")
	} else {
		Funciones.InstruccionsCommand(instruccion)
		main()
	}

}
