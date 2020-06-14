package main

import (
	"fmt"
	"os"
	"path/filepath"
	"ph1-assembly/constants"
	"ph1-assembly/extractor"
	"ph1-assembly/input"
	"ph1-assembly/output"
	"ph1-assembly/pherror"
	"strings"
)

// Options armazenam as opções do montador
type Options struct {
	InputFile  string
	OutputFile string
	Compress   bool
}

// Mount lê um arquivo fonte em assembly PH1 e monta para linguagem de máquina
// no padrão do emulador PH1
func Mount(opt *Options) {
	// Inicializa os logs de erro
	pherror.Setup(opt.InputFile)

	// Validação para permitir que o usuário tente mais vezes caso o nome do arquivo esteja
	// errado
	Input := input.ReadSource(opt.InputFile)

	// Primeira passagem: labels
	labels := extractor.ExtractLabels(Input)

	// Segunda passagem: instruções
	instructions, data := extractor.ExtractInstructionsAndData(Input, labels)

	// Gera o arquivo de saída a partir do nome definido no options
	output.CreateOutputFile(*instructions, *data, opt.OutputFile)
}

func main() {
	// Tratamento de erro
	defer func() {
		if err := recover(); err != nil {
			pherr := pherror.Format(err)
			fmt.Println(pherr)

			if pherr.Code != 0 {
				os.Exit(pherr.Code)
			}
			os.Exit(1)
		}
	}()

	if len(os.Args) < 2 {
		panic(pherror.MissingInputFile)
	}
	var input, output string

	// Pega o arquivo de entrada dos argumentos
	input = os.Args[1]

	// Verifica o output
	if len(os.Args) == 3 {
		output = os.Args[2]
	} else {

		// Gera o output através do nome do arquivo
		output = strings.Split(filepath.Base(input), ".")[0] + constants.OutputExtension
	}

	options := &Options{
		InputFile:  input,
		OutputFile: output,
		Compress:   false,
	}

	Mount(options)
}
