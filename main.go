package main

import (
	"fmt"
	"os"
	"ph1-assembly/input"
	"strings"
)

// Mount lê um arquivo fonte em assembly PH1 e monta para linguagem de máquina
// no padrão do emulador PH1
func Mount(opt *Options) {
	source, err := input.ReadSource(opt.Input)

	if err != nil {
		panic("Cannot read input file")
	}

	for _, srcLine := range source.Contents {
		fmt.Printf("%s: %s %s\n", srcLine.Label, srcLine.Name, srcLine.Operand)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Informe o nome do arquivo")
	}
	var input, output string

	// Pega o arquivo de entrada dos argumentos
	input = os.Args[1]

	// Verifica o output
	if len(os.Args) == 3 {
		output = os.Args[2]
	} else {

		// Gera o output através do nome do arquivo
		output = strings.Split(input, ".")[0] + ".ph1"
	}

	options := &Options{
		Input:    input,
		Output:   output,
		Compress: false,
	}

	Mount(options)
}
