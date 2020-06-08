package output

import (
	"fmt"
	"os"
	"ph1-assembly/extractor"
)

// CreateOutputFile cria e preenche o arquivo com todas as instruções
func CreateOutputFile(instruction []extractor.Instruction, outPutName string) {
	outputFile, err := os.Create(outPutName)

	if err != nil {
		panic(err)
	}

	// Para cada instructionValue contido na slice instruction escreve seu endereço e seu opCode, caso seja
	// uma instrução de dois bytes verifica os valores contidos no atributo Data da instrução e escreve seus
	// valores
	for _, instructionValue := range instruction {
		outputFile.WriteString(fmt.Sprintf("%02X %s\n", instructionValue.Address, instructionValue.OpCode))
		if instructionValue.HasOperand {
			outputFile.WriteString(fmt.Sprintf("%02X %02X\n", instructionValue.Data.Address,
				instructionValue.Data.Value))
		}
	}
}
