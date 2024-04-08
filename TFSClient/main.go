package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TSF/TFSClient/Business"
	"github.com/TSF/TFSClient/Commons"
	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "> ",
		HistoryFile: "~/.my_history",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	fmt.Println("Por favor escribe un comando")
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		palabras := strings.Split(line, " ")

		if palabras[0] == "Particionar" {
			chunkNumber, err := strconv.Atoi(palabras[2])

			file, err := os.Open(palabras[1])
			Business.PartitionFile(file, chunkNumber)
			fmt.Println(err)
		} else if palabras[0] == "Buscar" {
			palabras := strings.Split(line, " ")
			fmt.Println(palabras[1])
			Business.RestoreFile(palabras[1])
		} else if palabras[0] == "Conectar" {
			var request Commons.NameNodeRequest
			request.ChunksQuantity = 2
			request.File_name = "prueba"
			respuesta, err := Business.WritteAlert(request)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(respuesta)
			}
		} else {
			fmt.Println("Comando no reconocido: ", line)
			fmt.Println("Comandos permitidos: ")
			fmt.Println("Particionar *path del archivo y numero de chuncks todo separado por un unico espacio")
			fmt.Println("Buscar *Nombre del archivo, con su extensión")
		}
	}
}
