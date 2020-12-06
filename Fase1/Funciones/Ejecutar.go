package Funciones

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func InstruccionsCommand(lcomando string) {
	lcomando = strings.ToLower(lcomando)
	if strings.HasPrefix(lcomando, "#") {
		return
	} else if strings.HasPrefix(lcomando, "exec") {
		Exec(lcomando)
	} else if strings.HasPrefix(lcomando, "pause") {
		var pausa string
		fmt.Scanf("%s\n", &pausa)
	} else if strings.HasPrefix(lcomando, "mkdisk") {
		CreateDisk()
	} else if strings.HasPrefix(lcomando, "rmdisk") {
		fmt.Println("RMDISK")
	} else if strings.HasPrefix(lcomando, "fdisk") {
		fmt.Println("FKDISK ")
	} else if strings.HasPrefix(lcomando, "mount") {
		fmt.Println("mount montar disco")
	} else if strings.HasPrefix(lcomando, "unmount") {
		fmt.Println("unmount desmontar disco")
	} else {
		fmt.Println("Instruccion desconocida")
	}
}

func Exec(comando string) {
	comando = strings.ReplaceAll(comando, "exec ", "")
	parametros := strings.Split(comando, " ")
	if len(parametros) > 1 {
		fmt.Println("Un parametro desconocido")
	} else {
		Verificarparametro(parametros)
	}
}

func Verificarparametro(parametro []string) {
	if strings.HasPrefix(parametro[0], "-path->") {
		ruta := strings.ReplaceAll(parametro[0], "-path->", "")
		AbrirArchivo(ruta)
	} else {
		fmt.Println("Parametro Desconocido")
	}
}

func AbrirArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Print("Comando: ")
		fmt.Println(scanner.Text())
		InstruccionsCommand(scanner.Text())
	}
}
