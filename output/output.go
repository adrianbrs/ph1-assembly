package output

import (
	"fmt"
	"os"
)

//Instruction representa uma instrução
type Instruction struct {
	Address    int
	OpCode     string
	Value      int
	HasOperand bool
}

// CreateOutputFile cria e preenche o arquivo com todas as instruções
func CreateOutputFile(instruction []Instruction) {
	outputFile, err := os.Create("outputFile.phu")

	if err != nil {
		panic(err)
	}

	// Para cada instructionValue contido na slice instruction escreve seu endereço e seu opCode, caso seja
	// uma instrução de dois bytes incrementa o endereço para manter a continuidade
	// e escreve o endereço onde está contido seu valor
	for _, instructionValue := range instruction {
		outputFile.WriteString(fmt.Sprintf("%02X %s\n", instructionValue.Address, instructionValue.OpCode))
		if instructionValue.HasOperand {
			outputFile.WriteString(fmt.Sprintf("%02X %02X\n", instructionValue.Address+1,
				instructionValue.Value))
		}
	}
}
